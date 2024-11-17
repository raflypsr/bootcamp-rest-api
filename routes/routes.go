package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fp-super-bootcamp-go/controllers"
	"fp-super-bootcamp-go/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/user", controllers.GetAllUser)
	r.GET("/user/:id", controllers.GetUserById)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("user/:id", controllers.DeleteUser)

	bookMiddlewareRoute := r.Group("/book")
	bookMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	bookMiddlewareRoute.POST("", controllers.CreateBook)
	bookMiddlewareRoute.PATCH("/:id", controllers.UpdateBook)
	bookMiddlewareRoute.DELETE("/:id", controllers.DeleteBook)

	r.GET("/book", controllers.GetAllBook)
	r.GET("/book/:id", controllers.GetBookById)

	reviewMiddlewareRoute := r.Group("/review")
	reviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewMiddlewareRoute.POST("", controllers.CreateReview)
	reviewMiddlewareRoute.PATCH("/:id", controllers.UpdateReview)
	reviewMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

	r.GET("/review", controllers.GetAllReview)
	r.GET("/review/:id", controllers.GetReviewById)

	profileMiddlewareRoute := r.Group("/profile")
	profileMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	profileMiddlewareRoute.POST("", controllers.CreateProfile)
	profileMiddlewareRoute.PATCH("/:id", controllers.UpdateProfile)
	profileMiddlewareRoute.DELETE("/:id", controllers.DeleteProfile)

	r.GET("/profile", controllers.GetAllProfile)
	r.GET("/user/:id/profile", controllers.GetProfileByUserId)

	likeMiddlewareRoute := r.Group("/like")
	likeMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	likeMiddlewareRoute.POST("", controllers.CreateLike)
	likeMiddlewareRoute.DELETE("/:id", controllers.DeleteLike)

	r.GET("/like", controllers.GetAllLike)

	commentMiddlewareRoute := r.Group("/comment")
	commentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	commentMiddlewareRoute.POST("", controllers.CreateComment)
	commentMiddlewareRoute.PATCH("/:id", controllers.UpdateComment)
	commentMiddlewareRoute.DELETE("/:id", controllers.DeleteComment)

	r.GET("/comment", controllers.GetAllComment)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
