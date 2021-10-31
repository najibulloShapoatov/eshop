package v1

import (
	"eshop/internal/service"

	_ "eshop/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	apiV1 := router.Group("/api/v1")

	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1.POST("/sign-up", h.signUp)
	apiV1.POST("/sign-in", h.signIn)

	authenticated := apiV1.Group("/auth", h.userIdentity)
	{
		authenticated.GET("/products", h.getAllProducts)
		authenticated.GET("/products/:id", h.getProductById)
		manager := authenticated.Group("/manager", h.isManager)
		{
			manager.POST("/products", h.createProduct)
			manager.PUT("/products/:id", h.updateProduct)
			manager.DELETE("/products/:id", h.deleteProduct)
			manager.GET("/carts", h.getAllCarts)
		}

		user := authenticated.Group("/user", h.isUser)
		{
			user.GET("/cart", h.getCart)
			user.POST("/cart/product", h.productToCart)
		}
	}

	return router
}
