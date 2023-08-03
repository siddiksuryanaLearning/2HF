package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"2hf/controllers"
	"2hf/middlewares"
	"net/http"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Hunting For Halal Food"})
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", HomepageHandler)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", controllers.GetAllUserFromCurrentUser)

	r.GET("/advertise", controllers.GetAllAdvertise)
	r.GET("/advertise/:id", controllers.GetAdvertiseById)
	r.GET("/advertise-current-user", controllers.GetAllAdvertiseFromCurrentUser)
	advertiseMiddlewareRoute := r.Group("/advertise")
	advertiseMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	advertiseMiddlewareRoute.POST("/", controllers.CreateAdvertise)
	advertiseMiddlewareRoute.PATCH("/:id", controllers.UpdateAdvertise)
	advertiseMiddlewareRoute.DELETE("/:id", controllers.DeleteAdvertise)

	r.GET("/payment", controllers.GetAllPayment)
	r.GET("/payment/:id", controllers.GetPaymentById)
	r.GET("/payment-current-user", controllers.GetAllPaymentFromCurrentUser)
	paymentMiddlewareRoute := r.Group("/payment")
	paymentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	paymentMiddlewareRoute.POST("/", controllers.CreatePayment)
	paymentMiddlewareRoute.PATCH("/:id", controllers.UpdatePayment)
	paymentMiddlewareRoute.DELETE("/:id", controllers.DeletePayment)

	r.GET("/vocation", controllers.GetAllVocation)
	r.GET("/vocation/:id", controllers.GetVocationById)
	r.GET("/vocation-current-user", controllers.GetAllVocationFromCurrentUser)
	vocationMiddlewareRoute := r.Group("/vocation")
	vocationMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	vocationMiddlewareRoute.POST("/", controllers.CreateVocation)
	vocationMiddlewareRoute.PATCH("/:id", controllers.UpdateVocation)
	vocationMiddlewareRoute.DELETE("/:id", controllers.DeleteVocation)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
