package main

import (
	"encoding/json"
	"flag"
	"gate-zkmerkle-proof/prover/prover_server"
	"io/ioutil"

	"gate-zkmerkle-proof/prover/config"
	"gate-zkmerkle-proof/utils"
)

func main() {
	proverConfig := &config.Config{}
	content, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(content, proverConfig)
	if err != nil {
		panic(err.Error())
	}
	remotePasswdConfig := flag.String("remote_password_config", "", "fetch password from aws secretsmanager")
	rerun := flag.Bool("rerun", false, "flag which indicates rerun proof generation")
	flag.Parse()
	if *remotePasswdConfig != "" {
		s, err := utils.GetMysqlSource(proverConfig.MysqlDataSource, *remotePasswdConfig)
		if err != nil {
			panic(err.Error())
		}
		proverConfig.MysqlDataSource = s
	}
	prover := prover_server.NewProver(proverConfig)
	prover.Run(*rerun)
}
