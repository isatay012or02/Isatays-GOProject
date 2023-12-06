package routes

import (
	"github.com/gin-gonic/gin"
	"book-services/config"
	"book-services/entities/domain"
	books "book-services/handlers/book"
	"book-services/handlers/category"
)

func Run(book books.BookController, category category.CategoryController) {
	db := config.GetConnection()
	db.AutoMigrate(&domain.Book{})

	server := gin.Default()

	bookGroup := server.Group("/api/book")
	{
		bookGroup.GET("/", book.FindAll)
		bookGroup.POST("/", book.CreateBook)
		bookGroup.PUT("/:id", book.UpdateBook)
		bookGroup.GET("/:id", book.FindID)
		bookGroup.DELETE("/:id", book.DeleteBook)
	}

	conf := config.GetConfig()

	server.Run(conf.RunnApp.Host + ":" + conf.RunnApp.Port)
}
