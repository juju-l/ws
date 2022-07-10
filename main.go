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
						defer con.Close()

						rls := Yml[map[string]map[string][]string]("r.yml")

						ver := time.Now().Format("v010206r")

						if r,e := rls[ver]; !e {
							for k,v := range sh.Sh {
								rwLock.Lock(); bufReaderList[k] = Run(strings.Join(v, ";")); rwLock.Unlock()
							}
							rls[ver] = make(map[string][]string)
							/////rwLock.Lock(); Write("r.yml", rls); rwLock.Unlock()
							for{
								if len(bufReaderList) == 0 {
									break
								}
							for k, v := range bufReaderList {
								//go func() {
									//for {
										log, _, err := v.ReadLine()
										// reader := transform.NewReader(bytes.NewReader(log), simplifiedchinese.GBK.NewDecoder())
										if err != nil /*|| io.EOF == err*/ {
											delete(bufReaderList, k); delete(wsConList, c.Query("id")); break
										}
										con.WriteMessage(1, []byte("{\""+k+"\":\""+c.Query("id")+"--->"+string(log)+"\"}"))
										kLog := rls[ver][k]
										kLog = append(kLog, string(log))
										rls[ver][k] = kLog
										/////rwLock.Lock(); Write("r.yml", rls); rwLock.Unlock()
										time.Sleep(
											time.Millisecond * 100,
										)
									//}
								//} ()
							}
							}
						} else {
							for k,v := range r {
								con.WriteMessage(1, []byte("{\""+k+"\":\""+strings.Join(v, "\\n")+"\"}"))
							}
						}
						
						//
					} else {
						c.HTML(200, "default.htm", nil)
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