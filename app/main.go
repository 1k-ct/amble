package main

import (
	"log"

	"github.com/1k-ct/amble/app/handler/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// config, err := database.NewLocalDB("user", "password", "sample")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db, err := config.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := db.Where("").Error; err != nil {
	// 	log.Fatal(err)
	// }

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	r := router.InitRouters()
	if err := r.Run(); err != nil {
		return
	}

	// config, _ := database.NewLocalDB("user", "password", "sample")
	// config.NewMakeDB(
	// 	&model.Tweet{}, &model.Reply{}, &model.LikedUser{},
	// 	&model.RetweetedUser{},
	// )

	// r := gin.Default()
	// r.Static("/assets", "./assets")
	// r.SetHTMLTemplate(html)

	// r.GET("/", func(c *gin.Context) {
	// 	if pusher := c.Writer.Pusher(); pusher != nil {
	// 		// use pusher.Push() to do server push
	// 		if err := pusher.Push("/assets/app.js", nil); err != nil {
	// 			log.Printf("Failed to push: %v", err)
	// 		}
	// 	}
	// 	c.HTML(200, "https", gin.H{
	// 		"status": "success",
	// 	})
	// })

	// // Listen and Server in https://127.0.0.1:8080
	// r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
