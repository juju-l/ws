package main

import (
  //"fmt"
  "github.com/gorilla/websocket"
  "net/http"
  "strings"
  "time"
)

func newWs() *wsEngine {
  var ws *wsEngine
      ws = &wsEngine{
      shList: make(map[string]*shCmd),
      wsConList: make(map[string]*websocket.Conn),
      rList: make(map[string][]string),
      }
  return ws
}

func (ws *wsEngine) sendBroadcastMsg() {
  i := 0;t := ws.rList;for {
    if len(t) == 0 {
      break
    }
    for k, v := range t {
      //
      if len(v)-1 - i < 0 { continue }
      time.Sleep(time.Millisecond * 1000)
      for _,con := range ws.wsConList { con.WriteMessage(1, []byte("{\""+k+"\":\"--->"+v[i]+"\"}")) }
      i++
    }
    /*if len(ws.shList) != 0 {
      continue
    }*/
    i ++
  }
}

func (ws *wsEngine) run (cfg *appConfig) *wsEngine {
  //
  if len(ws.rList) == 0 {
    if cfg.Ready != nil {
      /*err := */newSh(strings.Join(*cfg.Ready, ";")).cmd.Wait()
    }
    for k, v := range cfg.Sh {
      //
      ws.shList[k] = newSh(strings.Join(v, ";"))
      //
    }
    for k, v := range ws.shList {
      go func() {
        i := 0;for {
          if len(v.rst) > 0 {
            if len(v.rst) == i {
              if v.isComplete {
                //
                delete(ws.shList, k)
                break
              } else {
                continue
              }
            }
            
            ws.rList[k] = append(ws.rList[k], v.rst[i])
            time.Sleep(time.Millisecond * 1000)
            i++
          }
        }
      } ()
    }
    /*for {
      if len(ws.shList) == 0 {
        break
      }
    }*/
    if cfg.Call != nil {
      /*err := */newSh(strings.Join(*cfg.Call, ";")).cmd.Wait()
    }
  }
  return ws
}

func (ws *wsEngine) cliRegister (id string, w http.ResponseWriter, r *http.Request) *wsEngine {
  cli := ws.wsConList[id]
  if cli == nil {
      u := websocket.Upgrader{}
      cli, _ = u.Upgrade(w, r, nil)
      ws.wsConList[id] = cli
  }
  return ws
}

type wsEngine struct {
  shList map[string]*shCmd
  wsConList map[string]*websocket.Conn
  rList map[string][]string
}

func init() {
  //
}