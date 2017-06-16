package ctr

import (
	"fmt"
	"lib/uweb"
	//"mod"
)

func init() {
	// simple get
	uweb.Get("/login", func(c *uweb.Context) {
		fmt.Println("login..............")
		c.Render.Html(200, "login", uweb.Map{
			"title": "hello demo1",
		})
	})
	uweb.Get("/lcq123", func(c *uweb.Context) {
		fmt.Println("lcq123.............")
		c.Render.Html(200, "lcq123", uweb.Map{
			"title": "hello lcq123",
		})
	})
}
func init() {
	// simple get
	/*	uweb.Get("/login", func(c *uweb.Context) {
			data := map[string]string{
				"key": "value",
			}
			c.Render.Html(200, "login", data)
		})
		/*
				 // post
				 uweb.Post("/api/login/", func(c *uweb.Context) {
				 	c.Render.Json(201, uweb.Map{
					  "key1": "value1",
					})
				 })

				 // not support regexp match
				 uweb.Put("/account/:user_id", func (c *uweb.Context) {
				     userId := c.Req.Params["user_id"]
				 	 println(userId)
				 	 account.Noop(userId)
				 	 c.Res.Plain(201, "success")
			     })
	*/
}
