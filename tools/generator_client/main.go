package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type GeneratorConfig struct {
	TargetDirectory   string `yaml:"target_directory"`
	MaxCountFiles     int    `yaml:"max_count_files"`
	PromptTemplate    string `yaml:"prompt_template"`
	PlaceholderSchema string `yaml:"placeholder_schema"`
	FilepathSchema    string `yaml:"filepath_schema"`
	ModelToUse        string `yaml:"model_to_use"`
	LlmUrl            string `yaml:"llm_url"`
}

type RequestBody struct {
	Model  string
	Prompt string
	Stream bool
}

type LlmResponse struct {
	Model      string `json:"model"`
	CreatedAt  string `json:"created_at"`
	Response   string `json:"response"`
	Done       bool   `json:"done"`
	DoneReason string `json:"done_reason"`
}

func BuildRequestBody(config GeneratorConfig) RequestBody {
	schema := ReadSchemaFromFile(config.FilepathSchema)
	prompt := strings.Replace(config.PromptTemplate, config.PlaceholderSchema, schema, -1)
	return RequestBody{
		Model:  config.ModelToUse,
		Prompt: prompt,
		Stream: false,
	}
}

func ReadConfig(filepath string) GeneratorConfig {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read config file %s: %v\n", filepath, err)
	}
	var config GeneratorConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file %s: %v\n", filepath, err)
	}
	return config
}

func ReadSchemaFromFile(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read schema file %s: %v\n", filepath, err)
	}
	return string(data)
}

func WriteResponseToFile(resp LlmResponse, fileName string) error {
	err := os.WriteFile(fileName, []byte(resp.Response), 0644)
	if err != nil {
		return fmt.Errorf("failed to write response body: %v", err)
	}
	log.Printf("Wrote response to file %s\n", fileName)
	return nil
}

func QueryGenerator(client *http.Client, config GeneratorConfig, jsonBytes []byte) (LlmResponse, error) {
	resp, err := client.Post(config.LlmUrl, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalf("Failed to post: %v", err)
	}
	if resp.StatusCode != 200 {
		if resp.ContentLength > 0 {
			bodyBytes, _ := io.ReadAll(resp.Body)
			log.Println(string(bodyBytes))
		}
		log.Fatalf("error when consuming post request: %v", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	var responseMap LlmResponse
	err = json.Unmarshal(bodyBytes, &responseMap)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	return responseMap, err
}

func main() {
	const configPath = "assets/config.yml"
	config := ReadConfig(configPath)
	jsonBody, err := json.Marshal(BuildRequestBody(config))
	if err != nil {
		log.Fatalf("Failed to marshal json: %v", err)
	}

	client := &http.Client{}

	log.Printf("Starting to query an llm %d times to generate test data", config.MaxCountFiles)

	for i := 0; i < config.MaxCountFiles; i++ {
		resp, err := QueryGenerator(client, config, jsonBody)
		if err != nil {
			log.Fatalf("Failed to query generator: %v", err)
		}
		filepath := fmt.Sprintf("%s/response_%d.xml", config.TargetDirectory, i)
		err = WriteResponseToFile(resp, filepath)
		if err != nil {
			log.Fatalf("Failed to write response: %v", err)
		}
	}

}
