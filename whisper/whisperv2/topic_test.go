// Copyright 2015 The go-DATx Authors
// This file is part of the go-DATx library.
//
// The go-DATx library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-DATx library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-DATx library. If not, see <http://www.gnu.org/licenses/>.

package whisperv2

import (
	"bytes"
	"testing"
)

var topicCreationTests = []struct {
	data []byte
	hash [4]byte
}{
	{hash: [4]byte{0x8f, 0x9a, 0x2b, 0x7d}, data: []byte("test name")},
	{hash: [4]byte{0xf2, 0x6e, 0x77, 0x79}, data: []byte("some other test")},
}

func TestTopicCreation(t *testing.T) {
	// Create the topics individually
	for i, tt := range topicCreationTests {
		topic := NewTopic(tt.data)
		if !bytes.Equal(topic[:], tt.hash[:]) {
			t.Errorf("binary test %d: hash mismatch: have %v, want %v.", i, topic, tt.hash)
		}
	}
	for i, tt := range topicCreationTests {
		topic := NewTopicFromString(string(tt.data))
		if !bytes.Equal(topic[:], tt.hash[:]) {
			t.Errorf("textual test %d: hash mismatch: have %v, want %v.", i, topic, tt.hash)
		}
	}
	// Create the topics in batches
	binaryData := make([][]byte, len(topicCreationTests))
	for i, tt := range topicCreationTests {
		binaryData[i] = tt.data
	}
	textualData := make([]string, len(topicCreationTests))
	for i, tt := range topicCreationTests {
		textualData[i] = string(tt.data)
	}

	topics := NewTopics(binaryData...)
	for i, tt := range topicCreationTests {
		if !bytes.Equal(topics[i][:], tt.hash[:]) {
			t.Errorf("binary batch test %d: hash mismatch: have %v, want %v.", i, topics[i], tt.hash)
		}
	}
	topics = NewTopicsFromStrings(textualData...)
	for i, tt := range topicCreationTests {
		if !bytes.Equal(topics[i][:], tt.hash[:]) {
			t.Errorf("textual batch test %d: hash mismatch: have %v, want %v.", i, topics[i], tt.hash)
		}
	}
}

var topicMatcherCreationTest = struct {
	binary  [][][]byte
	textual [][]string
	matcher []map[[4]byte]struct{}
}{
	binary: [][][]byte{
		{},
		{
			[]byte("Topic A"),
		},
		{
			[]byte("Topic B1"),
			[]byte("Topic B2"),
			[]byte("Topic B3"),
		},
	},
	textual: [][]string{
		{},
		{"Topic A"},
		{"Topic B1", "Topic B2", "Topic B3"},
	},
	matcher: []map[[4]byte]struct{}{
		{},
		{
			{0x25, 0xfc, 0x95, 0x66}: {},
		},
		{
			{0x93, 0x6d, 0xec, 0x09}: {},
			{0x25, 0x23, 0x34, 0xd3}: {},
			{0x6b, 0xc2, 0x73, 0xd1}: {},
		},
	},
}

func TestTopicMatcherCreation(t *testing.T) {
	test := topicMatcherCreationTest

	matcher := newTopicMatcherFromBinary(test.binary...)
	for i, cond := range matcher.conditions {
		for topic := range cond {
			if _, ok := test.matcher[i][topic]; !ok {
				t.Errorf("condition %d; extra topic found: 0x%x", i, topic[:])
			}
		}
	}
	for i, cond := range test.matcher {
		for topic := range cond {
			if _, ok := matcher.conditions[i][topic]; !ok {
				t.Errorf("condition %d; topic not found: 0x%x", i, topic[:])
			}
		}
	}

	matcher = newTopicMatcherFromStrings(test.textual...)
	for i, cond := range matcher.conditions {
		for topic := range cond {
			if _, ok := test.matcher[i][topic]; !ok {
				t.Errorf("condition %d; extra topic found: 0x%x", i, topic[:])
			}
		}
	}
	for i, cond := range test.matcher {
		for topic := range cond {
			if _, ok := matcher.conditions[i][topic]; !ok {
				t.Errorf("condition %d; topic not found: 0x%x", i, topic[:])
			}
		}
	}
}

var topicMatcherTests = []struct {
	filter [][]string
	topics []string
	match  bool
}{
	// Empty topic matcher should match everything
	{
		filter: [][]string{},
		topics: []string{},
		match:  true,
	},
	{
		filter: [][]string{},
		topics: []string{"a", "b", "c"},
		match:  true,
	},
	// Fixed topic matcher should match strictly, but only prefix
	{
		filter: [][]string{{"a"}, {"b"}},
		topics: []string{"a"},
		match:  false,
	},
	{
		filter: [][]string{{"a"}, {"b"}},
		topics: []string{"a", "b"},
		match:  true,
	},
	{
		filter: [][]string{{"a"}, {"b"}},
		topics: []string{"a", "b", "c"},
		match:  true,
	},
	// Multi-matcher should match any from a sub-group
	{
		filter: [][]string{{"a1", "a2"}},
		topics: []string{"a"},
		match:  false,
	},
	{
		filter: [][]string{{"a1", "a2"}},
		topics: []string{"a1"},
		match:  true,
	},
	{
		filter: [][]string{{"a1", "a2"}},
		topics: []string{"a2"},
		match:  true,
	},
	// Wild-card condition should match anything
	{
		filter: [][]string{{}, {"b"}},
		topics: []string{"a"},
		match:  false,
	},
	{
		filter: [][]string{{}, {"b"}},
		topics: []string{"a", "b"},
		match:  true,
	},
	{
		filter: [][]string{{}, {"b"}},
		topics: []string{"b", "b"},
		match:  true,
	},
}

func TestTopicMatcher(t *testing.T) {
	for i, tt := range topicMatcherTests {
		topics := NewTopicsFromStrings(tt.topics...)

		matcher := newTopicMatcherFromStrings(tt.filter...)
		if match := matcher.Matches(topics); match != tt.match {
			t.Errorf("test %d: match mismatch: have %v, want %v", i, match, tt.match)
		}
	}
}
