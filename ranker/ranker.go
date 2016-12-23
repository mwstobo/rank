package ranker

import (
	"fmt"
	"github.com/mwstobo/rank/rankings"
	"github.com/mwstobo/rank/util"
)

type Ranker struct {
	Ranking rankings.Ranking
}

func NewRanker(ranking rankings.Ranking) *Ranker {
	return &Ranker{
		Ranking: ranking,
	}
}

func (ranker *Ranker) AddItem(item string) error {
	var baseIndex, rankingLength, middle int
	var comparativeItem string

	rankingLength = ranker.Ranking.Length()
	baseIndex = 0

	for {
		middle = rankingLength / 2
		if rankingLength == 0 {
			break
		}

		comparativeItem = ranker.Ranking.Select(baseIndex + middle)

		isHigher := func() {
			rankingLength -= rankingLength - middle
		}

		isLower := func() {
			rankingLength -= middle + 1
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
	ranker.Ranking.Insert(insertIndex, item)

	return nil
}

func (ranker *Ranker) ListItems() {
	if ranker.Ranking.Length() == 0 {
		fmt.Println("No items")
		return
	}

	for i := 0; i < ranker.Ranking.Length(); i += 1 {
		item := ranker.Ranking.Select(i)
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func (ranker *Ranker) DeleteItem(itemIndex int) {
	ranker.Ranking.Delete(itemIndex)
}
