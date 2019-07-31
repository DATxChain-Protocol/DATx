// Copyright 2016 The go-DATx Authors
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

package datxclient

import "github.com/DATx-Protocol/go-DATx"

// Verify that Client implements the DATx interfaces.
var (
	_ = DATx.ChainReader(&Client{})
	_ = DATx.TransactionReader(&Client{})
	_ = DATx.ChainStateReader(&Client{})
	_ = DATx.ChainSyncReader(&Client{})
	_ = DATx.ContractCaller(&Client{})
	_ = DATx.GasEstimator(&Client{})
	_ = DATx.GasPricer(&Client{})
	_ = DATx.LogFilterer(&Client{})
	_ = DATx.PendingStateReader(&Client{})
	// _ = DATx.PendingStateEventer(&Client{})
	_ = DATx.PendingContractCaller(&Client{})
)
