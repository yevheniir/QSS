package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// YamlConfig ...
type YamlConfig struct {
	Influx struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	}
	Queries string `yaml:"queries"`
}

func parceYAML() YamlConfig {
	fmt.Println("Parsing YAML file")

	var fileName string
	flag.StringVar(&fileName, "f", "", "YAML file to parse.")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide yaml file by using -f option")
		os.Exit(1)
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		os.Exit(1)
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Result: %v\n", yamlConfig)
	return yamlConfig
}
