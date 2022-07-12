package main

import (
  "fmt"
  "bufio";/*"os";*/"os/exec"
  "time"
)

func newSh(s string) *shCmd {
  var sh *shCmd
      sh = &shCmd{ cmd: exec.Command("sh", "-xe", "-c", s, "2>&1") }
      /*sh.cmd.Stdin = os.Stdin;*/ stdout, _ := sh.cmd.StdoutPipe(); sh.cmd.Stderr = sh.cmd.Stdout
      sh.cmd.Start()
      time.Sleep(time.Millisecond * 100)
      b := bufio.NewReader(stdout)
      go func() {
      for { log, err := b.ReadBytes('\n'); if err != nil /*|| io.EOF == err*/ { /**/; sh.isComplete = true; break }; sh.rst = append(sh.rst, string(log[:len(log)-1])) }
      } ()
      fmt.Println(b)
  return sh
}

type shCmd struct {
  cmd *exec.Cmd
  isComplete bool
  rst []string
}