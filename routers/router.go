package routers

import (
	_ "kusnandartoni/starter/docs" //swager files
	"kusnandartoni/starter/midleware/jwt"
	"kusnandartoni/starter/pkg/setting"
	v1 "kusnandartoni/starter/routers/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter :
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/", v1.HealthCheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/auth/login", v1.Login)
	r.GET("/api/auth/get-token", v1.GetToken)
	r.POST("/api/auth/forgot", v1.Forgot)
	r.GET("/api/auth/verify", v1.Verify)
	r.PUT("/api/auth/reset", v1.Reset)
	r.POST("/api/auth/register", v1.Register)

	apiV1 := r.Group("/api/v1")
	{
		class := apiV1.Group("/class")
		class.Use(jwt.JWT())
		{
			class.GET("", v1.GetClasses)
			class.POST("", v1.AddClass)
			class.PUT(":id", v1.EditClass)
			class.DELETE(":id", v1.DeleteClass)
		}

	}

	return r
}
