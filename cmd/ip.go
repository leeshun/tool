package main

import (
	"flag"
	"fmt"
	"time"

	ip2 "github.com/leeshun/tool/watch/ip"
)

var (
	val = flag.Int("time", 10, "watch time val")
)

func init() {
	flag.Parse()
	if *val <= 0 {
		*val = 5
	}
}

func main() {
	ticker := time.NewTicker(time.Duration(*val) * time.Second)
	for range ticker.C {
		go func() {
			info := ip2.GetIPInfo()
			if info != "" {
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Println(fmt.Sprintf("%v\n%v", timeStr, info))
			}
		}()
	}
}
