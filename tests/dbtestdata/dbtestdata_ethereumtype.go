package dbtestdata

import (
	"encoding/hex"
	"github.com/grupokindynos/coins-explorer/bchain"
)

// Addresses
const (
	EthAddr3e         = "3e3a3d69dc66ba10737f531ed088954a9ec89d97"
	EthAddr55         = "555ee11fbddc0e49a9bab358a8941ad95ffdb48f"
	EthAddr20         = "20cd153de35d469ba46127a0c8f18626b59a256a"
	EthAddr9f         = "9f4981531fda132e83c44680787dfa7ee31e4f8d"
	EthAddr4b         = "4bda106325c335df99eab7fe363cac8a0ba2a24d"
	EthAddr7b         = "7b62eb7fe80350dc7ec945c0b73242cb9877fb1b"
	EthAddrContract4a = "4af4114f73d1c1c903ac9e0361b379d1291808a2" // ERC-20 (VTY)
	EthAddrContract0d = "0d0f936ee4c93e25944694d6c121de94d9760f11" // ERC-20 (MTT)
	EthAddrContract47 = "479cc461fecd078f766ecc58533d6f69580cf3ac" // non ERC20

	EthTxidB1T1  = "cd647151552b5132b2aef7c9be00dc6f73afc5901dde157aab131335baaa853b"
	EthTx1Packed = "08e8dd870210a6a6f0db051a6908ece40212050430e234001888a40122081bc0159d530e60003220cd647151552b5132b2aef7c9be00dc6f73afc5901dde157aab131335baaa853b3a14555ee11fbddc0e49a9bab358a8941ad95ffdb48f42143e3a3d69dc66ba10737f531ed088954a9ec89d97480a22070a025208120101"
	EthTxidB1T2  = "a9cd088aba2131000da6f38a33c20169baee476218deea6b78720700b895b101"
	EthTx2Packed = "08e8dd870210a6a6f0db051aa20108d001120509502f900018d5e1042a44a9059cbb000000000000000000000000555ee11fbddc0e49a9bab358a8941ad95ffdb48f00000000000000000000000000000000000000000000021e19e0c9bab24000003220a9cd088aba2131000da6f38a33c20169baee476218deea6b78720700b895b1013a144af4114f73d1c1c903ac9e0361b379d1291808a2421420cd153de35d469ba46127a0c8f18626b59a256a22a8010a02cb391201011a9e010a144af4114f73d1c1c903ac9e0361b379d1291808a2122000000000000000000000000000000000000000000000021e19e0c9bab24000001a20ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef1a2000000000000000000000000020cd153de35d469ba46127a0c8f18626b59a256a1a20000000000000000000000000555ee11fbddc0e49a9bab358a8941ad95ffdb48f"
	EthTxidB2T1  = "c2c3dd1ecb00e8a6d81f793d24387cf2947a313e94ab03b1fb22cd63320f6c91"
	EthTx3Packed = "08e9dd870210d4b5f0db051a6708c20112050218711a001888a401220710bc3578bd37d83220c2c3dd1ecb00e8a6d81f793d24387cf2947a313e94ab03b1fb22cd63320f6c913a149f4981531fda132e83c44680787dfa7ee31e4f8d4214555ee11fbddc0e49a9bab358a8941ad95ffdb48f480722070a025208120101"
	EthTxidB2T2  = "c92919ad24ffd58f760b18df7949f06e1190cf54a50a0e3745a385608ed3cbf2"
	EthTx4Packed = "08e9dd870210d4b5f0db051aa50b08f6be0712043b9aca001890a10f2ac40a4f15078700000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000003c00000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000000000000000000000000000000048000000000000000000000000000000000000000000000000000000000000004e00000000000000000000000000000000000000000000000000000000000000002000000000000000000000000555ee11fbddc0e49a9bab358a8941ad95ffdb48f0000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d0000000000000000000000000d0f936ee4c93e25944694d6c121de94d9760f110000000000000000000000004af4114f73d1c1c903ac9e0361b379d1291808a200000000000000000000000000000000000000000000000000000000000000000000000000000000000000007b62eb7fe80350dc7ec945c0b73242cb9877fb1b0000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d0000000000000000000000004af4114f73d1c1c903ac9e0361b379d1291808a20000000000000000000000000d0f936ee4c93e25944694d6c121de94d9760f110000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000a5ef5a7656bfb0000000000000000000000000000000000000000000000000000004ba78398d5c5000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000166cfe0b9579b4ecf7a2801880f644009a324671a79754ea57c3a103c6e70d3dbef6ba69a08000000000000000000000000000000000000000000000000004f937d86afb90000000000000000000000000000000000000000000000000ab280fd8037d500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000166cfb784b7c1f3fbe8b75484603ab8adc58aaee3a46245a6579fac7077b5570018b4e0d4eb0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000308fd0e798ac00000000000000000000000000000000000000000000000006a8313d60b1f80000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001b000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000029de0ccec59e8948e3d905b40e5542335ebc1eb4674db517d2f6392ec7fdeb3d45f3449d313ee2589819c6c79eb1c1b047adae68565c1608e3a1d1d70823febb0000000000000000000000000000000000000000000000000000000000000000234d06fe17f1202e8b07177a30eb64d14adc08cdb3fa1b3e3e0bea0f9672c02175b77c01c51d3c7e460723b27ecbc7801fd6482559a8c9999593f9a4d149c73843220c92919ad24ffd58f760b18df7949f06e1190cf54a50a0e3745a385608ed3cbf23a14479cc461fecd078f766ecc58533d6f69580cf3ac42144bda106325c335df99eab7fe363cac8a0ba2a24d482422d40b0a03034d301201011a9e010a140d0f936ee4c93e25944694d6c121de94d9760f1112200000000000000000000000000000000000000000000000006a8313d60b1f606b1a20ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef1a20000000000000000000000000555ee11fbddc0e49a9bab358a8941ad95ffdb48f1a200000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d1a9e010a144af4114f73d1c1c903ac9e0361b379d1291808a21220000000000000000000000000000000000000000000000000000308fd0e798ac01a20ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef1a200000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d1a20000000000000000000000000555ee11fbddc0e49a9bab358a8941ad95ffdb48f1aa1030a14479cc461fecd078f766ecc58533d6f69580cf3ac1280020000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d0000000000000000000000000d0f936ee4c93e25944694d6c121de94d9760f110000000000000000000000004af4114f73d1c1c903ac9e0361b379d1291808a20000000000000000000000000000000000000000000000006a8313d60b1f606b000000000000000000000000000000000000000000000000000308fd0e798ac0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005e083a16f4b092c5729a49f9c3ed3cc171bb3d3d0c22e20b1de6063c32f399ac1a200d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb31a20000000000000000000000000555ee11fbddc0e49a9bab358a8941ad95ffdb48f1a2000000000000000000000000000000000000000000000000000000000000000001a205af266c0a89a07c1917deaa024414577e6c3c31c8907d079e13eb448c082594f1a9e010a144af4114f73d1c1c903ac9e0361b379d1291808a2122000000000000000000000000000000000000000000000000000031855667df7a81a20ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef1a200000000000000000000000007b62eb7fe80350dc7ec945c0b73242cb9877fb1b1a200000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d1a9e010a140d0f936ee4c93e25944694d6c121de94d9760f1112200000000000000000000000000000000000000000000000006a8313d60b1f80001a20ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef1a200000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d1a200000000000000000000000007b62eb7fe80350dc7ec945c0b73242cb9877fb1b1aa1030a14479cc461fecd078f766ecc58533d6f69580cf3ac1280020000000000000000000000004bda106325c335df99eab7fe363cac8a0ba2a24d0000000000000000000000004af4114f73d1c1c903ac9e0361b379d1291808a20000000000000000000000000d0f936ee4c93e25944694d6c121de94d9760f1100000000000000000000000000000000000000000000000000031855667df7a80000000000000000000000000000000000000000000000006a8313d60b1f800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f2b0d62c44ed08f2a5adef40c875d20310a42a9d4f488bd26323256fe01c7f481a200d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb31a200000000000000000000000007b62eb7fe80350dc7ec945c0b73242cb9877fb1b1a2000000000000000000000000000000000000000000000000000000000000000001a20b0b69dad58df6032c3b266e19b1045b19c87acd2c06fb0c598090f44b8e263aa"
)

func unpackTxs(packed []string, parser bchain.BlockChainParser) []bchain.Tx {
	r := make([]bchain.Tx, len(packed))
	for i, p := range packed {
		b, err := hex.DecodeString(p)
		if err != nil {
			panic(err)
		}
		tx, _, err := parser.UnpackTx(b)
		if err != nil {
			panic(err)
		}
		r[i] = *tx
	}
	return r
}

// GetTestEthereumTypeBlock1 returns block #1
func GetTestEthereumTypeBlock1(parser bchain.BlockChainParser) *bchain.Block {
	return &bchain.Block{
		BlockHeader: bchain.BlockHeader{
			Height:        4321000,
			Hash:          "0xc7b98df95acfd11c51ba25611a39e004fe56c8fdfc1582af99354fcd09c17b11",
			Size:          31839,
			Time:          1534858022,
			Confirmations: 2,
		},
		Txs: unpackTxs([]string{EthTx1Packed, EthTx2Packed}, parser),
	}
}

// GetTestEthereumTypeBlock2 returns block #2
func GetTestEthereumTypeBlock2(parser bchain.BlockChainParser) *bchain.Block {
	return &bchain.Block{
		BlockHeader: bchain.BlockHeader{
			Height:        4321001,
			Hash:          "0x2b57e15e93a0ed197417a34c2498b7187df79099572c04a6b6e6ff418f74e6ee",
			Size:          2345678,
			Time:          1534859988,
			Confirmations: 1,
		},
		Txs: unpackTxs([]string{EthTx3Packed, EthTx4Packed}, parser),
	}
}
