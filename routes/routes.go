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
	apiPolicy := r.Group("/api/documents")
	{
		apiPolicy.GET("/getPolicy", controllers.GetPolicy)
		apiPolicy.GET("/getPrivacy", controllers.GetPrivacy)
	}
	apiUser := r.Group("/api/user")
	{
		apiUser.GET("/getUserInfo", controllers.GetUserInformation)
		apiUser.PUT("/updateUserInfo", controllers.UpdateUser)
	}
	apiNotes := r.Group("/api/notes")
	{
		apiNotes.POST("/createNote", controllers.CreateNote)
		apiNotes.GET("/getAllNotes", controllers.GetAllNotes)
		apiNotes.GET("/getNote/:id", controllers.GetNoteByID)
		apiNotes.PUT("/updateNote/:id", controllers.UpdateNote)
		apiNotes.DELETE("/deleteNote/:id", controllers.DeleteNote)
	}
}
