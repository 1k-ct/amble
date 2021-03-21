package main

import (
	"net/http"

	"github.com/1k-ct/twitter-dem/app/handler/router"
)

func main() {
	// bookPersistence := persistence.NewBookPersistence()
	// bookUseCase := usecase.NewBookUseCase(bookPersistence)
	// bookHandler := rest.NewBookHandler(bookUseCase)

	// router := httprouter.New()
	// router.GET("/api/v1/books", bookHandler.Index)
	// router.GET("/api/v1/book", bookHandler.Home)

	router := router.BookRouter()

	http.ListenAndServe(":3000", router)
}
