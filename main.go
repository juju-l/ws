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
	//rls := Yml[map[string]map[string][]string]("r.yml")
	r.GET("/", func(c *gin.Context) {
		if c.Query("id") != "" {
			con := wsConList[c.Query("id")]
			if con == nil {
				/*var err error;*/ con, _ = wssObject.Upgrade(c.Writer, c.Request, nil)
				wsConList[c.Query("id")] = con
			}
			defer delete(wsConList, c.Query("id"))

			//ver := time.Now().Format("v010206r")

			if cfg.Ready != nil {
				newSh(strings.Join(*cfg.Ready, ";")).cmd.Wait()
			}

			shCmdList := make(map[string]*shCmd)
			for k, v := range cfg.Sh {
				shCmdList[k] = newSh(strings.Join(v, ";"))
			}
			//
			for i, j := range shCmdList {
				go func(k string, v *shCmd) {
					i := 0
					for {
						if len(v.rst) > 0 {
							fmt.Println(v)
							if len(v.rst) == i {
								if v.isComplete {
									fmt.Println(shCmdList)
									delete(shCmdList, k)
									break
								} else {
									continue
								}
							}
							fmt.Printf("%v--->%s", i, v.rst[i])
							con.WriteMessage(1, []byte("{\""+k+"\":\""+c.Query("id")+"--->"+v.rst[i]+"\"}"))
							time.Sleep(time.Millisecond * 1000)
							i++
						}
					}
				}(i, j)
			}
			for {
				if len(shCmdList) == 0 {
					break
				}
			}
			//

			if cfg.Call != nil {
				newSh(strings.Join(*cfg.Call, ";")).cmd.Wait()
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
