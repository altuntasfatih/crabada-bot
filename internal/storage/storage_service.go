package storage

import (
	"crabada-bot/pkg/constant"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	NoExpiration     time.Duration = -1
	ExpireInOneHour                = 1 * time.Hour
	CleanupInOneHour               = 1 * time.Hour
	filePath         string        = "cache"
)

type ReadWriter interface {
	Read(path string) (error, []byte)
	Write(path string, value []byte) error
}

type Item struct {
	BattlePoint    int `json:"battlePoint"`
	MinePoint      int `json:"minePoint"`
	ReinforceCount int `json:"reinforceCount"`
}

type StorageService struct {
	cache SetGetFlusher
}

func NewDefaultStorageService(cache SetGetFlusher) StorageService {
	return StorageService{cache: cache}
}

func (s *StorageService) Get(key string) (error, Item) {
	log.Info().Msgf("StorageService.Get is called getting value for key [%s]", key)
	value, found := s.cache.Get(key)
	if !found {
		return errors.New(fmt.Sprintf("value not found for key [%s]\n\t", key)), Item{}
	}
	return nil, value.(Item)
}

func (s *StorageService) Set(key string, value interface{}) {
	log.Info().Msgf("StorageService.Set is called saving key: %s, value: %v", key, value)
	s.cache.Set(key, value, ExpireInOneHour)
}

func (s *StorageService) Flush() {
	fmt.Printf("%v StorageService.Flush is called \n", time.Now())
	s.cache.Flush()
}

func (s *StorageService) Persist() error {
	items := s.cache.Items()
	log.Info().Msgf("StorageService.Persist:Persist %d entries to file", len(s.cache.Items()))
	jsonString, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = s.Write(filePath, jsonString)
	if err != nil {
		return errors.Wrap(err, "error while persisting storage")
	}
	return nil
}

func (s *StorageService) Load() error {
	err, jsonString := s.Read(filePath)
	if err != nil {
		return err
	}

	var items map[string]cache.Item
	err = json.Unmarshal(jsonString, &items)
	if err != nil {
		return err
	}

	log.Info().Msgf("StorageService.Load: %d entries are loaded from file", len(items))

	for key, item := range items {
		_, found := s.cache.Get(key)
		if found {
			continue
		}
		ob := item.Object
		tempItem := Item{
			BattlePoint:    int(ob.(map[string]interface{})["battlePoint"].(float64)),
			MinePoint:      int(ob.(map[string]interface{})["minePoint"].(float64)),
			ReinforceCount: int(ob.(map[string]interface{})["reinforceCount"].(float64)),
		}
		s.cache.Set(key, tempItem, ExpireInOneHour)
	}
	return nil
}

func (s *StorageService) Read(path string) (error, []byte) {
	log.Info().Msgf("filePersistenceService.Read is called")
	homeDirectory, _ := os.UserHomeDir()

	filename := filepath.Join(homeDirectory, path, "latest-data.json")

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err, nil
	}
	return nil, bytes
}

func (s *StorageService) Write(path string, value []byte) error {
	log.Info().Msgf("filePersistenceService.Write is called")
	homeDirectory, _ := os.UserHomeDir()

	filename := filepath.Join(homeDirectory, path, fmt.Sprintf("%s-data.json", time.Now().Format(constant.TimeLogFieldFormat)))
	err := ioutil.WriteFile(filename, value, 0644)
	if err != nil {
		return err
	}
	filename = filepath.Join(homeDirectory, path, "latest-data.json")
	err = ioutil.WriteFile(filename, value, 0644)
	if err != nil {
		return err
	}
	return nil
}