package handler

import (
	"net/http"
	"os"
)

func FileReadSt_Handler(w http.ResponseWriter, r *http.Request) {

	FB, FErr := os.ReadFile("files/file1.txt")

	if FErr != nil {
		w.Write([]byte("Error in reading file" + FErr.Error()))
	} else {
		w.Write(FB)
	}
}
