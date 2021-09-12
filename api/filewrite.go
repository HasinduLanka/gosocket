package handler

import (
	"net/http"
	"os"
	"time"
)

func FileWrite_Handler(w http.ResponseWriter, r *http.Request) {

	FWErr := os.WriteFile("files/file1.html", []byte("<html><body>Hello World <p> Time is : "+string(time.Now().String())+" </p> </body></html>"), os.ModePerm)

	if FWErr != nil {
		w.Write([]byte("Error in writing file"))
	}

	FB, FErr := os.ReadFile("files/file1.html")

	if FErr != nil {
		w.Write([]byte("Error in reading file"))
	} else {
		w.Write(FB)
	}
}
