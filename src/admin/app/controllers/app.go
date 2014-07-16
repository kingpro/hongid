package controllers

//后台首页
import (
	"admin/app/models"
	"admin/utils/consts"
	"github.com/revel/config"
	"github.com/revel/revel"
	"os"
	"runtime"
	"strings"
)

type App struct {
	*revel.Controller
}

//首页
func (c *App) Index(admin *models.Admin) revel.Result {

	if admin_info, ok := GetAdminInfoBySession(c.Session); ok {

		//控制器
		c.RenderArgs["Controller"] = c.Name
		//动作
		c.RenderArgs["Action"] = c.Action
		//模型
		c.RenderArgs["Model"] = c.MethodName
		//导航菜单
		menu := new(models.Menu)
		c.RenderArgs["Menus"] = menu.GetAdminMenu(0, admin_info)
		//管理员信息
		c.RenderArgs["AdminInfo"] = admin_info
		//是否锁屏
		if c.Session[consts.C_Session_LockS] == "" || c.Session[consts.C_Session_LockS] == "0" {
			c.RenderArgs[consts.C_Session_LockS] = consts.C_Lock_0
		} else {
			c.RenderArgs[consts.C_Session_LockS] = consts.C_Lock_1
		}
	} else {
		return c.Redirect("/Login/")
	}

	return c.RenderTemplate("App/Index.html")
}

func (c *App) Main(admin *models.Admin) revel.Result {

	if admin_info, ok := GetAdminInfoBySession(c.Session); ok {

		//判断是否是系统的分隔符
		separator := "/"
		if os.IsPathSeparator('\\') {
			separator = "\\"
		}

		config_file := (revel.BasePath + "/conf/config.conf")
		config_file = strings.Replace(config_file, "/", separator, -1)
		config_conf, _ := config.ReadDefault(config_file)

		system_info := make(map[string]string)

		//版本
		version, _ := config_conf.String("website", "website.version")
		system_info["version"] = version

		//前台网站地址
		sitedomain, _ := config_conf.String("website", "website.sitedomain")
		system_info["sitedomain"] = sitedomain

		//操作系统
		system_info["os"] = strings.ToUpper(runtime.GOOS + " " + runtime.GOARCH)

		//Go版本
		system_info["go_varsion"] = strings.ToUpper(runtime.Version())

		//MySQL版本
		system_info["mysql_varsion"] = admin.GetMysqlVer()

		//快捷面板
		admin_panel := new(models.AdminPanel)
		panel_list := admin_panel.GetPanelList(admin_info)

		c.Render(admin_info, system_info, panel_list)
	} else {
		c.Render()
	}

	return c.RenderTemplate("App/Main.html")
}
