package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main()  {
	port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		port = "8080"
	}

	http.HandleFunc("/api/generatePdf", function.GeneratePdf)
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
