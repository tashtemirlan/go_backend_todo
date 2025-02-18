package routes

import (
	"github.com/gin-gonic/gin"
	"material_todo_go/controllers"
)

func SetupRoutes(r *gin.Engine) {
	apiLogin := r.Group("/api/auth")
	{
		apiLogin.POST("/login", controllers.Login)
		apiLogin.POST("/signup", controllers.Signup)
		apiLogin.POST("/forget-password/generateCode", controllers.SendResetCode)
		apiLogin.POST("/forget-password/changePassword", controllers.ResetPassword)
	}
	apiUser := r.Group("/api/user")
	{
		apiUser.GET("/getUserInfo", controllers.GetUserInformation)
		apiUser.PUT("/updateUserInfo", controllers.UpdateUser)
	}
}
