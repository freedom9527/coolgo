package loader

import (
	"coolgo/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//func init()  {
//	app := App{}
//	cfg := app.GetConfig()
//	app.GetDatabase(cfg)
//	fmt.Printf("haha: %v", cfg.GetString("mysql.db_name"))
//	fmt.Println()
//}

type App struct {
	Cfg *config.Config
	Db *gorm.DB
	Router *gin.Engine
}

func (a *App) SysInit() error  {
	cfg := a.GetConfig()
	err := a.GetDatabase(cfg)
	if err != nil {
		return err
	}
	return nil
}

// get system config
func (a *App) GetConfig() config.Config {
	if a.Cfg == nil {
		c := config.Config{}
		c.Load()
		a.Cfg = &c
	}
	return *a.Cfg
}
// init default database
func (a *App) GetDatabase( cfg config.Config) error  {
	var db *gorm.DB
	var err error
	driver := cfg.GetString("db_driver")
	user := cfg.GetString("db_name")
	pwd := cfg.GetString("db_pwd")
	host := cfg.GetString("db_host")
	port := cfg.GetString("db_port")
	name := cfg.GetString("db_name")
	char := cfg.GetString("db_charset")
	switch driver {
		case "mysql":
			connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, pwd, host, port, name, char)
			db, err = gorm.Open(driver, connectStr)
		case "sqlite3":
			db, err = gorm.Open(driver, name)
	}
	//defer db.Close()
	if err != nil {
		return err
	}
	a.Db = db
	return nil
}




