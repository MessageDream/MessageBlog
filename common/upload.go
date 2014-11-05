package common

import (
	"bytes"
	"crypto/md5"
	//"log"
	. "github.com/qiniu/api/conf"
	qiniu_io "github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"io"
	"io/ioutil"
	"strings"
)

func UploadFile(dir string, read io.Reader) (string, error) {
	ACCESS_KEY = Webconfig.UploadConfig.QiniuAccessKey
	SECRET_KEY = Webconfig.UploadConfig.QiniuSecretKey

	extra := &qiniu_io.PutExtra{
		//	Bucket:         "messagedream",
		MimeType: "",
	}

	var policy = rs.PutPolicy{
		Scope: "messagedream",
	}

	filename := strings.Replace(UUID(), "-", "", -1) + ".jpg"
	body, _ := ioutil.ReadAll(read)

	h := md5.New()
	h.Write(body)

	key := "static/upload/" + dir + "/" + filename
	ret := new(qiniu_io.PutRet)

	buf := bytes.NewBuffer(body)

	return key, qiniu_io.Put(nil, ret, policy.Token(nil), key, buf, extra)
}
