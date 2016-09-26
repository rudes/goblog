package main

import (
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	_testRoot := os.Getenv("GOPATH") + "/src/github.com/rudes/otherletters.net/static/"
	expected := Config{
		Title:       "Example Title",
		Subtitle:    "Example Sub-Title",
		Author:      "Example Writer",
		AuthorEmail: "writer@example.com",
	}
	actual, err := getConfig(_testRoot + "example.toml")
	if err != nil {
		t.Errorf("Expected nil, got : %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, actual)
	}
}

func TestGetConfigBadFile(t *testing.T) {
	_testRoot := os.Getenv("GOPATH") + "/src/github.com/rudes/otherletters.net/static/"
	_, err := getConfig(_testRoot + "invalid.toml")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
