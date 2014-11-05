package routers

const NOT_FOUND_MESSAGE = "Sorry, the page you were looking does not exist."

type PromptRouter struct {
	baseRouter
}

func (this *PromptRouter) Get() {
	code := this.Ctx.Input.Param(":code")
	if code == "404" {
		vars := make(map[string]interface{})
		vars["Message"] = NOT_FOUND_MESSAGE
		data := MakeData(vars)
		this.Data["Data"] = data
		this.TplNames = "404.html"
	}
}
