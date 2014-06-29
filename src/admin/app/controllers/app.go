package controllers

//后台首页
import (
	"admin/app/models"
	"github.com/revel/config"
	"github.com/revel/revel"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type App struct {
	*revel.Controller
}

//首页
func (c App) Index(admin *models.Admin) revel.Result {

	revel.TRACE.Println("后台首页")

	title := "首页--HongID后台管理系统"

	if adminId, ok := c.Session["AdminID"]; ok {
		AdminID, err := strconv.ParseInt(adminId, 10, 64)
		if err != nil {
			revel.WARN.Printf("解析Session错误: %v", err)
		}

		admin_info := admin.GetById(AdminID)
		if admin_info.Id <= 0 {
			return c.Redirect("/Login/")
		}

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
		if c.Session["LockScreen"] == "" || c.Session["LockScreen"] == "0" {
			c.RenderArgs["LockScreen"] = "0"
		} else {
			c.RenderArgs["LockScreen"] = "1"
		}
	} else {
		return c.Redirect("/Login/")
	}

	c.Render(title)
	return c.RenderTemplate("App/Index.html")
}

func (c App) Main(admin *models.Admin) revel.Result {

	title := "首页--HongID后台管理系统"

	if adminId, ok := c.Session["AdminID"]; ok {
		AdminId, err := strconv.ParseInt(adminId, 10, 64)
		if err != nil {
			revel.WARN.Printf("session解析错误: %v", err)
		}

		admin_info := admin.GetById(AdminId)

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

		c.Render(title, admin_info, system_info, panel_list)
	} else {
		c.Render(title)
	}

	return c.RenderTemplate("App/Main.html")
}
