package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Contract struct {
	Name                 string `yaml:"name"`
	Path                 string `yaml:"path"`
	Method               string `yaml:"method"`
	ResponseBodyContains string `yaml:"response_body_contains"`
	RequestBody          string `yaml:"body"`
	HTTPCodeIs           int    `yaml:"http_code_is"`
}

type Contracts struct {
	Contracts []Contract `yaml:"contracts"`
}

func main() {
	// Parse command line arguments
	yamlFilePath := flag.String("f", "smoke_test.yaml", "Path to YAML config file")
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

		// Create the request
		var req *http.Request
		var err error

		switch strings.ToUpper(contract.Method) {
		case "GET":
			req, err = http.NewRequest(http.MethodGet, fullURL, nil)
		case "POST":
			req, err = http.NewRequest(http.MethodPost, fullURL, bytes.NewBufferString(contract.RequestBody))
		case "PUT":
			req, err = http.NewRequest(http.MethodPut, fullURL, bytes.NewBufferString(contract.RequestBody))
		case "DELETE":
			req, err = http.NewRequest(http.MethodDelete, fullURL, nil)
		default:
			fmt.Printf("%s API %s test %s! Unsupported method: %s\n", crossIcon, contract.Name, crossIcon, contract.Method)
			continue
		}

		if err != nil {
			fmt.Printf("%s API %s test %s! Error creating request: %v\n", crossIcon, contract.Name, crossIcon, err)
			continue
		}

		// Hit the API
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s API %s test %s! Error hitting API: %v\n", crossIcon, contract.Name, crossIcon, err)
			continue
		}
		defer resp.Body.Close()

		// Check HTTP status code if specified
		if contract.HTTPCodeIs != 0 && resp.StatusCode != contract.HTTPCodeIs {
			fmt.Printf("%s API %s test %s! Expected HTTP code: %d, Actual: %d\n", crossIcon, contract.Name, crossIcon, contract.HTTPCodeIs, resp.StatusCode)
			continue
		}

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s API %s test %s! Error reading response body: %v\n", crossIcon, contract.Name, crossIcon, err)
			continue
		}

		// Check if response body contains expected string
		if contract.ResponseBodyContains != "" && !strings.Contains(string(body), contract.ResponseBodyContains) {
			fmt.Printf("%s API %s test %s! Expected: %s, Actual: %s\n", crossIcon, contract.Name, crossIcon, contract.ResponseBodyContains, string(body))
			continue
		}

		fmt.Printf("%s API %s test %s!\n", checkIcon, contract.Name, checkIcon)
	}
}
