// Copyright 2015 The go-datx Authors
// This file is part of the go-datx library.
//
// The go-datx library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-datx library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-datx library. If not, see <http://www.gnu.org/licenses/>.

package gasprice

import (
	"context"
	"math/big"
	"sort"
	"sync"

	"github.com/KunkaYU/go-DATx/common"
	"github.com/KunkaYU/go-DATx/internal/ethapi"
	"github.com/KunkaYU/go-DATx/params"
	"github.com/KunkaYU/go-DATx/rpc"
)

var maxPrice = big.NewInt(500 * params.Shannon)

type Config struct {
	Blocks     int
	Percentile int
	Default    *big.Int `toml:",omitempty"`
}

// Oracle recommends gas prices based on the content of recent
// blocks. Suitable for both light and full clients.
type Oracle struct {
	backend   ethapi.Backend
	lastHead  common.Hash
	lastPrice *big.Int
	cacheLock sync.RWMutex
	fetchLock sync.Mutex

	checkBlocks, maxEmpty, maxBlocks int
	percentile                       int
}

// NewOracle returns a new oracle.
func NewOracle(backend ethapi.Backend, params Config) *Oracle {
	blocks := params.Blocks
	if blocks < 1 {
		blocks = 1
	}
	percent := params.Percentile
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	return &Oracle{
		backend:     backend,
		lastPrice:   params.Default,
		checkBlocks: blocks,
		maxEmpty:    blocks / 2,
		maxBlocks:   blocks * 5,
		percentile:  percent,
	}
}

// SuggestPrice returns the recommended gas price.
func (gpo *Oracle) SuggestPrice(ctx context.Context) (*big.Int, error) {
	gpo.cacheLock.RLock()
	lastHead := gpo.lastHead
	lastPrice := gpo.lastPrice
	gpo.cacheLock.RUnlock()

	head, _ := gpo.backend.HeaderByNumber(ctx, rpc.LatestBlockNumber)
	headHash := head.Hash()
	if headHash == lastHead {
		return lastPrice, nil
	}

	gpo.fetchLock.Lock()
	defer gpo.fetchLock.Unlock()

	// try checking the cache again, maybe the last fetch fetched what we need
	gpo.cacheLock.RLock()
	lastHead = gpo.lastHead
	lastPrice = gpo.lastPrice
	gpo.cacheLock.RUnlock()
	if headHash == lastHead {
		return lastPrice, nil
	}

	blockNum := head.Number.Uint64()
	ch := make(chan getBlockPricesResult, gpo.checkBlocks)
	sent := 0
	exp := 0
	var txPrices []*big.Int
	for sent < gpo.checkBlocks && blockNum > 0 {
		go gpo.getBlockPrices(ctx, blockNum, ch)
		sent++
		exp++
		blockNum--
	}
	maxEmpty := gpo.maxEmpty
	for exp > 0 {
		res := <-ch
		if res.err != nil {
			return lastPrice, res.err
		}
		exp--
		if len(res.prices) > 0 {
			txPrices = append(txPrices, res.prices...)
			continue
		}
		if maxEmpty > 0 {
			maxEmpty--
			continue
		}
		if blockNum > 0 && sent < gpo.maxBlocks {
			go gpo.getBlockPrices(ctx, blockNum, ch)
			sent++
			exp++
			blockNum--
		}
	}
	price := lastPrice
	if len(txPrices) > 0 {
		sort.Sort(bigIntArray(txPrices))
		price = txPrices[(len(txPrices)-1)*gpo.percentile/100]
	}
	if price.Cmp(maxPrice) > 0 {
		price = new(big.Int).Set(maxPrice)
	}

	gpo.cacheLock.Lock()
	gpo.lastHead = headHash
	gpo.lastPrice = price
	gpo.cacheLock.Unlock()
	return price, nil
}

type getBlockPricesResult struct {
	prices []*big.Int
	err    error
}

// getLowestPrice calculates the lowest transaction gas price in a given block
// and sends it to the result channel. If the block is empty, price is nil.
func (gpo *Oracle) getBlockPrices(ctx context.Context, blockNum uint64, ch chan getBlockPricesResult) {
	block, err := gpo.backend.BlockByNumber(ctx, rpc.BlockNumber(blockNum))
	if block == nil {
		ch <- getBlockPricesResult{nil, err}
		return
	}
	txs := block.Transactions()
	prices := make([]*big.Int, len(txs))
	for i, tx := range txs {
		prices[i] = tx.GasPrice()
	}
	ch <- getBlockPricesResult{prices, nil}
}

type bigIntArray []*big.Int

func (s bigIntArray) Len() int           { return len(s) }
func (s bigIntArray) Less(i, j int) bool { return s[i].Cmp(s[j]) < 0 }
func (s bigIntArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
