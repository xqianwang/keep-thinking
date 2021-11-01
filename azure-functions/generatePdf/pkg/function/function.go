package function

import (
	"fmt"
	"net/http"
)

func GeneratePdf(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Generate Pdf function triggered")
	
}
