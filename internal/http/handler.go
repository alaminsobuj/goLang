package http

import (
	"net/http"
	"database/sql" // এখানে sql প্যাকেজটি import করতে হবে
)

type Handler struct {
	DB *sql.DB
}

// Example handler function
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from handler!"))
}
