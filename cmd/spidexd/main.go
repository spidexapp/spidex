package main

import (
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/spidexapp/spidex/app"
	cmdcfg "github.com/spidexapp/spidex/cmd/config"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()
	rootCmd, _ := NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "spidexd", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
