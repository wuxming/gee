package main

import (
	"log"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		var color string
		switch c.StatusCode / 100 {
		case 2:
			color = green
		case 3:
			color = white
		case 4:
			color = yellow
		default:
			color = red
		}
		log.Printf("["+color+"%d\033[0m] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
