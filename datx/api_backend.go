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

package datx

import (
	"context"
	"math/big"

	"github.com/DATx-Protocol/go-DATx/accounts"
	"github.com/DATx-Protocol/go-DATx/common"
	"github.com/DATx-Protocol/go-DATx/common/math"
	"github.com/DATx-Protocol/go-DATx/core"
	"github.com/DATx-Protocol/go-DATx/core/bloombits"
	"github.com/DATx-Protocol/go-DATx/core/state"
	"github.com/DATx-Protocol/go-DATx/core/types"
	"github.com/DATx-Protocol/go-DATx/core/vm"
	"github.com/DATx-Protocol/go-DATx/datx/downloader"
	"github.com/DATx-Protocol/go-DATx/datx/gasprice"
	"github.com/DATx-Protocol/go-DATx/datxdb"
	"github.com/DATx-Protocol/go-DATx/event"
	"github.com/DATx-Protocol/go-DATx/params"
	"github.com/DATx-Protocol/go-DATx/rpc"
)

// EthApiBackend implements ethapi.Backend for full nodes
type EthApiBackend struct {
	datx *Ethereum
	gpo *gasprice.Oracle
}

func (b *EthApiBackend) ChainConfig() *params.ChainConfig {
	return b.datx.chainConfig
}

func (b *EthApiBackend) CurrentBlock() *types.Block {
	return b.datx.blockchain.CurrentBlock()
}

func (b *EthApiBackend) SetHead(number uint64) {
	b.datx.protocolManager.downloader.Cancel()
	b.datx.blockchain.SetHead(number)
}

func (b *EthApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.datx.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.datx.blockchain.CurrentBlock().Header(), nil
	}
	return b.datx.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *EthApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.datx.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.datx.blockchain.CurrentBlock(), nil
	}
	return b.datx.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *EthApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.datx.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.datx.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *EthApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.datx.blockchain.GetBlockByHash(blockHash), nil
}

func (b *EthApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.datx.chainDb, blockHash, core.GetBlockNumber(b.datx.chainDb, blockHash)), nil
}

func (b *EthApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.datx.blockchain.GetTdByHash(blockHash)
}

func (b *EthApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.datx.BlockChain(), nil)
	return vm.NewEVM(context, state, b.datx.chainConfig, vmCfg), vmError, nil
}

func (b *EthApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.datx.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *EthApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.datx.BlockChain().SubscribeChainEvent(ch)
}

func (b *EthApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.datx.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *EthApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.datx.BlockChain().SubscribeLogsEvent(ch)
}

func (b *EthApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.datx.txPool.AddLocal(signedTx)
}

func (b *EthApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.datx.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *EthApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.datx.txPool.Get(hash)
}

func (b *EthApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.datx.txPool.State().GetNonce(addr), nil
}

func (b *EthApiBackend) Stats() (pending int, queued int) {
	return b.datx.txPool.Stats()
}

func (b *EthApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.datx.TxPool().Content()
}

func (b *EthApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.datx.TxPool().SubscribeTxPreEvent(ch)
}

func (b *EthApiBackend) Downloader() *downloader.Downloader {
	return b.datx.Downloader()
}

func (b *EthApiBackend) ProtocolVersion() int {
	return b.datx.EthVersion()
}

func (b *EthApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *EthApiBackend) ChainDb() datxdb.Database {
	return b.datx.ChainDb()
}

func (b *EthApiBackend) EventMux() *event.TypeMux {
	return b.datx.EventMux()
}

func (b *EthApiBackend) AccountManager() *accounts.Manager {
	return b.datx.AccountManager()
}

func (b *EthApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.datx.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *EthApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.datx.bloomRequests)
	}
}
