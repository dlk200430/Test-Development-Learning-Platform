package routes

import (
	"github.com/testdev-learning/DataHub/backend/controllers"
	"github.com/testdev-learning/DataHub/backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.GetProfile)

		testCases := protected.Group("/testcases")
		{
			testCases.POST("", controllers.CreateTestCase)
			testCases.GET("", controllers.GetTestCases)
			testCases.GET("/:id", controllers.GetTestCase)
			testCases.PUT("/:id", controllers.UpdateTestCase)
			testCases.DELETE("/:id", controllers.DeleteTestCase)
		}

		exec := protected.Group("/exec")
		{
			exec.POST("/testcase/:id", controllers.ExecuteTestCase)
		}

		reports := protected.Group("/reports")
		{
			reports.GET("", controllers.GetTestReports)
			reports.GET("/:id", controllers.GetTestReport)
		}
	}

	return r
}