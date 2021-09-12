package handler

import (
	"net/http"
	"os"
)

func FileRead_Handler(w http.ResponseWriter, r *http.Request) {

	FB, FErr := os.ReadFile("file1.html")

	if FErr != nil {
		w.Write([]byte("Error in reading file"))
	} else {
		w.Write(FB)
	}
}
