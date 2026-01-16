package main

import (
	"context"

	"github.com/LgoLgo/cqupt-grabber/cqupt"
)

func main() {
	tool := cqupt.New()

	cookie := "这里是一个cookie"

	// 支持同时抢多门课程
	loads := []string{
		"这里是第一节课",
		"这里是第二节课",
	}

	tool.Grabber.LoopRob(context.Background(), cookie, loads)
}
