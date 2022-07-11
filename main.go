package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	wssObject = websocket.Upgrader{}
	wsConList = make(map[string]*websocket.Conn)
	cfg       *appConfig
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./*.htm")
	rls := *Yml[map[string]map[string][]string]("r.yml")
	isExec := false
	shCmdList := make(map[string]*shCmd)
	r.GET("/", func(c *gin.Context) {
		if c.Query("id") != "" {
			con := wsConList[c.Query("id")]
			if con == nil {
				/*var err error;*/ con, _ = wssObject.Upgrade(c.Writer, c.Request, nil)
				wsConList[c.Query("id")] = con
			}
			defer delete(wsConList, c.Query("id"))

			ver := time.Now().Format("v010206r")

			if isExec == false { isExec = true
				if cfg.Ready != nil {
					newSh(strings.Join(*cfg.Ready, ";")).cmd.Wait()
				}

				//shCmdList := make(map[string]*shCmd)
				for k, v := range cfg.Sh {
					//
					shCmdList[k] = newSh(strings.Join(v, ";"))
					//
				}
				for i, j := range shCmdList {
					go func(k string, v *shCmd) {
						i := 0;for {
							if len(v.rst) > 0 {
								if len(v.rst) == i {
									if v.isComplete {
										//
										delete(shCmdList, k)
										break
									} else {
										continue
									}
								}
								if rls[ver] == nil {
									rls[ver] = make(map[string][]string)
								} else {
									rls[ver][k] = append(rls[ver][k], v.rst[i])
								}
								con.WriteMessage(1, []byte("{\""+k+"\":\""+c.Query("id")+"--->"+v.rst[i]+"\"}"))
								time.Sleep(time.Millisecond * 1000)
								i++
							}
						}
					}(i, j)
				}
				for {
					if len(shCmdList) == 0 {
						fmt.Print(rls[ver])
						////////Write("r.yml", rls)
						break
					}
				}
				// isExec = true

				if cfg.Call != nil {
					newSh(strings.Join(*cfg.Call, ";")).cmd.Wait()
				}
			} else {
				i := 0;for {
					if len(shCmdList) == 0 {
						break
					}
						fmt.Print(rls[ver])
						for k,v := range rls[ver] {
							con.WriteMessage(1, []byte("{\""+k+"\":\""+c.Query("id")+"--->"+v[i]+"\"}"))
						}
						//
					/*if len(shCmdList) != 0 {
						continue
					}*/i ++
				}
			}
		} else {
			c.HTML(200, "default.htm", gin.H{"tstmp": time.Now().Unix()})
		}
	})
	r.Run(":8080")
}

type appConfig struct {
	Ready *[]string
	Sh    map[string][]string
	Call  *[]string
}

func init() {
	cfg = Yml[appConfig]("c.yml")
}