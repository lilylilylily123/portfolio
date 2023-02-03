package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./public/main/"))
	fs2 := http.FileServer(http.Dir("./public/main/projects/"))
	http.Handle("/main/", http.StripPrefix("/main/", fs))
	http.Handle("/projects/", http.StripPrefix("/projects/", fs2))
	http.ListenAndServe(":6969", nil)
}
