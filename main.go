package main

import (
	"github.com/gin-gonic/gin"
)

	var ws *wsEngine

type tst struct {
	//
}

func main() {
	r := gin.Default()
			r.LoadHTMLGlob("./*.htm")
					ws = NewEngineWs(Yml[shCmd]("c.yml").Sh)
			r.GET("/", func(c *gin.Context) {
					if c.Query("id") != "" {
						ws.CliRegister(c.Query("id"), c.Writer, c.Request).RevBroadcastMsg()
					} else {
						c.HTML(200, "default.htm", nil)
					}
			})
					///ws.DestroyWs()
	r.Run(":8080")
}

type shCmd struct {
	Ready *[]string
	Sh map[string][]string
	Call *[]string
}

func init() {
	//
}