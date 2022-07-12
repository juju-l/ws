package main

import (
	//"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
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

func (ws *wsEngine) sendMsg(id string) {
  i := 0;
  //t := ws.rList;会同步删除
  t := make(map[string][]string)
  for k,v := range ws.rList {
	  t[k] = v
  }
  for {
    if len(t) == 0 {
      break
    }
    for k, v := range t {
      if len(ws.shList) == 0 { 
				if len(v) == i {
					delete(t, k);
					//break
				}
			}
      if len(v)-1 - i < 0 { 
				continue 
			}
      time.Sleep(time.Millisecond * 1000)
      /*for _,con := range ws.wsConList { */ws.wsConList[id].WriteMessage(1, []byte("{\""+k+"\":\"--->"+v[i]+"\"}")) //}
      //i++
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
    for m, n := range ws.shList {
      go func(k string, v shCmd) {
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
            i++
          }
        }
      } (m, *n)
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