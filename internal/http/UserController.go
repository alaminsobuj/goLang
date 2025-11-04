package http

import (
	"database/sql" 
	"net/http"
	 
)

type UserController  struct {
	DB *sql.DB
}
 

// GET /users
func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello user! ✅"))
		// SQL query
}
