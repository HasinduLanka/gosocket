package handler

import (
	"io/fs"
	"net/http"
	"os"
)

func FileReadSt_Handler(w http.ResponseWriter, r *http.Request) {

	FDErr := os.MkdirAll("files", fs.ModeDir)

	if FDErr != nil {
		w.Write([]byte("Error in mkdir " + FDErr.Error()))
	}

	FB, FErr := os.ReadFile("files/file1.txt")

	if FErr != nil {
		w.Write([]byte("Error in reading file" + FErr.Error()))
	} else {
		w.Write(FB)
	}
}
