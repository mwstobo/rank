package main

import (
	"fmt"
	"github.com/mwstobo/rank/config"
	"github.com/mwstobo/rank/ranker"
	"github.com/mwstobo/rank/rankings"
	"github.com/mwstobo/rank/store"
	"github.com/mwstobo/rank/ui"
	"os"
)

func main() {
	err := config.ParseConfig()
	if err == config.Help {
		os.Exit(0)
	} else if err != nil {
		fmt.Printf("Error parsing config: %v\n", err)
		os.Exit(1)
	}

	storage := store.NewJsonStorage(config.Filename)
	rankingMap, err := storage.Import()
	if err != nil {
		fmt.Printf("Error importing ranking file: %v\n", err)
		os.Exit(1)
	}

	ranking := rankings.NewArrayRanking(rankingMap)
	ranker := ranker.NewRanker(ranking)

	if config.Ui == config.INTERACTIVE {
		app := ui.NewInteractiveApp(ranker, storage)
		app.Run()
	} else {
		app := ui.NewCliApp(ranker, storage)
		switch config.Command {
		case config.ADD_COMMAND:
			app.AddAction(config.AddItem)
		case config.LIST_COMMAND:
			app.ListAction()
		case config.DELETE_COMMAND:
			app.DeleteAction(config.DeleteItemNumber)
		default:
			fmt.Printf("Config has bad command %s!", config.Command)
		}
		app.SaveAction()
	}
}
