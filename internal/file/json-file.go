package file

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"pokedex/internal/log"
)

func WriteFile(fileName string, objectToWrite interface{}) error {
	log.Debugf("Writing file: %s", fileName)

	f, err := os.Create(fileName)
	if err != nil {
		log.Error(err)
		return err
	}
	defer f.Close()

	result, err := json.MarshalIndent(objectToWrite, "", "  ")
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = f.WriteString(string(result))
	if err != nil {
		log.Error(err)
		return err
	}

	if err := f.Sync(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func ReadFile(fileName string, objectToWrite interface{}) error {
	log.Debugf("Reading file: %s", fileName)

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := json.Unmarshal(dat, objectToWrite); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
