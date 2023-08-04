package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kozhamseitova/ustaz-hub-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("/api/v1")
	{
		auth := apiV1.Group("/auth")
		{
			auth.POST("/user-register", h.createUser)
			auth.POST("/user-login", h.loginUser)
		}

		api := apiV1.Group("/", h.userIdentity)
		{
			users := api.Group("/users")
			{
				users.PUT("/", h.updateUser)
				users.DELETE("/:id", h.deleteUser)
				users.GET("/:user_id/products", h.getProductsByUserId)
				users.GET("/:user_id/posts", h.getPostsByUserId)
			}

			products := api.Group("/products")
			{
				products.POST("/", h.createProduct)
				products.GET("/:id", h.getProductById)
				products.GET("/", h.getAllProducts)
				products.PUT("/", h.updateProduct)
				products.DELETE("/:id", h.deleteProduct)
				products.GET("/:id/comments", h.getAllProductsComments)
			}

			posts := api.Group("/posts")
			{
				posts.POST("/", h.createPost)
				posts.GET("/:id", h.getPostById)
				posts.GET("/", h.getAllPosts)
				posts.PUT("/", h.updatePost)
				posts.DELETE("/:id", h.deletePost)
				posts.GET("/:id/comments", h.getAllPostsComments)
			}

			comments := api.Group("/comments")
			{
				comments.POST("/", h.createComment)
				comments.GET("/:id", h.getCommentsByParentId)
				//delete
			}

		}
	}

	return router
}
