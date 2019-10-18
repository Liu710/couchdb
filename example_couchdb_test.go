package couchdb_test

import (
  "testing"

  "github.com/Liu710/couchdb"
)

var docID string

func TestNewCouchDB(t *testing.T) {
  _, cErr := couchdb.NewCouchDB(couchdb.CouchDBConfig{
    Host:     "http://localhost:5984",
    Username: "test_user",
    Password: "test_password",
    Database: "test_db",
  })
  if cErr != nil {
    t.Fatal("Failed to initialize couchdb")
  }
}

func TestGetAllDocs(t *testing.T) {
  c, cErr := couchdb.NewCouchDB(couchdb.CouchDBConfig{
    Host:     "http://localhost:5984",
    Username: "test_user",
    Password: "test_password",
    Database: "test_db",
  })
  if cErr != nil {
    t.Fatal("Failed to initialize couchdb")
  }
  result, err := c.GetAllDocs()
  if err != nil {
    t.Fatal("Failed to get all docs")
  }
  if len(result) < 1 {
    t.Fatal("No doc is in the database")
  }
  doc := result[0].(map[string]interface{})
  docID = doc["_id"].(string)
}

func TestGetDoc(t *testing.T) {
  c, cErr := couchdb.NewCouchDB(couchdb.CouchDBConfig{
    Host:     "http://localhost:5984",
    Username: "test_user",
    Password: "test_password",
    Database: "test_db",
  })
  if cErr != nil {
    t.Fatal("Failed to initialize couchdb")
  }
  _, err1 := c.GetDoc("id_does_not_exist")
  if err1 != nil && err1 != couchdb.ErrDocNotFound {
    t.Fatal("Failed to get all docs")
  }
  _, err2 := c.GetDoc(docID)
  if err2 != nil {
    t.Fatal("Failed to get all docs")
  }
}
