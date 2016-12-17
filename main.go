package main

import (
	"fmt"
	"github.com/mwstobo/rank/interactive"
	"github.com/mwstobo/rank/ranker"
	"github.com/mwstobo/rank/store"
	"os"
)

func main() {
	var filename, command string
	switch len(os.Args) {
	case 3:
		command = os.Args[2]
		fallthrough
	case 2:
		filename = os.Args[1]
	default:
		usage()
	}

	storage := store.NewJsonStorage(filename)

	ranking, err := storage.Import()
	if err != nil {
		fmt.Printf("Error importing ranking file: %v\n", err)
		os.Exit(1)
	}

	ranker := ranker.NewRanker(ranking)

	if command == "" {
		app := interactive.NewInteractiveApp(ranker, storage)
		app.Run()
	} else {
		fmt.Printf("Command: %s\n", command)
	}
}

func usage() {
	fmt.Println("Invalid input!")
	os.Exit(1)
}
