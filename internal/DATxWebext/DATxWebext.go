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

// package DATxWebext contains gdatx specific DATxWeb.js extensions.
package DATxWebext

var Modules = map[string]string{
	"admin":      Admin_JS,
	"chequebook": Chequebook_JS,
	"clique":     Clique_JS,
	"debug":      Debug_JS,
	"datx":       Eth_JS,
	"miner":      Miner_JS,
	"net":        Net_JS,
	"personal":   Personal_JS,
	"rpc":        RPC_JS,
	"shh":        Shh_JS,
	"swarmfs":    SWARMFS_JS,
	"txpool":     TxPool_JS,
	"dpos":       Dpos_JS,
}

const Chequebook_JS = `
DATxWeb._extend({
	property: 'chequebook',
	methods: [
		new DATxWeb._extend.Method({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: DATxWeb._extend.utils.toDecimal
		}),
		new DATxWeb._extend.Method({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Method({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const Dpos_JS = `
DATxWeb._extend({
	property: 'dpos',
	methods: [
		new DATxWeb._extend.Method({
			name: 'getValidators',
			call: 'dpos_getValidators',
			params: 1,
			inputFormatter: [DATxWeb._extend.formatters.inputBlockNumberFormatter]
		}),
		new DATxWeb._extend.Method({
			name: 'getConfirmedBlockNumber',
			call: 'dpos_getConfirmedBlockNumber',
			params: 0,
			outputFormatter: DATxWeb._extend.utils.toBigNumber
		}),
	]
});
`

const Clique_JS = `
DATxWeb._extend({
	property: 'clique',
	methods: [
		new DATxWeb._extend.Method({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Method({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
	],
	properties: [
		new DATxWeb._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const Admin_JS = `
DATxWeb._extend({
	property: 'admin',
	methods: [
		new DATxWeb._extend.Method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new DATxWeb._extend.Method({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new DATxWeb._extend.Method({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new DATxWeb._extend.Method({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
	],
	properties: [
		new DATxWeb._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new DATxWeb._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new DATxWeb._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const Debug_JS = `
DATxWeb._extend({
	property: 'debug',
	methods: [
		new DATxWeb._extend.Method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'getBlockRlp',
			call: 'debug_getBlockRlp',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			params: 1,
			outputFormatter: console.log
		}),
		new DATxWeb._extend.Method({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new DATxWeb._extend.Method({
			name: 'metrics',
			call: 'debug_metrics',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new DATxWeb._extend.Method({
			name: 'stacks',
			call: 'debug_stacks',
			params: 0,
			outputFormatter: console.log
		}),
		new DATxWeb._extend.Method({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new DATxWeb._extend.Method({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new DATxWeb._extend.Method({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new DATxWeb._extend.Method({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new DATxWeb._extend.Method({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new DATxWeb._extend.Method({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new DATxWeb._extend.Method({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new DATxWeb._extend.Method({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Method({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new DATxWeb._extend.Method({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new DATxWeb._extend.Method({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new DATxWeb._extend.Method({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
	],
	properties: []
});
`

const Eth_JS = `
DATxWeb._extend({
	property: 'datx',
	methods: [
		new DATxWeb._extend.Method({
			name: 'sign',
			call: 'datx_sign',
			params: 2,
			inputFormatter: [DATxWeb._extend.formatters.inputAddressFormatter, null]
		}),
		new DATxWeb._extend.Method({
			name: 'resend',
			call: 'datx_resend',
			params: 3,
			inputFormatter: [DATxWeb._extend.formatters.inputTransactionFormatter, DATxWeb._extend.utils.fromDecimal, DATxWeb._extend.utils.fromDecimal]
		}),
		new DATxWeb._extend.Method({
			name: 'signTransaction',
			call: 'datx_signTransaction',
			params: 1,
			inputFormatter: [DATxWeb._extend.formatters.inputTransactionFormatter]
		}),
		new DATxWeb._extend.Method({
			name: 'submitTransaction',
			call: 'datx_submitTransaction',
			params: 1,
			inputFormatter: [DATxWeb._extend.formatters.inputTransactionFormatter]
		}),
		new DATxWeb._extend.Method({
			name: 'getRawTransaction',
			call: 'datx_getRawTransactionByHash',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (DATxWeb._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'datx_getRawTransactionByBlockHashAndIndex' : 'datx_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [DATxWeb._extend.formatters.inputBlockNumberFormatter, DATxWeb._extend.utils.toHex]
		}),
	],
	properties: [
		new DATxWeb._extend.Property({
			name: 'pendingTransactions',
			getter: 'datx_pendingTransactions',
			outputFormatter: function(txs) {
				var formatted = [];
				for (var i = 0; i < txs.length; i++) {
					formatted.push(DATxWeb._extend.formatters.outputTransactionFormatter(txs[i]));
					formatted[i].blockHash = null;
				}
				return formatted;
			}
		}),
	]
});
`

const Miner_JS = `
DATxWeb._extend({
	property: 'miner',
	methods: [
		new DATxWeb._extend.Method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
		new DATxWeb._extend.Method({
			name: 'stop',
			call: 'miner_stop'
		}),
		new DATxWeb._extend.Method({
			name: 'setValidator',
			call: 'miner_setValidator',
			params: 1,
			inputFormatter: [DATxWeb._extend.formatters.inputAddressFormatter]
		}),
		new DATxWeb._extend.Method({
			name: 'setCoinbase',
			call: 'miner_setCoinbase',
			params: 1,
			inputFormatter: [DATxWeb._extend.formatters.inputAddressFormatter]
		}),
		new DATxWeb._extend.Method({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [DATxWeb._extend.utils.fromDecimal]
		}),
		new DATxWeb._extend.Method({
			name: 'getHashrate',
			call: 'miner_getHashrate'
		}),
	],
	properties: []
});
`

const Net_JS = `
DATxWeb._extend({
	property: 'net',
	methods: [],
	properties: [
		new DATxWeb._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const Personal_JS = `
DATxWeb._extend({
	property: 'personal',
	methods: [
		new DATxWeb._extend.Method({
			name: 'importRawKey',
			call: 'personal_importRawKey',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'sign',
			call: 'personal_sign',
			params: 3,
			inputFormatter: [null, DATxWeb._extend.formatters.inputAddressFormatter, null]
		}),
		new DATxWeb._extend.Method({
			name: 'ecRecover',
			call: 'personal_ecRecover',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
	],
	properties: [
		new DATxWeb._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RPC_JS = `
DATxWeb._extend({
	property: 'rpc',
	methods: [],
	properties: [
		new DATxWeb._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const Shh_JS = `
DATxWeb._extend({
	property: 'shh',
	methods: [
	],
	properties:
	[
		new DATxWeb._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: DATxWeb._extend.utils.toDecimal
		}),
		new DATxWeb._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SWARMFS_JS = `
DATxWeb._extend({
	property: 'swarmfs',
	methods:
	[
		new DATxWeb._extend.Method({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new DATxWeb._extend.Method({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new DATxWeb._extend.Method({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxPool_JS = `
DATxWeb._extend({
	property: 'txpool',
	methods: [],
	properties:
	[
		new DATxWeb._extend.Property({
			name: 'content',
			getter: 'txpool_content'
		}),
		new DATxWeb._extend.Property({
			name: 'inspect',
			getter: 'txpool_inspect'
		}),
		new DATxWeb._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.pending = DATxWeb._extend.utils.toDecimal(status.pending);
				status.queued = DATxWeb._extend.utils.toDecimal(status.queued);
				return status;
			}
		}),
	]
});
`
