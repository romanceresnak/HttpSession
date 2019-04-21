package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

const
(
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

//we declared a private cookie store to store session using secure cookies
var store *sessions.CookieStore

//init always run before main function
//this function create a new cookie and save it to store variable
func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func home(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session-name")
	var authenticated = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You are unauthorized to view the page",
				http.StatusForbidden)
			return
		}
		fmt.Fprintln(w, "Home Page")
	} else
	{
		http.Error(w, "You are unauthorized to view the page",
			http.StatusForbidden)
		return
	}
}

//for logging user to session
func login(w http.ResponseWriter, r *http.Request){
	//Get value from cookie store with same name
	session, _ := store.Get(r,"session-name")
	//Set authenticated to true
	session.Values["authenticated"]=true
	//Save request and responseWriter
	session.Save(r,w)
	//Print the result to console
	fmt.Println(w,"You have succesfully login")
}

func logout(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r,"sesion_name")
	session.Values["authenticated"]=false
	session.Save(r,w)
	fmt.Println(w, "You are logout")
}

func main() {
http.HandleFunc("/home", home)
http.HandleFunc("/login", login)
http.HandleFunc("/logout", logout)
err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
if err != nil {
	log.Fatal("error starting http server : ", err)
	return
	}
}