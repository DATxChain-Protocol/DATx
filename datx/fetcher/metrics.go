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

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/DATx-Protocol/go-DATx/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("datx/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("datx/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("datx/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("datx/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("datx/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("datx/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("datx/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("datx/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("datx/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("datx/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("datx/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("datx/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("datx/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("datx/fetcher/filter/bodies/out")
)
