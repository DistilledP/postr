package handler

import "net/http"

type RootHandler struct {
}

func (*RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
