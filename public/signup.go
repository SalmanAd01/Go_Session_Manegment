package public

import (
	"fmt"
	"net/http"
)

var My_Map = map[string]string{}

func SignupGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/signup.html")
}
func SignupPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	// Place For Authentication,Hashing Database
	My_Map["username"] = username
	My_Map["password"] = password
	fmt.Println("username:", username, "password:", password)
	http.Redirect(w, r, "/signin", http.StatusFound)
}
