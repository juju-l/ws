package main

import (
	"os"
	"os/exec"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"bufio"
	"fmt"
)

	var wss = websocket.Upgrader{}

func run(shCmd string) *bufio.Reader {
    cmd := exec.Command("sh", /*"-x",*/ "-c", shCmd+";echo -n \\#", "2>&1"); stdout,_ := cmd.StdoutPipe(); bufReader := bufio.NewReader(stdout); cmd.Stderr = os.Stderr; go cmd.Run(); return bufReader //
}

var (
	mux sync.Mutex
	buflist map[string]*bufio.Reader
	msgChannel = make(map[string]chan interface{})
	clilist = make(map[string]*websocket.Conn)
	sh *shCmd
)

func main() {
	r := gin.Default()
  r.LoadHTMLGlob("./*.htm")
	r.GET("/", func(c *gin.Context) {
      if c.Query("id") != "" {
        con, _ := wss.Upgrade(c.Writer, c.Request, nil); defer con.Close(); bufReader := run("packer --help"); for { log,_,_ := bufReader.ReadLine(); if string(log) == "#" { break }; con.WriteMessage(1, []byte("{\"consoleShow\":\""+string(log)+"\"}")) }; fmt.Println("shCmd exec finished...")
      } else {
        c.HTML(200, "default.htm", nil) //
      }
    })
  ///
	r.Run(":8080")
}

type shCmd struct {
  Ready *[]string
  Sh *map[string][]string
  Call *[]string
}

func init() {
	//
}