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
  /*#i := 0;*/idx:=make(map[string]int);s:=[]string{}
  ver := time.Now().Format("v010206r")
  /*;*/rls := *Yml[map[string]map[string][]string]("r.yml");t:=rls
    if rls[ver] == nil {
      /*;*/if rls==nil{t = make(map[string]map[string][]string)};t[ver]=make(map[string][]string)
    }
  for {
    if len(s) == len(ws.rList) {
      delete(ws.wsConList, id); /*if len(rls[ver])==0 {*/ Write("r.yml", t) /*}*/; break
    }
    for k, v := range ws.rList {
        if (t[ver][k] != nil||ws.shList[k].isComplete) && len(*v) == idx[k] {
          /*;*/is := false;for i := 0; i < len(s); i ++ { if s[i] == k { is = true } };if ! is { /*;*/if k == "ready" {for key,sh := range ws.shList{ if key != "ready"&&key != "call" { sh.cmd.Start() } }};s = append(s, k);if len(s) == len(ws.shList)-1 { ws.shList["call"].cmd.Start() };t[ver][k] = *v };continue
        }
        if len(*v)-1-idx[k] < 0 { /*;*/ /*#for { if len(*v)-1-i >= 0{ break } }*/; continue }
        /*//---*/time.Sleep(time.Millisecond * 10)
        err := ws.wsConList[id].WriteMessage(1, []byte("{\""+k+"\":\""+(*v)[idx[k]]+"\"}")) //websocket message send
        if err != nil { fmt.Println(err)
          delete(ws.wsConList, id);for _,v := range ws.shList { if !v.isComplete{ return } };/*if len(rls[ver]) == 0 {*/ Write("r.yml", t) /*}*/
        return }; idx[k] = idx[k] + 1
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
  //         rls = make(map[string]map[string][]string)
  //     }
      if rls[ver] == nil {
  //       rls[ver] = make(map[string][]string)
  if len(ws.rList) == 0 {
    // if cfg.Ready != nil {
      ready := strings.Join(*cfg.Ready, ";")
      sh := newSh(ready);ws.shList["ready"] = sh;sh.cmd.Start()
      ws.rList["ready"] = &ws.shList["ready"].rst
    // }
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
    // if cfg.Call != nil {
      /////
      /*;*/ws.shList["call"] = newSh(strings.Join(*cfg.Call, ";"))/*;*/
      ws.rList["call"] = &ws.shList["call"].rst
    // }
  }
      //
      } else {
  // if len(ws.rList) == 0 {
    for k, v := range rls[ver] {
      t := v
      // fmt.Println( &t )
      ws.rList[k] = &t
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