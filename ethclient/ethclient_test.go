// Copyright 2016 The go-datx Authors
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

package ethclient

import "github.com/meitu/go-datx"

// Verify that Client implements the datx interfaces.
var (
	_ = datx.ChainReader(&Client{})
	_ = datx.TransactionReader(&Client{})
	_ = datx.ChainStateReader(&Client{})
	_ = datx.ChainSyncReader(&Client{})
	_ = datx.ContractCaller(&Client{})
	_ = datx.GasEstimator(&Client{})
	_ = datx.GasPricer(&Client{})
	_ = datx.LogFilterer(&Client{})
	_ = datx.PendingStateReader(&Client{})
	// _ = datx.PendingStateEventer(&Client{})
	_ = datx.PendingContractCaller(&Client{})
)