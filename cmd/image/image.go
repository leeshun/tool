package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/leeshun/tool/ktqueue/image"

	"github.com/naturali/mgo"
)

var (
	mongoAddr = flag.String("mongo-addr", "127.0.0.1:27017", "mongo address")
	imageName = flag.String("image-name", "", "image name")
)

func main() {
	flag.Parse()
	sess, err := mgo.DialWithTimeout(*mongoAddr, time.Duration(10)*time.Second)
	if err != nil {
		fmt.Printf("connect mongo with address %v err: %v", mongoAddr, *mongoAddr)
		os.Exit(-1)
	}
	ks := image.NewKtQueueImageService(sess)
	ks.AddImage(*imageName)
}
