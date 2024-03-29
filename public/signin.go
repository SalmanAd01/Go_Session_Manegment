package public

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET")))

func SigninGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/signin.html")
}
func SigninPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	// Place For Hashing Database
	if username == My_Map["username"] && password == My_Map["password"] {
		session, _ := Store.Get(r, "auth-session")
		session.Values["username"] = username
		session.Save(r, w)
		fmt.Println("username:", username)
		http.Redirect(w, r, "/dashboard", 302)
	} else {
		fmt.Println("username:", username, "password:", password)
		http.Redirect(w, r, "/signup", 302)
	}
}
