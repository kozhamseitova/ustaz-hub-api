package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/user-register", h.createUser)
		auth.POST("/user-login", h.loginUser)
	}

	apiV1 := router.Group("/api/v1")
	{
		users := apiV1.Group("/users")
		{
			users.PUT("/", h.updateUser)
			users.DELETE("/", h.deleteUser)
			users.GET("/:user_id/products", h.getProductsByUserId)
			users.GET("/:user_id/posts", h.getPostsByUserId)
		}

		products := apiV1.Group("/products")
		{
			products.POST("/", h.createProduct)
			products.GET("/:id", h.getProductById)
			products.GET("/", h.getAllProducts)
			products.PUT("/", h.updateProduct)
			products.DELETE("/", h.deleteProduct)
			products.GET("/:id/comments", h.getAllProductsComments)
		}

		posts := apiV1.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/:id", h.getPostById)
			posts.GET("/", h.getAllPosts)
			posts.PUT("/", h.updatePost)
			posts.DELETE("/", h.deletePost)
			posts.GET("/:id/comments", h.getAllPostsComments)
		}

		comments := apiV1.Group("/comments")
		{
			comments.POST("/", h.createComment)
			comments.GET("/:id", h.getCommentsByParentId)
		}

	}

	return router
}
