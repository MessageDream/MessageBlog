MessageBlog
===========
[![Requirement >= Go 1.2](http://b.repl.ca/v1/Requirement-%3E%3D_Go%201.2-blue.png)]() [![Requirement >= beego 1.0.0](http://b.repl.ca/v1/Requirement-%3E%3D_beego%201.0.0-blue.png)]()

An open source blog project which powered by golang beego and mongodb and you can view the demo here:http://buxuxiao.com

## Install site locally

MessageBlog is a `go get` able project:

	$ go get github.com/MessageDream/MessageBlog

Install the dependent packages:

	$ go get github.com/astaxie/beego
	$ go get labix.org/v2/mgo

Switch to project root path:

	$ cd $GOPATH/src/github.com/MessageDream/MessageBlog/v2

Build and run with Go tools:

	$ go build blog.go
	$ ./blog

Open your browser and visit [http://localhost:8888](http://localhost:8888).

## Build as your site

This project can be easily transferred as your own blog site, there are some tips that you may want to know:

- In the file `conf/app.conf`:

-`appname = 北飘漂`:The title of your blog

-`httpport = 8888`:The port of your web site

-`dbconn = localhost:27017`:The mongodb connection string

-`logfile = logs/logs.log`:The log file path

-`username = admin`:The username of backstage

-`password = 123456`:The password of backstage

-`emailserver = smtp.163.com:25`:The email server of sending the notification

-`emailsender =admin@163.com`:The email sender address of sending the notification

-`emailpwd = admin`:The email sender's password

-`sitebase = localhost`:The IP address

-`siteurl = http://localhost:8888`:The domain name

-`staticurl = http://localhost:8888`:The server address of static File 

-`themename = bliss`:The theme name

-`pagecount = 5`:pagecount

-`qiniuaccesskey = *****************`:Your qiniu qiniuaccesskey

-`qiniusecretkey = *****************`:Your qiniu qiniusecretkey
