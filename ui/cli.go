package ui

import (
	"fmt"
	"github.com/mwstobo/rank/ranker"
	"github.com/mwstobo/rank/store"
)

type CliApp struct {
	ranker  *ranker.Ranker
	storage store.Storage
}

func NewCliApp(
	ranker *ranker.Ranker,
	storage store.Storage) *CliApp {

	return &CliApp{
		ranker:  ranker,
		storage: storage,
	}
}

func (app *CliApp) AddAction(item string) {
	err := app.ranker.AddItem(item)
	if err != nil {
		fmt.Printf("Error adding item: %v\n", err)
		return
	}
}

func (app *CliApp) ListAction() {
	app.ranker.ListItems()
}

func (app *CliApp) DeleteAction(itemNumber int) {
	if len(app.ranker.Ranking) == 0 {
		fmt.Println("No items")
		return
	}

	app.ranker.DeleteItem(itemNumber - 1)
}

func (app *CliApp) SaveAction() {
	app.storage.Export(app.ranker.Ranking)
}
