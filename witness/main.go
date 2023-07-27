package main

import (
	"encoding/json"
	"fmt"
	"gate-zkmerkle-proof/witness/src"
	zk_smt "github.com/gatechain/gate-zk-smt"
	"io/ioutil"

	"gate-zkmerkle-proof/utils"
	"gate-zkmerkle-proof/witness/config"
)

func main() {
	witnessConfig := loadConfig()
	accounts, cexAssetsInfo, accountTree := accountTree(witnessConfig)

	witnessService := src.NewWitness(accountTree, uint32(len(accounts)), accounts, cexAssetsInfo, witnessConfig)
	witnessService.Run()
	fmt.Println("witness service run finished...")
}

func loadConfig() *config.Config {
	witnessConfig := &config.Config{}
	jsonFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(fmt.Sprintf("load config err : %s", err.Error()))
	}
	err = json.Unmarshal(jsonFile, witnessConfig)
	if err != nil {
		panic(err.Error())
	}

	return witnessConfig
}

func accountTree(witnessConfig *config.Config) ([]utils.AccountInfo, []utils.CexAssetInfo, zk_smt.SparseMerkleTree) {
	accounts, cexAssetsInfo, err := utils.ReadUserAssets(witnessConfig.UserDataFile)
	fmt.Println("the user account total is", len(accounts))
	if err != nil {
		panic(err.Error())
	}
	accountTree, err := utils.NewAccountTree(witnessConfig.TreeDB.Driver, witnessConfig.TreeDB.Option.Addr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("the account tree init height is ", accountTree.LatestVersion())
	fmt.Printf("account tree root is %x\n", accountTree.Root())

	return accounts, cexAssetsInfo, accountTree
}
