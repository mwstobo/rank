package store

import (
	"encoding/json"
	"github.com/mwstobo/rank/ranker"
	"io/ioutil"
	"os"
)

type RankingJson struct {
	Ranking ranker.Ranking `json:"ranking"`
}

type JsonStorage struct {
	filename string
}

func NewJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{
		filename: filename,
	}
}

func (storage *JsonStorage) Import() (ranker.Ranking, error) {
	rankingData, err := ioutil.ReadFile(storage.filename)
	if os.IsNotExist(err) {
		return ranker.Ranking{}, nil
	} else if err != nil {
		return nil, err
	}

	rankingJson := &RankingJson{}
	err = json.Unmarshal(rankingData, rankingJson)
	if err != nil {
		return nil, err
	}

	if rankingJson.Ranking == nil {
		return ranker.Ranking{}, nil
	}
	return rankingJson.Ranking, nil
}

func (storage *JsonStorage) Export(ranking ranker.Ranking) error {
	rankingJson := RankingJson{ranking}
	rankingData, err := json.Marshal(rankingJson)
	rankingData = append(rankingData, '\n')
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(storage.filename, rankingData, 0644)
	if err != nil {
		return err
	}

	return nil
}
