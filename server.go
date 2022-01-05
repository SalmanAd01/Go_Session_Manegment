package main

import (
	auth "User_Session_Manegment/middleware"
	"User_Session_Manegment/public"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", auth.IsNotAuth(public.Home)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(public.SigninGet)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(public.SigninPost)).Methods("POST")
	router.HandleFunc("/signup", auth.IsNotAuth(public.SignupGet)).Methods("GET")
	router.HandleFunc("/signup", auth.IsNotAuth(public.SignupPost)).Methods("POST")
	router.HandleFunc("/dashboard", auth.IsAuth(public.Dashboard)).Methods("GET")
	router.HandleFunc("/logout", auth.IsAuth(public.Logout)).Methods("GET")
	http.ListenAndServe(":8080", router)
}
