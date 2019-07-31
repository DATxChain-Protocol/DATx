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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/DATx-Protocol/go-DATx/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("datx/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("datx/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("datx/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("datx/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("datx/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("datx/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("datx/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("datx/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("datx/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("datx/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("datx/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("datx/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("datx/downloader/states/in")
	stateDropMeter = metrics.NewMeter("datx/downloader/states/drop")
)
