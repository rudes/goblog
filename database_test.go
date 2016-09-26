package main

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestOpenDB(t *testing.T) {
	_, err := openDB()
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
func TestGetAllNilDB(t *testing.T) {
	cluster := gocql.NewCluster("db")
	cluster.Hosts = nil
	session, err := cluster.CreateSession()
	if err == nil {
		session.Close()
		t.Error("Expected an error, got nil")
	}
	_, err = getall(session)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
