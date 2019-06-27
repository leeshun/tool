package image

import (
	"fmt"
	"strings"

	"github.com/naturali/mgo"
	"github.com/naturali/mgo/bson"
)

const (
	hostName = "in.fds.so:5000"
)

type KtQueueImageService struct {
	sess *mgo.Session
}

func NewKtQueueImageService(sess *mgo.Session) *KtQueueImageService {
	return &KtQueueImageService{
		sess: sess,
	}
}

func (kis *KtQueueImageService) AddImage(imageName string) {
	tokens := strings.Split(imageName, "/")
	if len(tokens) != 2 {
		fmt.Printf("invalid image name, it should be %q\n",
			fmt.Sprintf("%v/%v:%v", hostName, "imageName", "imageTag"))
		return
	}
	if tokens[0] != hostName {
		fmt.Printf("invalid image name, image name should begin with %v\n", hostName)
		return
	}
	imageTokens := strings.Split(tokens[1], ":")
	if len(imageTokens) != 2 {
		fmt.Printf("invalid image name, it should end with %v\n",
			fmt.Sprintf("%v:%v", "imageName", "imageTag"))
		return
	}
	kis.addImageIntoMongo(imageTokens[0], imageTokens[1])
}

func (kis *KtQueueImageService) addImageIntoMongo(name, tag string) {
	s := kis.sess.Copy()
	defer s.Close()
	c := s.DB("ktqueue").C("images")
	account, err := c.Find(bson.M{"name": name, "tags": tag}).Count()
	if err != nil {
		fmt.Printf("find data err: %v\n", err)
		return
	}
	if account != 0 {
		fmt.Printf("the image already exist\n")
		return
	}
	info, err := c.Upsert(bson.M{"name": name}, bson.M{"$push": bson.M{"tags": tag}})
	if err != nil {
		fmt.Printf("insert data into mongo err: %v\n", err)
		return
	}
	fmt.Printf("insert data result: matched=%d, removed=%d, updated=%d\n",
		info.Matched, info.Removed, info.Updated)
}
