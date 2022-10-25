package main

import (
	"chain-api-imgo/setting"
	"chain-api-imgo/unittest"
	"flag"
	"log"
	_ "net/http/pprof"
)

func main() {
	var args = unittest.Args{}
	flag.StringVar(&args.Env, "env", "localtest", "config file, dev or prod")
	flag.IntVar(&args.N, "n", 5, "启动n个协程")
	flag.IntVar(&args.Loop, "loop", 1000, "每个协程循环创建loop个合约")
	flag.IntVar(&args.Tx, "tx", 1000, "每个合约调用tx次交易")
	flag.IntVar(&args.S, "s", 0, "sleep")
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	var err error
	err = setting.Setup(args.Env)
	if err != nil {
		log.Fatal(err)
	}

	//profileTest(&args)

	// Test 1
	// t := unittest.AllowAddress{
	// 	Abi:          "asserts/sol/MgtvNFTv2b1.abi",
	// 	Bin:          "asserts/sol/MgtvNFTv2b1.bin",
	// 	AddressFile:  "/Users/huyinghuan/sol-workspace/mgtv-nft/addresses.txt",
	// 	StartLine:    100000,
	// 	EndLine:      200000,
	// 	ContractName: "profile_KTwrecDbYacLsUIehTrCISQd",
	// }
	// if err := t.Init(); err != nil {
	// 	log.Println(err)
	// }
	// if err := t.Run(&args); err != nil {
	// 	log.Println(err)
	// }

	// Test 2
	// t := unittest.CrossChain{
	// 	Abi:          "asserts/sol/CrosschainV2.abi",
	// 	Bin:          "asserts/sol/CrosschainV2.bin",
	// 	ContractName: "profile_QBLUMvxlCuRxsNXYvviKUQEk",
	// }
	t := unittest.BugCheckt{
		// Abi:          "asserts/sol/contracts_bug_sol_BugContract.abi",
		// Bin:          "asserts/sol/contracts_bug_sol_BugContract.bin",
		Abi:          "asserts/sol/BugContract.abi",
		Bin:          "asserts/sol/BugContract.bin",
		ContractName: "",
	}
	if err := t.Init(); err != nil {
		log.Println(err)
	}
	if err := t.Run(&args); err != nil {
		log.Println(err)
	}
}
