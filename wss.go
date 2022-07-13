package main

import (
  //"fmt"

  "fmt"
  "net/http"
  "strings"
  "time"

  //"time"

  "github.com/gorilla/websocket"
)

func newWs() *wsEngine {
  var ws *wsEngine
  ws = &wsEngine{
    shList:    make(map[string]*shCmd),
    wsConList: make(map[string]*websocket.Conn),
    rList:     make(map[string]*[]string),
  }
  return ws
}

func (ws *wsEngine) sendMsg(id string) {
  s := 0

  // for {
  //   if len(ws.rList) != 0 {
  //     break
  //   }
  // }

  i := 0
  for {
    if s == len(ws.rList) {
      //
      delete(ws.wsConList, id)
      break
    }


    //
    for k, v := range ws.rList {
        if ws.shList[k].isComplete && len(*v) == i {
          s++
          continue
          //
        }

        if len(*v)-1-i < 0 {
          for { // 防止因时间差原因跳过本次数据，在此等待
            time.Sleep(time.Millisecond*100)
            if len(*v)-1-i >= 0{
              break
            }
          }
        }

        time.Sleep(time.Millisecond*100)
        err := ws.wsConList[id].WriteMessage(1, []byte("{\""+k+"\":\"--->"+(*v)[i]+"\"}"))
        if err != nil {
          fmt.Println(err)
          delete(ws.wsConList, id)
          return
        }
    }
    i++ //需要确保上面本次 i 的所有数组都取过一次，如果没有取到，除非本身不存在外，需要在上边进行手工等待，防止因时间差原因而错行


  }
  //

  // for {
  //   if s == len(ws.rList) {
  //     //
  //     delete(ws.wsConList, id)
  //     break
  //   }
  // }
}

func (ws *wsEngine) run(cfg *appConfig) *wsEngine {
  //
  if len(ws.rList) == 0 {
    if cfg.Ready != nil {
      /*err := */ newSh(strings.Join(*cfg.Ready, ";")).cmd.Wait()
    }
    for k, v := range cfg.Sh {
      //
      ws.shList[k] = newSh(strings.Join(v, ";"))

      // go func() {
      //   i := 0
      //   for {
      //     time.Sleep(time.Millisecond * 10)
      //     fmt.Printf("%s---%v", k, i)
      //     if ws.shList[k].isComplete {
      //       // var t []string
      //       // t = (*ws.shList[k]).rst
      //       // ws.rList[k] = &t
      //       // delete(ws.shList, k)
      //       break
      //     }
      //     if i > 100007 {
      //       break
      //     }
      //     i++
      //   }
      // }()

      ws.rList[k] = &ws.shList[k].rst
      //
    }
    if cfg.Call != nil {
      /*err := */ newSh(strings.Join(*cfg.Call, ";")).cmd.Wait()
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

type wsEngine struct {
  shList    map[string]*shCmd
  wsConList map[string]*websocket.Conn
  rList     map[string]*[]string
}

func init() {
  //
}
