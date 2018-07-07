package file

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"pokedex/internal/log"
)

func DownloadFile(filepath string, url string) error {
	log.Debugf("Downloading file: %s to %s", url, filepath)

	out, err := os.Create(filepath)
	if err != nil {
		log.Error(err)
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func StreamFileOut(filepath string) (*bytes.Buffer, error) {
	byteStream, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return bytes.NewBuffer(byteStream), nil
}
