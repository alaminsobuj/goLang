package http

import "net/http"

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ Hello! Server is running."))
	})
    uc := &UserController{}
		// Example additional route
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/users", uc.Index)
	return r
}
