package main

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qbase/server"
	"github.com/QOSGroup/qos/app"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	cmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"io"
)

func main() {
	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "qosd",
		Short:             "qos Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	// tendermint testnet cmd
	rootCmd.AddCommand(cmd.TestnetFilesCmd)
	rootCmd.AddCommand(bcli.LineBreak)

	server.AddCommands(ctx, cdc, rootCmd, app.QOSAppInit(),
		server.ConstructAppCreator(newApp, "qos"))

	executor := cli.PrepareBaseCmd(rootCmd, "qos", app.DefaultNodeHome)

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, storeTracer io.Writer) abci.Application {
	return app.NewApp(logger, db, storeTracer)
}
