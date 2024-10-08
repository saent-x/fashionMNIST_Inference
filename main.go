package main

import (
	"fmt"
	"github.com/saent-x/fashionMNIST_Inference/handlers"
	"github.com/saent-x/fashionMNIST_Inference/pages"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":3000", Setup())
	if err != nil {
		log.Fatalln(err)
	}
}

func Setup() http.Handler {
	mux := http.NewServeMux()

	handlers.Static(mux)
	handlers.HandleUpload(mux)
	pages.Home(mux)

	fmt.Println("server started...")

	return mux
}
