package main

import (
  "os"
  "gopkg.in/yaml.v2"
  // "fmt"
)

/*
 ***
 */
func Write(file string, dat interface{}, /**/) {
  os.WriteFile(
    "r.yml", append([]byte("\n"), Must(yaml.Marshal(dat)).([]byte)... /* , */), 0600,
    )
}

//
///*/
/*****/
///*/
//
func init() {
  //
}