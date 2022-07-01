package main

import (
	"os"
	"os/exec"
    // "time"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"bufio"
	"fmt"
)

var wss = websocket.Upgrader{}

func main() {
	r := gin.Default()
  r.LoadHTMLGlob("./*.htm")
	r.GET("/", func(c *gin.Context) {
      if c.Query("id") != "" {
        // wss := websocket.Upgrader{}
        con, _ := wss.Upgrade(c.Writer, c.Request, nil)
        defer con.Close()
        cmd := exec.Command("sh", /*"-x",*/ "-c", "packer build -var version_kic=b725891a packer-kic.json;echo \\#", "2>&1")
        stdout,_ := cmd.StdoutPipe()
        soBuffer := bufio.NewReader(stdout)
        cmd.Stderr = os.Stderr
        cmd.Start()
        for {
          log,_,_ := soBuffer.ReadLine()
          if string(log) == "#" {
            return
          }
          con.WriteMessage(1, log)
        }
        cmd.Wait()
        fmt.Println(".")
      } else {
        c.HTML(200, "default.htm", nil)
      }
    })
  ///
	r.Run(":8080")
}

func init() {
	//
}