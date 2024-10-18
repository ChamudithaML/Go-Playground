package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/lobby", lobbyPage)

	fmt.Println("Server is starting at :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		http.NotFound(w, r)
	}
}

func lobbyPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/lobby" {
		http.ServeFile(w, r, "static/lobby.html")
	} else {
		http.NotFound(w, r)
	}
}
