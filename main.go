package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func main() {
	var s Settings
	err := envconfig.Process("", &s)
	if err != nil {
		err = errors.Wrap(err, "processing config")
		log.Fatal(err)
	}

	for _ = range time.Tick(time.Second * 5) {
		alive, err := serverIsAlive(&s)
		if err != nil {
			err = errors.Wrap(err, "checking server liveness")
			log.Println("[ERROR]", err)
			continue
		}
		if !alive {
			err = resetServer(&s)
			if err != nil {
				err = errors.Wrap(err, "resetting server")
				log.Println("[ERROR]", err)
			}
			continue
		}
		log.Println("[INFO] server is alive")
	}
}

func serverIsAlive(s *Settings) (bool, error) {
	resp, err := http.Get(s.TrueNasPingUrl)
	if err != nil {
		err = errors.Wrap(err, "connecting to TrueNAS server")
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return false, nil
	}

	return true, nil
}
