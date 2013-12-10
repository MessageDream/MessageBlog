package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type KeyValuePair struct {
	Key   int64
	Value string
}

type KeyPairs struct {
	Items []*KeyValuePair
}

func (kp KeyPairs) Len() int {
	return len(kp.Items)
}

func (kp KeyPairs) Swap(i, j int) {
	kp.Items[i], kp.Items[j] = kp.Items[j], kp.Items[i]
}

func (kp KeyPairs) Less(i, j int) bool {
	if kp.Items[i].Key == kp.Items[j].Key {
		return kp.Items[i].Value < kp.Items[j].Value
	}
	return kp.Items[i].Key < kp.Items[j].Key
}

func UUID() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] & 0x0F) | 0x40
	b[8] = (b[8] &^ 0x40) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func TimeHumanReading(t1 int64) string {
	t := time.Unix(t1, 0).Local()
	return fmt.Sprintf("%v-%v-%v %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func TimeRFC3339(t1 int64) string {
	t := time.Unix(t1, 0).Local()
	return t.Format("2006-01-02T15:04:05Z")
}

func MD5Sum(in string) string {
	md5 := md5.New()
	md5.Write([]byte(in))
	return fmt.Sprintf("%x", md5.Sum(nil))
}

func SHA256Sum(in string) string {
	sha := sha256.New()
	sha.Write([]byte(in))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

/**
* user : example@example.com login smtp server user
* password: xxxxx login smtp server password
* host: smtp.example.com:port   smtp.163.com:25
* to: example@example.com;example1@163.com;example2@sina.com.cn;...
* subject:The subject of mail
* body: The content of mail
* mailtype: mail type html or text
 */
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func Encrypt_password(password string, salt []byte) string {
	if salt == nil {
		m := md5.New()
		m.Write([]byte(time.Now().String()))
		s := hex.EncodeToString(m.Sum(nil))
		salt = []byte(s[2:10])
	}
	mac := hmac.New(sha256.New, salt)
	mac.Write([]byte(password))
	//s := log.Sprintf("%x", (mac.Sum(salt)))
	s := hex.EncodeToString(mac.Sum(nil))

	hasher := sha1.New()
	hasher.Write([]byte(s))

	//result := log.Sprintf("%x", (hasher.Sum(nil)))
	result := hex.EncodeToString(hasher.Sum(nil))

	p := string(salt) + result

	return p
}

func Validate_password(hashed string, input_password string) bool {
	salt := hashed[0:8]
	if hashed == Encrypt_password(input_password, []byte(salt)) {
		return true
	} else {
		return false
	}
	return false
}

func GetPager(dir string, current, total int) template.HTML {
	if total == 1 || total == 0 {
		return template.HTML("")
	}
	url := Webconfig.SiteURL + "/" + dir + "/"
	var pre, next string
	first := strconv.Itoa(current - 1)
	if current == 1 {
		pre = ""
	} else {
		pre = `<a href="` + url + first + `"><i class="icon-left-open-1"></i> 上一页</a>`

		if current-2 >= 2 {
			pre += `<a class="box" href="` + Webconfig.SiteURL + `">«</a>`
		}
		pre += `<a class="box" href="` + url + first + `">‹</a>`
		if current-2 >= 1 {
			third := strconv.Itoa(current - 2)
			pre += `<a class="box" href="` + url + third + `">` + third + `</a>`
		}
		if current-1 >= 1 {
			second := strconv.Itoa(current - 1)
			pre += `<a class="box" href="` + url + second + `">` + second + `</a>`
		}
	}

	cur := `<span class="current box">` + strconv.Itoa(current) + `</span>`

	if current != total {
		nfirst := strconv.Itoa(current + 1)
		next = `<a href="` + url + nfirst + `" class="inactive box">` + nfirst + `</a>`
		if total > current+1 {
			nsecond := strconv.Itoa(current + 2)
			next += `<a href="` + url + nsecond + `" class="inactive box">` + nsecond + `</a>`
		}
		next += `<a class="box" href="` + url + nfirst + `">›</a>`
		if total > current+2 {
			next += `<a class="box" href="` + url + strconv.Itoa(total) + `/">»</a>`
		}
		next += `<a href="` + url + nfirst + `">下一页 <i class="icon-right-open-1"></i></a>`
	}
	return template.HTML(pre + cur + next)
}
