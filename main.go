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
	cookie, err := r.Cookie("hasVisited")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			cookie := &http.Cookie{
				Name:  "hasVisited",
				Value: "true",
				Path:  "/",
			}
			http.SetCookie(w, cookie)
			log.Println("cookie not set, redirecting to cookie setter:", err)
			http.Redirect(w, r, "/gallery/", 302)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		log.Println("cookie already set, skipping animations")
		http.Redirect(w, r, "/work/", 302)
	}
	fmt.Println(cookie)
	// http.Redirect(w, r, "/projects/", 302)
}

const FileDir = "public"

func main() {
	http.Handle("/main/", http.StripPrefix("/main", http.FileServer(http.Dir(FileDir+"/main"))))
	log.Println("listening on /main/")
	// http.Handle("/projects/", http.StripPrefix("/projects", http.FileServer(http.Dir(FileDir+"/projects"))))
	// http.HandleFunc("/projects/", alreadySeen)
	http.Handle("/gallery/", http.StripPrefix("/gallery", http.FileServer(http.Dir(FileDir+"/projects"))))
	log.Println("listening on /gallery/")
	http.HandleFunc("/project/", cookies)
	log.Println("redirecting from /project/")
	http.Handle("/work/", http.StripPrefix("/work", http.FileServer(http.Dir(FileDir+"/seen-projects"))))
	log.Println("listening on /work")
	http.Handle("/about/", http.StripPrefix("/about", http.FileServer(http.Dir(FileDir+"/about"))))
	log.Println("listening on /about/")
	http.HandleFunc("/", redirection)
	log.Println("listening for unknown urls")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
