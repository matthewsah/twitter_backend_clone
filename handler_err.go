package main

import "net/http"

func handler_err(w http.ResponseWriter, r *http.Request) {
	respond_with_error(w, 400, "Something went wrong")
}
