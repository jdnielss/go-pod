package main

import (
	"bytes"
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
	Body                 string `yaml:"body,omitempty"`
	ResponseBodyContains string `yaml:"response_body_contains,omitempty"`
	HTTPCodeIs           int    `yaml:"http_code_is,omitempty"`
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

		// Create HTTP request based on the method
		var req *http.Request
		var err error
		if contract.Method == http.MethodGet || contract.Method == http.MethodDelete {
			req, err = http.NewRequest(contract.Method, fullURL, nil)
		} else {
			req, err = http.NewRequest(contract.Method, fullURL, bytes.NewBuffer([]byte(contract.Body)))
			req.Header.Set("Content-Type", "application/json")
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

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s API %s test %s! Error reading response body: %v\n", crossIcon, contract.Name, crossIcon, err)
			continue
		}

		// Check HTTP status code
		if contract.HTTPCodeIs != 0 && resp.StatusCode != contract.HTTPCodeIs {
			fmt.Printf("%s API %s test %s! Expected HTTP code: %d, Actual: %d\n", crossIcon, contract.Name, crossIcon, contract.HTTPCodeIs, resp.StatusCode)
			continue
		}

		// Check if response body contains expected string
		if contract.ResponseBodyContains != "" && !strings.Contains(string(body), contract.ResponseBodyContains) {
			fmt.Printf("%s API %s test %s! Expected body to contain: %s, Actual: %s\n", crossIcon, contract.Name, crossIcon, contract.ResponseBodyContains, string(body))
			continue
		}

		fmt.Printf("%s API %s test %s!\n", checkIcon, contract.Name, checkIcon)
	}
}
