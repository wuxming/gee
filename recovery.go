package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	//跳过前三个，Callers、trace、Recovery这三个函数
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		//该函数的函数名，文件名和行数
		str.WriteString(fmt.Sprintf("\n\t%s:  %s:%d", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail("Internal Server Error")
			}
		}()
		c.Next()
	}
}
