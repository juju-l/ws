package main

import (
  "fmt"
  "github.com/gorilla/websocket"
  "net/http"
  "strings"
  "time"
)

type wsEngine struct {
  shList map[string]*shCmd
  wsConList map[string]*websocket.Conn
  rList map[string]*[]string
}

func (ws *wsEngine) sendMsg(id string) {
  var i,s int
  for {
    if s == len(ws.rList) {
      /*;*/ delete(ws.wsConList, id); break
    }
    for k, v := range ws.rList {
        if ws.shList[k].isComplete && len(*v) == i {
            s ++; continue; /*;*/
        }
        if len(*v)-1-i < 0 { for { if len(*v)-1-i >= 0{ break } } }
        time.Sleep(time.Millisecond*100)
        err := ws.wsConList[id].WriteMessage(1, []byte("{\""+k+"\":\"--->"+(*v)[i]+"\"}")) //ws message send
        if err != nil {
            fmt.Println(err); delete(ws.wsConList, id); return
        }
    }
    /**/
    i ++
    /**/
  }
  /**/
}

func (ws *wsEngine) run(cfg *appConfig) *wsEngine {
  //
  if len(ws.rList) == 0 {
    if cfg.Ready != nil {
      /*err := */ newSh(strings.Join(*cfg.Ready, ";")).cmd.Wait() //
    }
    for k, v := range cfg.Sh {
      //
      ws.shList[k] = newSh(strings.Join(v, ";"))

      // go func() {
      //   i := 0
      //   for {
      //     if i > 100007 {
      //       break
      //     }
      //     if ws.shList[k].isComplete {
      //       // var t []string
      //       // t = (*ws.shList[k]).rst
      //       // ws.rList[k] = &t
      //       // delete(ws.shList, k)
      //       break
      //     }
      //     i ++
      //   }
      //   //
      // }()

      ws.rList[k] = &ws.shList[k].rst
      //
    }
    if cfg.Call != nil {
      /*err := */ newSh(strings.Join(*cfg.Call, ";")).cmd.Wait() //
    }
  }
  return ws
}

func (ws *wsEngine) cliRegister(id string, w http.ResponseWriter, r *http.Request) *wsEngine {
  cli := ws.wsConList[id]
  if cli == nil {
      u := websocket.Upgrader{}
      cli, _ = u.Upgrade(w, r, nil)
      ws.wsConList[id] = cli
  }
  return ws
}

func newWs() *wsEngine {
  var ws *wsEngine
  ws = &wsEngine{
    shList: make(map[string]*shCmd),
    wsConList: make(map[string]*websocket.Conn),
    rList: make(map[string]*[]string),
  }
  return ws
}

func init() {
  //
}
