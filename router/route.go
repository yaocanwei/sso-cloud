/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaocanwei/sso-cloud/controller"
	"github.com/yaocanwei/sso-cloud/middleware"
)

func RegisterRoutes(mids ...gin.HandlerFunc) (r *gin.Engine, err error) {
	r = gin.Default()
	r.Use(gin.Recovery())
	r.Use(mids...)
	r.Use(middleware.AccessLogging(r))

	public := r.Group("/api")
	{
		public.GET("/login", controller.Login)
		public.GET("/validate", controller.Validate)
		public.GET("logout", controller.Logout)
	}

	return
}
