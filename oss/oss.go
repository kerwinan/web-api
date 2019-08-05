package oss

import (
	"github.com/dlintw/goconf"
)

// AK ...
type AK struct {
	EndPoint        string
	AccessKeyID     string
	AccesskeySecret string
}

// Bucket ...
var Bucket *AK

// Init ...
func Init(conf *goconf.ConfigFile) {
	Bucket = &AK{}
	Bucket.EndPoint, _ = conf.GetString("oss", "end_point")
	Bucket.AccessKeyID, _ = conf.GetString("oss", "access_key_id")
	Bucket.AccesskeySecret, _ = conf.GetString("oss", "access_key_secret")
	// c.OssClient, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	// OssClient, err = oss.New(endPoint, accessKeyID, accessKeySecret)
	// if err != nil {
	// 	beego.Error("New OssClient error: ", err.Error())
	// 	os.Exit(-1)
	// }
}
