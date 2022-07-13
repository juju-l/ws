package main

import (
  "bufio"
  "os"
  "os/exec"
)

/*
 ***
 */
func Run(shCmd string) *bufio.Reader {
  var b *bufio.Reader
  c := exec.Command("sh", "+xe", "-c", shCmd, "2>&1")
  c.Stderr = os.Stderr; stdout, _ := c.StdoutPipe(); //c.Stdin = os.Stdin;
  e := c.Start()
  if e != nil {
    panic(e)
  }
  b = bufio.NewReader(stdout)//一次性只读
  return b
}

//
///*/
/*****/
///*/
//
func init() {
  //
}