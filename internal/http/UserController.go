package http

import (
	"strings"
	"encoding/json"
	"net/http"

    "github.com/alaminsobuj/goLang/internal/models"
)

type UserController  struct {}
 

// GET /users
func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello user! ✅"))
		// SQL query
	users, err := models.GetAllUsers()
   if err != nil {
		// যদি DB query fail হয়, user কে error দেখান
		http.Error(w, "Error fetching users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users
func (uc *UserController) user_store(w http.ResponseWriter, r *http.Request) {

     w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(10 << 20) // 10MB memory limit 
	// ✅ Required for form-data / x-www-form-urlencoded
	if err != nil {
		http.Error(w, "Form parse error: "+err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
    // ✅ Validation Map
	validationErrors := make(map[string]string)
	// If empty, let us debug

	if name == "" {
		validationErrors["name"] = "Name is required"
	} 
	if email == "" {
		validationErrors["email"] = "Email is required"
	} else if !strings.Contains(email, "@") {
		validationErrors["email"] = "Email format is invalid"
	}

	// ✅ If validation failed
	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Validation failed",
			"errors":  validationErrors,
		})
		return
	}
	// if name == "" || email == "" {
	// 	http.Error(w, "No data found, check POSTMAN keys", http.StatusBadRequest)
	// 	return
	// }

	// json.NewEncoder(w).Encode(map[string]string{
	// 	"name":  name,
	// 	"email": email,
	// })
	// w.Header().Set("Content-Type", "application/json")
    // r.ParseForm()
	// name := r.FormValue("name")
	// email := r.FormValue("email")
    // models.InsertUser(name,email)
	// // Step 2: Respond with received data (no console)
	// json.NewEncoder(w).Encode(map[string]string{
	// 	"name":  name,
	// 	"email": email,
	// })


	// // // Step 1: Parse form
	// if err := r.ParseForm(); err != nil {
	// 	http.Error(w, "Form parse error: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }

     // Response to client
	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "User inserted successfully",
	// })

	// var data struct {
	// 	Name  string `json:"name"`
	// 	Email string `json:"email"`
	// }

	// json.NewEncoder(w).Encode(map[string]string{
	// 	"name":  data.Name,
	// 	"email": data.Email,
	// })
	// if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // ✅ Manual validation
	// if data.Name == "" {
	// 	http.Error(w, "Name is required", http.StatusBadRequest)
	// 	return
	// }

	// if data.Email == "" {
	// 	http.Error(w, "Email is required", http.StatusBadRequest)
	// 	return
	// }
    // return 
	// Insert via model
	if err := models.InsertUser(name, email); err != nil {
		http.Error(w, "DB Insert Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Success Response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": map[string]string{
			"name":  name,
			"email": email,
		},
		"status":  "success",
		"message": "User created successfully",
	})
	
}

// GET /users
func (uc *UserController) UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse form-data (required for form-data / x-www-form-urlencoded)
	// if err := r.ParseMultipartForm(10 << 20); err != nil {
	// 	http.Error(w, "Form parse error: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// ✅ Get values using FormValue
	name := r.FormValue("name")
	email := r.FormValue("email")
	id := r.FormValue("id")
    models.UpdateUser(name, email, id)
 
	// ✅ Call model method
	// if err := models.UpdateUser(uc.DB, name, email, id); err != nil {
	// 	http.Error(w, "DB Update Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// ✅ Success response
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"status":  true,
	// 	"message": "User updated successfully",
	// 	"id":      id,
	// })
}
