package book

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"book-services/entities/domain"
	"book-services/entities/web"
	"book-services/helpers"
	books "book-services/services/book"
)

type bookcontroller struct {
	bookcontroll books.BookService
}

func NewBookController(book books.BookService) BookController {
	return &bookcontroller{
		bookcontroll: book,
	}
}