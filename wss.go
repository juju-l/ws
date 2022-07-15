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
  /*#i := 0;*/idx:=make(map[string]int);s := []string{}
  ver := time.Now().Format("v010206r")
  /*;*/rls := *Yml[map[string]map[string][]string]("r.yml");t:=rls
    if rls[ver] == nil {
      /*;*/t = make(map[string]map[string][]string);t[ver] = make(map[string][]string)
    }
  for {
    if len(s) == len(ws.rList) {
      delete(ws.wsConList, id); if len(rls[ver]) == 0 { Write("r.yml", t) }; break
    }
    for k, v := range ws.rList {
        if (len(t[ver][k]) != 0||ws.shList[k].isComplete) && len(*v) == idx[k] {
          /*;*/is := false;for i := 0; i < len(s); i ++ { if s[i] == k { is = true } };if ! is { /*;*/s = append(s, k);t[ver][k] = *v };continue
        }
        if len(*v)-1-idx[k] < 0 { /*;*//*#for { if len(*v)-1-i >= 0{ break } }*/;continue }
        /*//---*/time.Sleep(time.Millisecond * 10)
        err := ws.wsConList[id].WriteMessage(1, []byte("{\""+k+"\":\""+(*v)[idx[k]]+"\"}")) //websocket message send
        if err != nil {
            fmt.Println(err); delete(ws.wsConList, id); return
        }; idx[k] = idx[k] + 1
    }
    /**/
    //#i ++
    /**/
  }
  /**/
  /***/
  /**/
  /**/
  /***/
  /**/
}

func (ws *wsEngine) run(cfg *appConfig) *wsEngine {
  ver := time.Now().Format("v010206r")
      rls := *Yml[map[string]map[string][]string]("r.yml")
  //     if rls == nil {
  //       rls = make(map[string]map[string][]string) 
  //     }
      if rls[ver] == nil {
  //       rls[ver] = make(map[string][]string)
  if len(ws.rList) == 0 {
    if cfg.Ready != nil {
      /*err := */ newSh(strings.Join(*cfg.Ready, ";")).cmd.Wait() //
    }
    for k, v := range cfg.Sh {
      //
      ws.shList[k] = newSh(strings.Join(v, ";"))

      // go func(k string) {
      //   //#i := 0
      //   for {
      //     //#if i > 100007 {
      //     //#  break
      //     //#}
      //     if ws.shList[k].isComplete {
      //       // var t []string
      //       //   t = (*ws.shList[k]).rst
      //       //   ws.rList[k] = &t
      //       //   delete(ws.shList, k)
      //       //
      //       rls[ver][k] = ws.shList[k].rst
      //       Write("r.yml", rls)
      //       //
      //       break
      //     }
      //     //#i ++
      //   }
      //   //
      // } (k)

      ws.rList[k] = &ws.shList[k].rst
      //
    }
    if cfg.Call != nil {
      /*err := */ newSh(strings.Join(*cfg.Call, ";")).cmd.Wait() //
    }
  }
      //
      } else {
  // if len(ws.rList) == 0 {
    for k, v := range rls[ver] {
      ws.rList[k] = &v
    }
  // }
      }
      ///
      ///
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
