package main

import (
	"os"
	"time"

	"github.com/jasonsoft/napnap"
)

type Payload struct {
	Message string `json:"message"`
}

func newPayLoad() Payload {
	return Payload{
		Message: "Hello, World!",
	}
}

func getHelloWorldEndpoint(c *napnap.Context) {
	payLoad := newPayLoad()

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(200, payLoad)
		return
	}

	c.String(200, payLoad.Message)
}

func getHostnameEndpoint(c *napnap.Context) {
	hostname, _ := os.Hostname()
	c.String(200, hostname)
}

func throwInternalErrorEndpoint(c *napnap.Context) {
	c.String(500, "internal error/發生未知錯誤")
}

func throwBadRequestEndpoint(c *napnap.Context) {
	c.String(400, "bad request/找不到")
}

func timeoutEndpoint(c *napnap.Context) {
	time.Sleep(3 * time.Minute)
}