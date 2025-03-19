package main

import (
	"encoding/json"
	"io"
	"main/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestBuildRequestBody(t *testing.T) {
	expected := RequestBody{
		Model:  "model",
		Prompt: "some test data",
		Stream: false,
	}

	config := GeneratorConfig{
		FilepathSchema:    "./testdata/schema.txt",
		PromptTemplate:    "some <placeholder> data",
		PlaceholderSchema: "<placeholder>",
		ModelToUse:        "model",
	}

	actual := BuildRequestBody(config)

	assert.Equal(t, actual, expected)
}

func TestQueryGenerator(t *testing.T) {
	request := RequestBody{
		Model:  "model",
		Prompt: "some test data",
		Stream: false,
	}
	jsonBytes, _ := json.Marshal(request)
	expected := LlmResponse{
		Model:      "model",
		CreatedAt:  time.Now().Format(time.RFC3339),
		Response:   "some response",
		Done:       false,
		DoneReason: "done",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/generate" {
			t.Errorf("Expected path to be /api/generate, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected method to be POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected content-type to be application/json, got %s", r.Header.Get("Content-Type"))
		}
		if r.Body == nil {
			t.Errorf("Expected request body to not be nil")
		}
		if r.Body != nil {
			requestBytes, _ := io.ReadAll(r.Body)
			var requestBody RequestBody
			_ = json.Unmarshal(requestBytes, &requestBody)
			assert.Equal(t, requestBody, request)
		}
		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(expected)
		_, _ = w.Write(responseBytes)
	}))
	defer server.Close()
	config := GeneratorConfig{
		TargetDirectory:   "./generated",
		MaxCountFiles:     1,
		PromptTemplate:    "some <placeholder> data",
		PlaceholderSchema: "<placeholder>",
		FilepathSchema:    "./testdata/schema.txt",
		ModelToUse:        "model",
		LlmUrl:            server.URL + "/api/generate",
	}

	actual, _ := QueryGenerator(&http.Client{}, config, jsonBytes)

	assert.Equal(t, actual, expected)
}

func TestReadConfig(t *testing.T) {
	expected := GeneratorConfig{
		TargetDirectory:   "./generated",
		MaxCountFiles:     5,
		PromptTemplate:    "This is a test case ### <schema> ###",
		PlaceholderSchema: "<schema>",
		FilepathSchema:    "./testdata/schema.xsd",
		ModelToUse:        "model",
		LlmUrl:            "http://localhost:8080",
	}

	actual := ReadConfig("./testdata/config.yml")

	assert.Equal(t, actual, expected)
}

func TestWriteResponseToFile(t *testing.T) {
	tempFile, _ := os.CreateTemp("./generated", ".testfile")
	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name())

	response := LlmResponse{
		Model:      "model",
		CreatedAt:  time.Now().String(),
		Response:   "Some generated Response",
		Done:       false,
		DoneReason: "done",
	}

	err := WriteResponseToFile(response, tempFile.Name())

	if err != nil {
		t.Error(err)
	}
	rawContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(rawContent), response.Response)
}
