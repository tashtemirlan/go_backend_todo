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
		apiLogin.POST("/validate-token", controllers.ValidateToken)
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
	apiTaskGroup := r.Group("/api/tasks_groups")
	{
		apiTaskGroup.POST("/createTaskGroup", controllers.CreateTaskGroup)
		apiTaskGroup.GET("/getTasksGroup", controllers.GetTaskGroups)
		apiTaskGroup.GET("/getTaskGroup/:id", controllers.GetTaskGroup)
		apiTaskGroup.PUT("/updateTaskGroup/:id", controllers.UpdateTaskGroup)
		apiTaskGroup.DELETE("/deleteTaskGroup/:id", controllers.DeleteTaskGroup)
	}
	apiTask := r.Group("/api/tasks")
	{
		apiTask.POST("/createTask", controllers.CreateTask)
		apiTask.GET("/getAllTasks", controllers.GetAllTasks)
		apiTask.GET("/getTask/:id", controllers.GetTask)
		apiTask.PUT("/updateTask/:id", controllers.UpdateTask)
		apiTask.DELETE("/deleteTask/:id", controllers.DeleteTask)
		apiTask.GET("/getTasks/todo", controllers.GetTasksByStatusTODO)
		apiTask.GET("/getTasks/in_progress", controllers.GetTasksByStatusInProgress)
		apiTask.GET("/getTask/finish-date", controllers.GetTasksByFinishDate)
	}
}
