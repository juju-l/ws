package main

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
  // "fmt"
)

/*
 ***
 */
func Yml [T any] (ymlfile string) *T {
  var t T;ti := &t;//
  yaml.Unmarshal(
      Must(ioutil.ReadFile(ymlfile)).([]byte), ti,
    )
  return ti
}

//
///*/
/*****/
///*/
//
func init() {
  //
}