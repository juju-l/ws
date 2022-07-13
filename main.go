package main

import (
	// "runtime"
  "github.com/gin-gonic/gin"
  "time"
)

var cfg *appConfig

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("./*.htm")
      ws := newWs()
  r.GET("/", func(c *gin.Context) {
    if c.Query("id") != "" {
      ws.run(cfg).cliRegister(c.Query("id"), c.Writer, c.Request).sendMsg(c.Query("id"))
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
  cfg = Yml[appConfig](
      "c.yml",
    )
}