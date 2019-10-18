package couchdb

import "errors"

var (
  errRequestFailed = errors.New("Error: Request failed")
  errNotFound      = errors.New("Error: Not Found")

  ErrInitFailed    = errors.New("Error: Cannot initialize couchdb")
  ErrCannotGetDocs = errors.New("Error: Cannot get docs from database")
  ErrCannotGetDoc  = errors.New("Error: Cannot get docs from database")
  ErrDocNotFound   = errors.New("Error: Cannot find document in the dabtase")
)
