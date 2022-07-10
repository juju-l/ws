package main

import (
	"bufio"
	// "bytes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	// "golang.org/x/text/encoding/simplifiedchinese"
	// "golang.org/x/text/transform"
	"sync"
	"strings"
	"time"
)

var (
	rwLock sync.RWMutex
	wssObject = websocket.Upgrader{}
	bufReaderList = make(map[string]*bufio.Reader)
	wsConList = make(map[string]*websocket.Conn)
	sh shCmd
	rls = Yml[map[string]map[string][]string]("r.yml")
	isExec = false
)

type tst struct {
	//
}

func main() {
	r := gin.Default()
			r.LoadHTMLGlob("./*.htm")
			r.GET("/", func(c *gin.Context) {
					if c.Query("id") != "" {
						con := wsConList[c.Query("id")]
						if con == nil {
						/*var err error;*/ con, _ = wssObject.Upgrade(c.Writer, c.Request, nil); wsConList[c.Query("id")] = con
						}
						defer delete(
							wsConList, c.Query("id"),
						)
						
						ver := time.Now().Format("v010206r")

						if isExec == false {
							rls[ver] = make(map[string][]string)
							for k,v := range sh.Sh {
								bufReaderList[k] = Run(strings.Join(v, ";"))
							}
							for k, v := range bufReaderList {
								go func() {
									for {
										log, _, err := v.ReadLine()
										if err != nil /*|| io.EOF == err*/ {
											delete(bufReaderList, k)
											// delete(wsConList, c.Query("id"))
											break
										}
										kLog := rls[ver][k]
										kLog = append(kLog, string(log))
										rls[ver][k] = kLog
									}
								} ()
							}
							// for {
								//
							// }
							isExec = !isExec
						}
						

						//if r,e := rls[ver]; !e {
							
							
							
								idx := 0
								s := len(rls[ver])
								
							for {
							
							
								for k, v := range rls[ver] {
									/*if len(v) < idx && bufReaderList[k] == nil {
										s = s - 1;continue
									}*/
									if string(v[idx]) == "#" {
										s = s - 1;continue
									}
 									con.WriteMessage(1, []byte("{\""+k+"\":\""+c.Query("id")+"--->"+string(v[idx])+"\"}"))
									time.Sleep(
										time.Millisecond * 100,
									)
									idx += 1
								}
								
								
								if s == 0 {
									break
								}
								
							}
							
						rwLock.Lock(); Write("r.yml", rls); rwLock.Unlock() // 写入结果到 yml 文件	
							
							
						/*} else {
							for k,v := range r {
								con.WriteMessage(1, []byte("{\""+k+"\":\""+strings.Join(v, "\\n")+"\"}"))
							}
						}*/
						
						//
					} else {
						c.HTML(200, "default.htm", gin.H { "tstmp":time.Now().Unix() })
					}
			})
	r.Run(":8080")
}

type shCmd struct {
	Ready *[]string
	Sh map[string][]string
	Call *[]string
}

func init() {
	sh = Yml[shCmd]("c.yml")
}