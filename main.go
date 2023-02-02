package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./public/start/"))
	fs2 := http.FileServer(http.Dir("./public/start/projects/"))
	http.Handle("/start/", http.StripPrefix("/start/", fs))
	http.Handle("/projects/", http.StripPrefix("/projects/", fs2))
	http.ListenAndServe(":6969", nil)
}
