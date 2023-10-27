package main

import "net/http"

func handler_readiness(w http.ResponseWriter, r *http.Request) {
	respond_with_JSON(w, 200, struct{}{})
}
