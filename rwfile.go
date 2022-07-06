package main

import (
	"os"
	"gopkg.in/yaml.v3"
	// "fmt"
)

/*
 ***
 */
func Write(file string, dat interface{}, /**/) {
  os.WriteFile(
	  "r.yml", Must(yaml.Marshal(dat)).([]byte), 0600,
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