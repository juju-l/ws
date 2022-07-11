package main

import (
  "bufio"
  "os"
  "os/exec"
)

func newSh(s string) *shCmd {
  var sh *shCmd
      sh = &shCmd{ cmd: exec.Command("sh", "+xe", "-c", s, "2>&1") }
      sh.cmd.Stderr = os.Stderr; stdout, _ := sh.cmd.StdoutPipe(); //sh.cmd.Stdin = os.Stdin
      sh.cmd.Start()
      b := bufio.NewReader(stdout)
      go func() { for { log, err := b.ReadBytes('\n'); if err != nil /*|| io.EOF == err*/ { /**/; sh.isComplete = true; break }; sh.rst = append(sh.rst, string(log[:len(log)-1])) } } ()
  return sh
}

type shCmd struct {
  cmd *exec.Cmd
  isComplete bool
  rst []string
}