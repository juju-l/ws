package main

import (
	"bufio"
	"net/http"
	"github.com/gorilla/websocket"
	"os"
	"os/exec"
	//   "time"
	"fmt"
)

var upGrader = websocket.Upgrader{}

func wssHandle(rsp http.ResponseWriter, req *http.Request) {
  wss, _ := upGrader.Upgrade(rsp, req, nil)
  cmd := exec.Command("packer", "build", "-var version_kic=b725891a", "packer-kic.json", "2>&1")
  stdout,_ := cmd.StdoutPipe(); soBuffer := bufio.NewReader(stdout); cmd.Stderr = os.Stderr
  cmd.Start(); for /*range time.Tick(0.1 * time.Second)*/ { line,_,_ := soBuffer.ReadLine(); wss.WriteMessage(websocket.TextMessage, line); fmt.Println(string( line )) }; //cmd.Wait()
  fmt.Println("test")
}

func main() {
  http.HandleFunc("/fw", wssHandle)
  http.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
    http.ServeFile(rsp, req, "default.htm")
  })
  http.ListenAndServe(":8080", nil)
}

func init() {
	//
}