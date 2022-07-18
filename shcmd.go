package main

import (
  "fmt"
  "bufio"; "io"; /*strings;*/ /*os;*/ "os/exec"
  "time"
)

func newSh(s string) *shCmd {
  var sh *shCmd
      sh = &shCmd{ cmd: exec.Command("sh", "+e", "-c", s, "2>&1") }
      /*sh.cmd.Stdin = os.Stdin;*/ stdout, _ := sh.cmd.StdoutPipe(); sh.cmd.Stderr = sh.cmd.Stdout
      sh.cmd.Start()
      // time.Sleep(time.Millisecond * 100)
      b := bufio.NewReader(stdout)
      go func() {
      for { log, err := b.ReadBytes('\n'); if err != nil || io.EOF == err /**/ { /**/; sh.isComplete = true; break }; sh.rst = append(sh.rst, fmt.Sprintf("%d->%s<-%d", time.Now().Unix(),string(log[:len(log)-1]),time.Now().Unix()) /**/) }
      } ()
      // fmt.Println(b)
  return sh
}

type shCmd struct {
  cmd *exec.Cmd
  isComplete bool
  rst []string
}