package ranker

import (
	"fmt"
	"github.com/mwstobo/rank/util"
)

type Ranker struct {
	Ranking Ranking
}

func NewRanker(ranking Ranking) *Ranker {
	return &Ranker{
		Ranking: ranking,
	}
}

func (ranker *Ranker) AddItem(item string) error {
	var baseIndex, middle int
	var comparativeItem string

	ranking := make([]string, len(ranker.Ranking))

	copy(ranking, ranker.Ranking)
	baseIndex = 0

	for {
		middle = len(ranking) / 2
		if len(ranking) == 0 {
			break
		}

		comparativeItem = ranking[middle]

		isHigher := func() {
			ranking = ranking[:middle]
		}

		isLower := func() {
			ranking = ranking[middle+1:]
			baseIndex += middle + 1
		}

		choices := []util.Choice{
			util.Choice{Command: "y", Action: isHigher},
			util.Choice{Command: "n", Action: isLower},
		}

		selectedChoice, err := util.PresentChoice(
			fmt.Sprintf("Is %s higher than %s?", item, comparativeItem),
			choices)
		if err != nil {
			return err
		}
		selectedChoice()
	}

	insertIndex := baseIndex + middle
	ranker.Ranking = append(ranker.Ranking, "")
	copy(ranker.Ranking[insertIndex+1:], ranker.Ranking[insertIndex:])
	ranker.Ranking[insertIndex] = item

	return nil
}

func (ranker *Ranker) ListItems() {
	if len(ranker.Ranking) == 0 {
		fmt.Println("No items")
		return
	}

	for position, item := range ranker.Ranking {
		fmt.Printf("%d. %s\n", position+1, item)
	}
}

func (ranker *Ranker) DeleteItem(itemIndex int) {
	ranker.Ranking = append(
		ranker.Ranking[:itemIndex],
		ranker.Ranking[itemIndex+1:]...)
}
