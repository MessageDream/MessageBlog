package common

type Config struct {
	// sys config
	Port        int
	Certificate string
	Path        string
	SiteBase    string
	SiteURL     string
	// site config
	SiteTitle     string
	SiteSubTitle  string
	AuthorName    string
	TimelineCount int
	ThemeName     string
}

var Webconfig *Config = nil

func init() {
	Webconfig = new(Config)
	// default config
	Webconfig.Port = 8888
	//config.Certificate = SHA256Sum("42")
	Webconfig.SiteBase = "localhost"
	Webconfig.SiteURL = "http://localhost:8888"
	Webconfig.SiteTitle = "TATTOO!"
	Webconfig.Path = "/"
	Webconfig.AuthorName = "root"
	Webconfig.TimelineCount = 3
	Webconfig.ThemeName = "sealscript"
}
