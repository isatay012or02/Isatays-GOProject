package book

import "github.com/gin-gonic/gin"

type BookController interface {
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindID(ctx *gin.Context)
}