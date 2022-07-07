package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"os"
	"os/exec"
	"sync"
	"time"
	"strings"
	"fmt"
)

	var wss = websocket.Upgrader{}

func run(shCmd string) *bufio.Reader {
    cmd := exec.Command("sh", /*"-x",*/ "-c", shCmd+";echo -n \\#", "2>&1"); stdout,_ := cmd.StdoutPipe(); bufReader := bufio.NewReader(stdout); cmd.Stderr = os.Stderr; go cmd.Run(); return bufReader //
}

var (
	mux sync.Mutex
	buflist map[string]*bufio.Reader
	msgChannel = make(map[string]chan /*interface{}*/[]byte)
	clilist = make(map[string]*websocket.Conn)
	sh *shCmd
)

func main() {
	r := gin.Default()
  r.LoadHTMLGlob("./*.htm")
	r.GET("/", func(c *gin.Context) {
      if c.Query("id") != "" {
			con := clilist[c.Query("id")]
			if con == nil {
			/*var err error;*/ con, _ = wss.Upgrade(c.Writer, c.Request, nil); clilist[c.Query("id")] = con
			}
			///defer con.Close()

			rls := Yml[map[string]map[string][]string]("r.yml")

			cha := msgChannel[c.Query("id")]
			if cha == nil {
			/*var err error;*/ cha = make(chan /*interface{}*/[]byte); msgChannel[c.Query("id")] = cha
			}
			///defer 销毁


			for {
				ver := time.Now().Format("v010206r")

				if r,e := (*rls)[ver]; e {
					for k,v := range r {
					  con.WriteMessage(1, []byte("{\""+k+"\":\""+strings.Join(v, "\\n")+"\"}"))
					}
				} else {
					if buflist == nil {
					buflist = make(map[string]*bufio.Reader) //
					for k,v := range *sh.Sh {
					buflist[k] = run(strings.Join(v, ";"))
					}
					//
					}
					for k,v := range buflist {
					  (*rls)[ver] = make(map[string][]string)
							for {
					/*select {
					  //
					  case */msg, ok := <- cha/*:*/
							err := con.WriteMessage(1, msg)
							if err != nil {
									if ok { go func() { /**/; cha <- msg } (); break } //offline
							}
							//
							if v == nil {
									break
							}
							log,_,_ := v.ReadLine()
							//
							if log != nil {
							if string(log) == "@" {
									/**/; delete(buflist, k) ;break
							}
							(*rls)[ver][k] = append(
									/*,*/ (*rls)[ver][k], string(log),
								)
							//
							cha <- []byte("{\""+k+"\":\""+string(log)+"\"}")
							//
							///con.WriteMessage(1, []byte("{\""+k+"\":\""+string(log)+"\"}"))
							//
							//time.Sleep(
							//		time.Millisecond*100,
							//	)
							} else {
									continue
							}
					  //
					/*}*/
							}
					///(*rls)[ver][k] = append((*rls)[ver][k], string(log))
					}
					//
					//
					if len(buflist) == 0 {
					break
					}
				}

				if len(buflist) == 0 {
					break
				}
			}


			buflist = nil
			delete(clilist, c.Query("id"))
			con.Close()

		  //Write("r.yml", *rls, /**/)
			fmt.Println("shCmd exec finished...")
			//

			///con, _ := wss.Upgrade(c.Writer, c.Request, nil); defer con.Close(); bufReader := run("packer --help"); for { log,_,_ := bufReader.ReadLine(); if string(log) == "#" { break }; con.WriteMessage(1, []byte("{\"consoleShow\":\""+string(log)+"\"}")) }; fmt.Println("shCmd exec finished...")
      } else {
        c.HTML(200, "default.htm", nil) //
      }
    })
  ///
	r.Run(":8080")
}

type shCmd struct {
  Ready *[]string
  Sh *map[string][]string
  Call *[]string
}

func init() {
    sh = Yml[shCmd](
        "c.yml",
      )
}