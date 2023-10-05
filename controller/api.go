package controller

import (
	"fmt"
	"net/http"
)

func Exec(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Success")
}
