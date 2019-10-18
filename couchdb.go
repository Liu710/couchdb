package couchdb

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "time"
)

type CouchDB struct {
  client *http.Client
  url    *url.URL
}

type CouchDBConfig struct {
  Host     string
  Database string
  Username string
  Password string
  Timeout  int
}

// Create new CouchDB from config
func NewCouchDB(cfg CouchDBConfig) (*CouchDB, error) {
  if cfg.Timeout == 0 {
    cfg.Timeout = 10
  }
  client := &http.Client{
    Timeout: time.Duration(cfg.Timeout) * time.Second,
  }

  u, err := url.Parse(cfg.Host)
  if err != nil {
    return nil, ErrInitFailed
  }
  u.User = url.UserPassword(
    cfg.Username,
    cfg.Password,
  )
  u.Path = "/" + cfg.Database

  couchdb := &CouchDB{
    client: client,
    url:    u,
  }
  return couchdb, nil
}

// Response from database
type DBResp struct {
  Rows []struct {
    Id    string      `json:"id"`
    Key   string      `json:"key"`
    Value interface{} `json:"value"`
    Doc   interface{} `json:"doc"`
  } `json:"rows"`
}

// Send http request to database and read the response
func (c *CouchDB) requestDB(path string) ([]byte, error) {
  if path == "" {
    return nil, errRequestFailed
  }

  resp, err := c.client.Get(c.url.String() + path)
  if err != nil {
    return nil, errRequestFailed
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    if resp.StatusCode == 404 {
      return nil, errNotFound
    }
    return nil, errRequestFailed
  }

  result, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, errRequestFailed
  }

  return result, nil
}

// Get all docs from database
func (c *CouchDB) GetAllDocs() ([]interface{}, error) {
  path := "/_all_docs?include_docs=true"
  dbRespBytes, err := c.requestDB(path)
  if err != nil {
    return nil, ErrCannotGetDocs
  }

  var dbResp DBResp
  if err = json.Unmarshal(dbRespBytes, &dbResp); err != nil {
    return nil, ErrCannotGetDocs
  }
  result := make([]interface{}, len(dbResp.Rows))
  for i, row := range dbResp.Rows {
    result[i] = row.Doc
  }
  return result, nil
}

// Get doc from database
func (c *CouchDB) GetDoc(docID string) (interface{}, error) {
  path := "/" + docID
  dbRespBytes, err := c.requestDB(path)
  if err != nil {
    if err == errNotFound {
      return nil, ErrDocNotFound
    }
    return nil, ErrCannotGetDocs
  }

  var result interface{}
  if err = json.Unmarshal(dbRespBytes, &result); err != nil {
    return nil, ErrCannotGetDocs
  }
  return result, nil
}
