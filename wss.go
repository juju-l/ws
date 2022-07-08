package main

import (
	"bufio"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

	var wss = websocket.Upgrader{}

/*
 ***
 */
func (ws *wsEngine) DestroyWs() {
	ws = nil
}

/*
 ***
 */
func NewEngineWs(sh map[string][]string) *wsEngine {
	var ws *wsEngine
			ws = &wsEngine{
					Mux: &sync.Mutex{},
					// Buf: make(map[string]*bufio.Reader),
					MsgChannel: make(map[string]chan []byte),
					Con: make(map[string]*websocket.Conn),
					Sh: sh,
			}
	return ws
}

/*
 ***
 */
func (ws *wsEngine) CliRegister(id string, w http.ResponseWriter, r *http.Request) *wsEngine {
	cli := ws.Con[id]
	if cli == nil {
			cli, _ = wss.Upgrade(w, r, nil)
			ws.MsgChannel[id] = make(chan []byte, 1)
			ws.Con[id] = cli
			
	}
	return ws
}

/*
 ***
 */
func (ws *wsEngine) RevBroadcastMsg() {
			for {
	for k,v := range ws.Con {
					if ws.Buf != nil && len(ws.Buf) == 0 {
							break
					}
			if ws.Buf != nil {
					for l,m := range ws.Buf {
							select {
							case log, ok := <- ws.MsgChannel[k] :
									var err error
									err = v.WriteMessage(1, []byte("{\""+l+"\":\""+string(log)+"\"}"))
									if err != nil {
											delete(ws.Con, k)
											if ok {
													go func() { ws.MsgChannel[k] <- log } ()
											}
											return
									}
									//
									time.Sleep( time.Millisecond*1500 )
							default :
									log,_,_ := m.ReadLine()
									if log != nil {
											if string(log) == "#" {
													/**/; delete(ws.Buf, l); break
											}
											/*go func () { */ws.MsgChannel[k] <- log/* } ()*/
											///
									} else {
											continue
									}
							}
					}
			} else {
					ws.Buf = make(map[string]*bufio.Reader)
					for i,j := range ws.Sh {
							ws.Mux.Lock(); ws.Buf[i] = Run(strings.Join(j, ";")); ws.Mux.Unlock()
					}
					///
			}
					// 写入yml
	}
					if ws.Buf != nil && len(ws.Buf) == 0 {
							break
					}
			}
}

/*
 ***
 */
func Run(str string) *bufio.Reader {
	var cmd *exec.Cmd
			cmd = exec.Command("sh", /*"-x",*/ "-c", str+";echo -n \\#", "2>&1")
			cmd.Stderr = os.Stderr; stdout,_ := cmd.StdoutPipe(); cmd.Start()
			buf := bufio.NewReader(stdout)
	return buf //
}

/*
 ***
 */
type wsEngine struct {
					Mux *sync.Mutex
					Buf map[string]*bufio.Reader
					MsgChannel map[string]chan []byte
					Con map[string]*websocket.Conn
					Sh map[string][]string
}

//
///*/
/*****/
///*/
//
func init() {
	//
}