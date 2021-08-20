package main

import (
	"log"

	"github.com/1k-ct/amble/pkg/database"
)

func main() {
	// bookPersistence := persistence.NewBookPersistence()
	// bookUseCase := usecase.NewBookUseCase(bookPersistence)
	// bookHandler := rest.NewBookHandler(bookUseCase)

	// router := httprouter.New()
	// router.GET("/api/v1/books", bookHandler.Index)
	// router.GET("/api/v1/book", bookHandler.Home)

	file := "/.env"
	config, err := database.NewDB(file)
	if err != nil {
		log.Fatal(err)
	}
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ok")
	defer db.Close()

	// router := router.BookRouter()

	// http.ListenAndServe(":3000", router)
}
