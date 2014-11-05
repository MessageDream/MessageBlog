package root

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)
import (
	"../../common"
)

import (
	"github.com/astaxie/beego"
)

const (
	EMPTYSTR = ""
	Script   = `<script type="text/javascript">window.parent.CKEDITOR.tools.callFunction("%d","%s","%v");</script>`
)

type RootUploadRouter struct {
	rootBaseRouter
}

func (this *RootUploadRouter) Post() {
	num, err := this.GetInt("CKEditorFuncNum")
	//	file, handler, err := this.GetFile("upload")
	file, _, err := this.GetFile("upload")
	if err != nil {
		beego.Error(err)
		return
	}

	defer file.Close()
	url, err := common.UploadFile("image", file)
	if err != nil {
		beego.Error(err)
		fmt.Fprintf(this.Ctx.ResponseWriter, Script, num, common.Webconfig.StaticURL+"/"+url, err)
	} else {
		beego.Error(err)
		fmt.Fprintf(this.Ctx.ResponseWriter, Script, num, common.Webconfig.StaticURL+"/"+url, "")
	}
	//	url, err := SaveFile("static/uploads", handler.Filename, file)//保存到本地
	// if err != nil {
	// 	fmt.Fprintf(this.Ctx.ResponseWriter, Script, num, common.Webconfig.SiteURL+"/"+url, err)
	// } else {
	// 	fmt.Fprintf(this.Ctx.ResponseWriter, Script, num, common.Webconfig.SiteURL+"/"+url, "")
	// }

}

//检查文件夹是否存在
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
	//panic("not reached")
}

//保存文件
func SaveFile(filePath, name string, file multipart.File) (string, error) {
	if !IsDirExists(filePath) {
		err := os.MkdirAll(filePath, 0666)
		if err != nil {
			beego.Error(err)
			return EMPTYSTR, err
		}
	}

	f, err := os.OpenFile(filePath+"/"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		beego.Error(err)
		return EMPTYSTR, err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		beego.Error(err)
		return EMPTYSTR, err
	}
	return filePath + "/" + name, err
}
