package main

import (
	"encoding/json"
	"fmt"
	"gate-zkmerkle-proof/client"
	"gate-zkmerkle-proof/config"
	"gate-zkmerkle-proof/global"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func main() {
	// load file
	global.Cfg = &config.Config{}
	jsonFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(fmt.Sprintf("load config err : %s", err.Error()))
	}
	err = json.Unmarshal(jsonFile, global.Cfg)
	if err != nil {
		panic(err.Error())
	}

	rootCmd := &cobra.Command{
		Use:   "gproof",
		Short: "Command line interface for interacting with gate-zkmerkle-proof",
	}

	rootCmd.AddCommand(
		client.KeygenCommand(),
		client.WitnessCommand(),
		client.ProverCommand(),
		client.UserProofCommand(),
		client.VerifyCommand(),
		client.ToolCommand(),
	)

	if err = rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
