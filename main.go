package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type Contract struct {
	Name                 string `yaml:"name"`
	Path                 string `yaml:"path"`
	Method               string `yaml:"method"`
	ResponseBodyContains string `yaml:"response_body_contains"`
}

type Contracts struct {
	Contracts []Contract `yaml:"contracts"`
}

func main() {
	// Parse command line arguments
	yamlFilePath := flag.String("f", "config.yaml", "Path to YAML config file")
	baseURL := flag.String("h", "http://localhost:8080", "Base URL of the API")
	flag.Parse()

	// Load YAML configuration
	yamlFile, err := ioutil.ReadFile(*yamlFilePath)
	if err != nil {
		panic(err)
	}

	var contracts Contracts
	err = yaml.Unmarshal(yamlFile, &contracts)
	if err != nil {
		panic(err)
	}

	// Define check and cross icons
	checkIcon := "✔"
	crossIcon := "✘"

	// Iterate over contracts and hit the APIs
	for _, contract := range contracts.Contracts {
		fmt.Printf("Testing contract: %s\n", contract.Name)

		// Construct full URL
		fullURL := strings.TrimSuffix(*baseURL, "/") + contract.Path

		// Hit the API
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf("%s API %s test %s! Error hitting API: %v\n", crossIcon, contract.Name, crossIcon, err)
			continue
		}
		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s API %s test %s! Error reading response body: %v\n", crossIcon, contract.Name, crossIcon, err)
			continue
		}

		// Check if response body contains expected string
		if strings.Contains(string(body), contract.ResponseBodyContains) {
			fmt.Printf("%s ls  %s ll %s!\n", checkIcon, contract.Name, checkIcon)
		} else {
			fmt.Printf("%s API %s test %s! Expected: %s, Actual: %s\n", crossIcon, contract.Name, crossIcon, contract.ResponseBodyContains, string(body))
		}
	}
}
