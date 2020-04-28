package routers
import (
	"coolgo/controllers"
	"coolgo/loader"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *loader.App) *gin.Engine {
	r := gin.Default()

	r.GET("/:controller/index", (&controllers.AdminController{Db:app.Db}).IndexEndpoint)
	r.GET("/:controller/edit/:id", (&controllers.AdminController{}).EditEndpoint)
	r.GET("/:controller/delete/:id", (&controllers.AdminController{}).DeleteEndpoint)

	return r
}