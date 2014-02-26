package config

import (
  "encoding/json"
  "log"
  "os"
)

// Config contains all of the configuration settings required to bring up a
// local irc server.
type Config struct {
  Name string
  Port int
}

// FromJSONFile reads a Config struct from a file containing a JSON encoded
// value.
func FromJSONFile(path string) Config {
  var cfg Config
  file, err := os.Open(path)
  if err != nil {
    log.Fatalf("Could not open config file: %v.", err)
  }

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&cfg)
  if err != nil {
    log.Fatalf("Problem parsing config data: %v", err)
  }

  return cfg
}
