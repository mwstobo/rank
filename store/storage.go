package store

import (
	"github.com/mwstobo/rank/ranker"
)

type Storage interface {
	Import() (ranker.Ranking, error)
	Export(ranker.Ranking) error
}
