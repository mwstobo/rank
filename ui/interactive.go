package ui

import (
	"fmt"
	"github.com/mwstobo/rank/ranker"
	"github.com/mwstobo/rank/store"
	"github.com/mwstobo/rank/util"
)

type InteractiveApp struct {
	ranker  *ranker.Ranker
	storage store.Storage
	quit    chan bool
}

func NewInteractiveApp(
	ranker *ranker.Ranker,
	storage store.Storage) *InteractiveApp {

	return &InteractiveApp{
		ranker:  ranker,
		storage: storage,
		quit:    make(chan bool, 1),
	}
}

func (app *InteractiveApp) Run() {
	choices := []util.Choice{
		util.Choice{"Add an item", "1", app.AddAction},
		util.Choice{"View your list", "2", app.ListAction},
		util.Choice{"Delete an item", "3", app.DeleteAction},
		util.Choice{"Save and quit", "4", app.SaveAndQuitAction},
		util.Choice{"Quit without saving", "5", app.QuitAction},
	}

	for {
		select {
		case <-app.quit:
			return
		default:
			selectedChoice, err := util.VerbosePresentChoice(
				"What do you want to do?",
				choices)
			if err == util.EOF {
				app.SaveAndQuitAction()
				continue
			} else if err != nil {
				fmt.Printf("Error reading choice: %v\n", err)
				app.SaveAndQuitAction()
				continue
			}
			selectedChoice()
			fmt.Println("")
		}
	}
}

func (app *InteractiveApp) AddAction() {
	item, err := util.GetInput("What do you want to add?")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		app.SaveAndQuitAction()
		return
	}

	err = app.ranker.AddItem(item)
	if err != nil {
		fmt.Printf("Error adding item: %v\n", err)
		app.SaveAndQuitAction()
		return
	}
}

func (app *InteractiveApp) ListAction() {
	app.ranker.ListItems()
}

func (app *InteractiveApp) DeleteAction() {
	var itemNumber int
	var err error

	if app.ranker.Ranking.Length() == 0 {
		fmt.Println("No items")
		return
	}

	app.ListAction()
	for {
		itemNumber, err = util.GetInputInt("What do you want to delete?")
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			app.SaveAndQuitAction()
			return
		}

		if itemNumber > 0 && itemNumber <= app.ranker.Ranking.Length() {
			break
		}
		fmt.Printf("Invalid item number %d\n", itemNumber)
	}

	app.ranker.DeleteItem(itemNumber - 1)
}

func (app *InteractiveApp) SaveAndQuitAction() {
	app.storage.Export(app.ranker.Ranking)
	app.QuitAction()
}

func (app *InteractiveApp) QuitAction() {
	app.quit <- true
}
