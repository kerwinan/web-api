package client

import (
	"fmt"
	"os"

	myoss "dana-tech.com/web-api/oss"
	"github.com/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
)

// OssClient ...
var OssClient *oss.Client

// NewOssClient ...
func NewOssClient() {
	OssClient, err := oss.New(myoss.Bucket.EndPoint, myoss.Bucket.AccessKeyID, myoss.Bucket.AccesskeySecret)
	if err != nil {
		beego.Error("New OssClient err: ", err.Error())
		os.Exit(-1)
	}
	isExist, err := OssClient.IsBucketExist("sher-bucket")
	if err != nil {
		beego.Error("bucket not exist: ", err.Error())
		os.Exit(-1)
	}
	fmt.Println(isExist)
}
