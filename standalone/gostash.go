package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hobbs/mustache"
	"io/ioutil"
	"log"
	"os"
)

func usage() {
	flag.PrintDefaults()
	log.Fatal("Must provide json data file and template.")
}

func main() {

	data := flag.String("d", "", "Data file")
	template := flag.String("t", "", "Template file")
	flag.Parse()

	if len(*template) == 0 {
		usage()
	}

	var raw []byte
	var err error
	if len(*data) == 0 {
		raw, err = ioutil.ReadAll(os.Stdin)
		if err != nil || len(raw) == 0 {
			usage()
		}
	} else {
		raw, err = ioutil.ReadFile(*data)
		if err != nil || len(raw) == 0 {
			log.Fatal("Could not load file", err)
		}
	}

	var parsed interface{}
	err = json.Unmarshal(raw, &parsed)
	if err != nil {
		log.Fatal("Failed to parse JSON file ", err)
	}

	fmt.Println(mustache.RenderFile(*template, parsed))
}
