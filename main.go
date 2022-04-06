package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type conf struct {
	ApiKey string `yaml:"api_key"`
}

func (c *conf) getConfig(apiKey string) *conf {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadFile(dirname + "/.greycom/config.yaml")
	if err != nil {
		//log.Printf("yamlFile.Get err   #%v ", err)
		if len(apiKey) != 0 {

			// check if .greycom dir exists
			path := dirname + "/.greycom"
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(path, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}

			err := os.WriteFile(dirname+"/.greycom/config.yaml", []byte("api_key: "+apiKey), 0755)
			if err != nil {
				fmt.Printf("Unable to write file: %v", err)
				os.Exit(0)
			}
		} else {
			fmt.Println("No config file found and no API Specified. Please re-run with -apikey flag.")
			os.Exit(0)
		}
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {

	var target string
	flag.StringVar(&target, "t", "", "Target IP Address to query")
	var apiKey string
	flag.StringVar(&apiKey, "apikey", "", "Your Greynoise Community API Key")
	flag.Parse()

	if len(target) == 0 {
		fmt.Println("Target IP (-t) is required.")
		os.Exit(0)
	}

	var c conf
	c.getConfig(apiKey)

	api_key := c.ApiKey

	fmt.Println("`~> -GreyCom v1.0-")

	req, err := http.NewRequest("GET", "https://api.greynoise.io/v3/community/"+target, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("key", api_key)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)

}
