package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gcfg.v1"
)

func config() (string, string, string, bool, string, int64) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	cfgpath := dir + "\\eveapi.cfg"

	file, err := ioutil.ReadFile(cfgpath) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	cfgStr := string(file)

	cfg := struct {
		Eve struct {
			Timeout  string
			KeyID    string
			VCode    string
			Onenotif bool
			Bot      string
			Userid   int64
		}
	}{}
	err = gcfg.ReadStringInto(&cfg, cfgStr)
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	return cfg.Eve.Timeout, cfg.Eve.KeyID, cfg.Eve.VCode, cfg.Eve.Onenotif, cfg.Eve.Bot, cfg.Eve.Userid

}
