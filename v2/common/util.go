package common

import (
	"bytes"
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

import (
	"github.com/astaxie/beego"
)

type KeyValuePair struct {
	Key   int64
	Value string
}

type KeyPairs struct {
	Items []*KeyValuePair
}

var SubscribeEmailContent string = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd">
<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta content="IE=11.0000"
          http-equiv="X-UA-Compatible">
    <title>
    </title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="robots" content="noindex, nofollow, noarchive">
    <meta name="goolgebot" content="noindex, nofollow, noarchive">
    <style type="text/css">
        body {
            margin: 0px;
            padding: 0px;
            font-size: 11px;
            font-family: Verdana,Sans-Serif,Tahoma,helviteca;
            background-color: #f8f8f8;
            overflow: visible;
            width: 100%;
            height: 100%;
        }

        .bl_author_img {
            text-align: center;
            left: 50%;
            margin-top: 35px;
            top: -45px;
            overflow: hidden;
            -moz-border-radius: 500px;
            -webkit-border-radius: 500px;
            border-radius: 500px;
            width: 90px;
            background: #fff;
            padding: 5px;
        }

        .bl_author_bio {
            padding-left: 25px;
            padding-right: 25px;
            padding-bottom: 15px;
            color: #333;
        }

        .widget-body .bl_author_img img {
            -moz-border-radius: 500px;
            -webkit-border-radius: 500px;
            border-radius: 500px;
            height: 90px;
            max-width: none;
            width: 90px;
        }
    </style>

    <meta name="GENERATOR" content="MSHTML 11.00.9600.16476">
</head>
<body>
    <div id="divEmailBody" style="padding: 10px; width: 100%; height: 100%; overflow: auto; font-size: 12px;">
        <div id="divEmail">
            <center>
                <div style="margin: 0px auto; padding: 0px; width: 705px; height: auto;">
                    <h1 style="background: rgb(199, 199, 199); margin: 0px; padding: 0px; border-radius: 5px 5px 0px 0px; height: 35px; color: rgb(255, 255, 255); line-height: 35px; font-size: 14px;">感谢您一直以来的支持</h1>
                    <h2 style="background: rgb(41, 43, 40); margin: 0px; padding: 0px; height: 100px; text-align: center; line-height: 120px; font-size: 36px;">
                        <a style="text-decoration: none;" href="{{.WebUrl}}">
                            <font style="color: rgb(255, 0, 102);">北</font><font style="color: rgb(255, 255, 102);">飘</font><font style="color: rgb(102, 255, 204);">漂</font>
                        </a>
                    </h2>
                    <h1 style="background: rgb(41, 43, 40); height: 50px; margin: 0px; padding: 0px; text-align: center; font-size: 12px; ">
                        <small>
                            <font style="color: rgb(255, 0, 102);">
                                一个乐于分享| .Net | C/C++ | Golang | TypeScript | NoSql | linux |等技术的IT博客
                            </font>
                        </small>
                    </h1>
                    <div style="margin: 0px; padding: 0px; width: 705px; height:auto; overflow: hidden;">
                        <div style="background: rgb(240, 240, 240); margin: 0px; padding: 0px; width: 187px; height: auto; text-align: center; float: left;">
                            <div class="widget-body">
                                <center>
                                    <div class="bl_author_img">
                                        <img src="{{.StaticUrl}}/static/bliss/img/messagedream.jpg">
                                    </div>
                                </center>
                                <div class="bl_author_bio">
                                    <h3>MessageDream</h3>
                                    <p class="muted">敢于尝试新东西，勇于面对新挑战。</p>
                                    <span><a href="http://weibo.com/u/2017203357" target="_blank"><span class="sprite-round-icon-sina">新浪微博</span></a></span><br />
                                    <span><a href="http://t.qq.com/messagedream" target="_blank"><span class="sprite-round-icon-qq">腾讯微博</span></a></span>
                                </div>
                                <p></p>
                            </div>
                        </div><div style="background: rgb(255, 255, 255); margin: 0px; padding: 0px; width: 518px; height: auto; float: right;">
                            <h3 style="margin: 0px;">&nbsp; </h3>
                            <h3 style="margin: 0px;">
                                <font color="#808080">
                                    <a style="text-decoration: none;"
                                       href="{{.WebUrl}}/article/{{.ArticleName}}">
                                        &nbsp;
                                        <font color="#808080">{{.Title}}</font>
                                    </a>
                                </font>
                            </h3>
                            <h3 style="margin: 0px;">&nbsp;</h3>
                            <div style="margin: 0px auto 15px; padding: 0px 0px 15px; width: 415px; overflow: hidden; border-bottom-color: rgb(176, 176, 176); border-bottom-width: 1px; border-bottom-style: dotted;">
                                {{.Summary}} 
                                <a style="FONT-SIZE: 12px; TEXT-DECORATION: none; FONT-FAMILY: Microsoft yahei;" target="_blank" href="{{.WebUrl}}/article/{{.ArticleName}}">
                                 <font color="#2f82de">查看全部>></font>
                                </a>
                            </div>
                            
                        </div>
                    </div>
                    <div style="margin: 0px; padding: 25px 0px 0px; border-radius: 0px 0px 6px 6px; width: 705px; height: 90px; color: rgb(237, 237, 237); background-color: rgb(199, 199, 199); -moz-border-radius: 0px 0px 6px 6px; -webkit-border-radius: 0px 0px 6px 6px;">
                        <p style="margin: 0px; padding: 0px; text-align: center; color: rgb(237, 237, 237); line-height: 25px; font-family: Microsoft yahei; font-size: 12px;">感谢您阅读本期邮件，若不想收到此邮件，请点击<a style="margin: 0px 5px; color: rgb(255, 255, 255); font-family: Microsoft yahei; font-size: 12px; text-decoration: none;" href="{{.WebUrl}}/desubscribe/{{.Uid}}">退订</a></p>
                        <p style="margin: 0px; padding: 0px; text-align: center; line-height: 25px;">
                            <a style="color: rgb(237, 237, 237); font-family: Microsoft yahei; font-size: 12px; text-decoration: none;"
                               href="mailto:{{.MailSender}}">mailto:{{.MailSender}}</a>
                        </p>
                    </div>
                </div>
            </center>
        </div>
    </div>
</body>
</html>
`
var SubConfirmContent string = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd">
<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta content="IE=11.0000"
          http-equiv="X-UA-Compatible">
    <title>北飘漂博客订阅确认</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="robots" content="noindex, nofollow, noarchive">
    <meta name="goolgebot" content="noindex, nofollow, noarchive">
    <style type="text/css">
        body {
            margin: 0px;
            padding: 0px;
            font-size: 11px;
            font-family: Verdana,Sans-Serif,Tahoma,helviteca;
            background-color: #f8f8f8;
            overflow: visible;
            width: 100%;
            height: 100%;
        }

        .bl_author_img {
            text-align: center;
            left: 50%;
            margin-top: 35px;
            top: -45px;
            overflow: hidden;
            -moz-border-radius: 500px;
            -webkit-border-radius: 500px;
            border-radius: 500px;
            width: 90px;
            background: #fff;
            padding: 5px;
        }

        .bl_author_bio {
            padding-left: 25px;
            padding-right: 25px;
            padding-bottom: 15px;
            color: #333;
        }

        .widget-body .bl_author_img img {
            -moz-border-radius: 500px;
            -webkit-border-radius: 500px;
            border-radius: 500px;
            height: 90px;
            max-width: none;
            width: 90px;
        }
    </style>

    <meta name="GENERATOR" content="MSHTML 11.00.9600.16476">
</head>
<body>
    <div id="divEmailBody" style="padding: 10px; width: 100%; height: 100%; overflow: auto; font-size: 12px;">
        <div id="divEmail">
            <center>
                <div style="margin: 0px auto; padding: 0px; width: 705px; height: auto;">
                    <h1 style="background: rgb(199, 199, 199); margin: 0px; padding: 0px; border-radius: 5px 5px 0px 0px; height: 35px; color: rgb(255, 255, 255); line-height: 35px; font-size: 14px;"></h1>
                    <h2 style="background: rgb(41, 43, 40); margin: 0px; padding: 0px; height: 100px; text-align: center; line-height: 120px; font-size: 36px;">
                        <a style="text-decoration: none;" href="{{.WebUrl}}">
                            <font style="color: rgb(255, 0, 102);">北</font><font style="color: rgb(255, 255, 102);">飘</font><font style="color: rgb(102, 255, 204);">漂</font>
                        </a>
                    </h2>
                    <h1 style="background: rgb(41, 43, 40); height: 50px; margin: 0px; padding: 0px; text-align: center; font-size: 12px; ">
                        <small>
                            <font style="color: rgb(255, 0, 102);">
                                一个乐于分享| .Net | C/C++ | Golang | TypeScript | NoSql | linux |等技术的IT博客
                            </font>
                        </small>
                    </h1>
                    <div style="margin: 0px; padding: 0px; width: 705px; height:auto; overflow: hidden;">
                        <div style="background: rgb(240, 240, 240); margin: 0px; padding: 0px; width: 187px; height: auto; text-align: center; float: left;">
                            <div class="widget-body">
                                <center>
                                    <div class="bl_author_img">
                                        <img src="{{.StaticUrl}}/static/bliss/img/messagedream.jpg">
                                    </div>
                                </center>
                                <div class="bl_author_bio">
                                    <h3>MessageDream</h3>
                                    <p class="muted">敢于尝试新东西，勇于面对新挑战。</p>
                                    <span><a href="http://weibo.com/u/2017203357" target="_blank"><span class="sprite-round-icon-sina">新浪微博</span></a></span><br />
                                    <span><a href="http://t.qq.com/messagedream" target="_blank"><span class="sprite-round-icon-qq">腾讯微博</span></a></span>
                                </div>
                                <p></p>
                            </div>
                        </div><div style="background: rgb(255, 255, 255); margin: 0px; padding: 0px; width: 518px; height: auto; float: right;">
                            <h3 style="margin: 0px;">&nbsp; </h3>
                            <h3 style="margin: 0px;">
                                <font color="#808080">
                                    <a style="text-decoration: none;"
                                       href="{{.WebUrl}}">
                                        &nbsp;
                                        <font color="#808080">感谢您订阅北飘漂博客</font>
                                    </a>
                                </font>
                            </h3>
                            <h3 style="margin: 0px;">&nbsp;</h3>
                            <div style="margin: 0px auto 15px; padding: 0px 0px 15px; width: 415px; overflow: hidden; border-bottom-color: rgb(176, 176, 176); border-bottom-width: 1px; border-bottom-style: dotted;">
                                <p> 请点击以下链接确认:</p>
                                <p><a href="{{.WebUrl}}/subscribe/{{.Uid}}">{{.WebUrl}}/subscribe/{{.Uid}}</a></p>
                                <p> 若您没有发送过订阅需求，请忽略。</p>
                            </div>
                        </div>
                    </div>
                    <div style="margin: 0px; padding: 25px 0px 0px; border-radius: 0px 0px 6px 6px; width: 705px; height: 90px; color: rgb(237, 237, 237); background-color: rgb(199, 199, 199); -moz-border-radius: 0px 0px 6px 6px; -webkit-border-radius: 0px 0px 6px 6px;">
                        <p style="margin: 0px; padding: 0px; text-align: center; color: rgb(237, 237, 237); line-height: 25px; font-family: Microsoft yahei; font-size: 12px;">若您没有发送过订阅需求，请忽略。</p>
                        <p style="margin: 0px; padding: 0px; text-align: center; line-height: 25px;">
                            <a style="color: rgb(237, 237, 237); font-family: Microsoft yahei; font-size: 12px; text-decoration: none;"
                               href="mailto:{{.MailSender}}">mailto:{{.MailSender}}</a>
                        </p>
                    </div>
                </div>
            </center>
        </div>
    </div>
</body>
</html>
`

type SubscribeEmail struct {
	StaticUrl   string
	Title       string
	ArticleName string
	Summary     template.HTML
	WebUrl      string
	Uid         string
	MailSender  string
}

func PaseHtml(content string, data interface{}) *string {
	t, err := template.New("name").Parse(content)
	if err != nil {
		beego.Error(err)
		return nil
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	err = t.Execute(buf, data)
	if err != nil {
		beego.Error(err)
		return nil
	}

	ht := buf.String()
	return &ht
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
