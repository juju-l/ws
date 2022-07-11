package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func newSh(s string) *shCmd {
	var sh *shCmd
	sh = &shCmd{ cmd: exec.Command("sh", "+xe", "-c", s, "2>&1") }
	sh.cmd.Stderr = os.Stderr; stdout, _ := sh.cmd.StdoutPipe(); //sh.cmd.Stdin = os.Stdin
	sh.cmd.Start()
	time.Sleep(time.Millisecond * 100)
	buf := bufio.NewReader(stdout)
	fmt.Println(buf)
	go func() { for { log, err := buf.ReadBytes('\n'); if err != nil /*|| io.EOF == err*/ { /**/; sh.isComplete = true; break }; sh.rst = append(sh.rst, string(log[:len(log)-1])) } } ()
	return sh
}

type shCmd struct {
	//
	cmd *exec.Cmd
	isComplete bool
	rst []string
	//
}