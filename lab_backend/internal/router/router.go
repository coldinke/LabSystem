package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "lab_backend/docs"
	"lab_backend/internal/middleware"
	"lab_backend/internal/services"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Ping
	r.GET("/ping", services.Pong)

	// admin api
	admin := r.Group("/admin/api")
	{
		admin.POST("/login", services.Login)

		admin.Use(middleware.JWT())
		{
			// Count
			admin.GET("/users/count", services.GetUsersCount)
			admin.GET("/labs/count", services.GetLabsCount)
			admin.GET("/books/count", services.GetBooksCount)

			// Right
			admin.GET("/rights", services.GetRightsList)
			admin.PUT("/rights/:id", services.UpdateRight)
			admin.DELETE("/rights/:id", services.DeleteRight)
			//admin.POST("/rights", services.NewRole)

			// Role
			admin.GET("/roles", services.GetRolesList)
			admin.PUT("/roles/:id", services.UpdateRole)
			admin.DELETE("/roles/:id", services.DeleteRole)
			//admin.POST("/roles", services.NewRole)

			// User
			admin.GET("/users", services.GetUsersList)
			admin.PUT("/users/:id", services.UpdateUser)
			admin.DELETE("/users/:id", services.DeleteUser)
			admin.POST("/users", services.NewUser)

			// Lab
			admin.GET("/labs", services.GetLabsList)
			admin.GET("/labs/:id", services.GetLabDetail)
			admin.PUT("/labs/:id", services.UpdateLab)
			admin.DELETE("/labs/:id", services.DeleteLab)
			admin.POST("/labs", services.NewLab)

			// Book
			admin.GET("/books", services.GetBookList)
			admin.GET("/books/process", services.GetBookListByStatus)
			admin.GET("/books/:username", services.GetBookListByUsername)
			admin.GET("/select/books/:id", services.GetBookListByLabID)
			admin.POST("/books", services.NewBook)
			admin.DELETE("/books/:id", services.DeleteBook)
			admin.PUT("/books/:id", services.UpdateBook)

			// Sensor
			admin.GET("/sensor/:labID", services.GetLastSensor)

			// Devices
			admin.POST("/devices", services.PublishMessage)
		}

	}

	return r
}
