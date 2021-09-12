package handler

import (
	"net/http"
	"os"
	"time"
)

func FileWrite_Handler(w http.ResponseWriter, r *http.Request) {

	os.MkdirAll("files", 0777)

	FWErr := os.WriteFile("files/file1.html", []byte("<html><body>Hello World <p> Time is : "+string(time.Now().String())+" </p> </body></html>"), os.ModePerm)

	if FWErr != nil {
		w.Write([]byte("Error in writing file" + FWErr.Error()))
	}

	FB, FErr := os.ReadFile("files/file1.html")

	if FErr != nil {
		w.Write([]byte("Error in reading file" + FErr.Error()))
	} else {
		w.Write(FB)
	}
}
