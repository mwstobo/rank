package store

import (
	"github.com/mwstobo/rank/rankings"
)

type Storage interface {
	Import() (map[int]string, error)
	Export(rankings.Ranking) error
}
