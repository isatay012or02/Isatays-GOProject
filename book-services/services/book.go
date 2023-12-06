package books

import (
	"errors"

	"book-services/entities/domain"
	"book-services/entities/web"
	books "book-services/repository"
)

type bookservice struct {
	book books.BookRepository
}

func NewBookRepository(bookrepo books.BookRepository) BookService {
	return &bookservice{
		book: bookrepo,
	}
}

func (bk *bookservice) CreateBook(book domain.Book) (web.BookResponse, error) {
	response, err := bk.book.Create(book)
	if err != nil {
		return web.BookResponse{}, errors.New("cant upload book")
	}

	resp := web.BookResponse{
		Id:   response.ID,
		Name: response.Name,

		Price: response.Price,
	}

	return resp, nil
}

func (bk *bookservice) UpdateBook(id uint, book domain.Book) (web.BookResponse, error) {
	response, err := bk.book.Update(id, book)
	if err != nil {
		return web.BookResponse{}, errors.New("cant update book")
	}

	return web.BookResponse{
		Id:   response.ID,
		Name: response.Name,
		Price: response.Price,
	}, nil
}


func (bk *bookservice) FindId(id uint) (web.BookResponse, error) {
	response, err := bk.book.FindId(id)
	if err != nil {
		return web.BookResponse{}, errors.New("cant find book")
	}

	return web.BookResponse{
		Id:   response.ID,
		Name: response.Name,

		Price: response.Price,
	}, nil
}

func (bk *bookservice) DeleteBook(id uint) (web.BookDeleteRespons, error) {
	err := bk.book.Delete(id)
	if err != nil {
		return web.BookDeleteRespons{}, errors.New("cant delete book")
	}

	return web.BookDeleteRespons{Id: id}, nil
}

func (bk *bookservice) FindAll() ([]domain.Book, error) {
	book, err := bk.book.FindAll()
	if err != nil {
		return book, errors.New("cant find all book")
	}
	return book, nil
}

