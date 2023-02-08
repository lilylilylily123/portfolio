package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func redirection(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/main/", 302)
}
func cookies(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "test",
		Value: "1",
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/projects", 302)
}

func alreadySeen(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("test")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Redirect(w, r, "/projects", 302)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		http.Redirect(w, r, "/seenprojects", 302)
	}
	fmt.Println(cookie)
}

const FileDir = "public"

func main() {
	http.Handle("/main/", http.StripPrefix("/main", http.FileServer(http.Dir(FileDir+"/main"))))
	//http.Handle("/projects/", http.StripPrefix("/projects", http.FileServer(http.Dir(FileDir+"/projects"))))
	http.HandleFunc("/projects/", alreadySeen)
	http.HandleFunc("/project", cookies)
	http.Handle("/seenprojects/", http.StripPrefix("/seenprojects", http.FileServer(http.Dir(FileDir+"/seen-projects"))))
	http.Handle("/about/", http.StripPrefix("/about", http.FileServer(http.Dir(FileDir+"/about"))))
	http.HandleFunc("/", redirection)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
