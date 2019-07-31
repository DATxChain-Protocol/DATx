// Copyright 2017 The go-DATx Authors
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

package dpos

import (
	"github.com/DATx-Protocol/go-DATx/common"
	"github.com/DATx-Protocol/go-DATx/consensus"
	"github.com/DATx-Protocol/go-DATx/core/types"
	"github.com/DATx-Protocol/go-DATx/rpc"

	"math/big"
)

// API is a user facing RPC API to allow controlling the delegate and voting
// mechanisms of the delegated-proof-of-stake
type API struct {
	chain consensus.ChainReader
	dpos  *Dpos
}

// GetValidators retrieves the list of the validators at specified block
func (api *API) GetValidators(number *rpc.BlockNumber) ([]common.Address, error) {
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	if header == nil {
		return nil, errUnknownBlock
	}

	epochTrie, err := types.NewEpochTrie(header.DposContext.EpochHash, api.dpos.db)
	if err != nil {
		return nil, err
	}
	dposContext := types.DposContext{}
	dposContext.SetEpoch(epochTrie)
	validators, err := dposContext.GetValidators()
	if err != nil {
		return nil, err
	}
	return validators, nil
}

// GetConfirmedBlockNumber retrieves the latest irreversible block
func (api *API) GetConfirmedBlockNumber() (*big.Int, error) {
	var err error
	header := api.dpos.confirmedBlockHeader
	if header == nil {
		header, err = api.dpos.loadConfirmedBlockHeader(api.chain)
		if err != nil {
			return nil, err
		}
	}
	return header.Number, nil
}
