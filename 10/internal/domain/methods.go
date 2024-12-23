package domain

import (
	"io"
	"net/http"
	"os"
)

func (d *Domain) Wget(url, destination string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		d.log.Error(err.Error())
		return "", err
	}

	defer resp.Body.Close()

	if destination == "" {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			d.log.Error(err.Error())
			return "", err
		}
		return string(body), nil
	}

	output, err := os.Create(destination)
	if err != nil {
		d.log.Error(err.Error())
		return "", err
	}

	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		d.log.Error(err.Error())
		return "", err
	}

	return "", nil
}
