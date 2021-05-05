package main

import (
	"fmt"
	"net/http"
	"time"

	password "github.com/JamesUmmec/CaesarCipher/app/password"
	server "github.com/JamesUmmec/CaesarCipher/local_server"
)

func main() {
	// Register template directory
	const PagePath = "./template/dist"
	http.Handle("page", http.FileServer(http.Dir(PagePath)))

	// Register routes to functions
	// unfinished yet.
	password.PasswordWorks()

	// Run Server
	addr := server.AvailableAddr()
	server.OpenInBrowser(fmt.Sprintf("http://%s/page/index.html", addr))
	server.ImokWatch(500*time.Millisecond, 5)
	http.ListenAndServe(addr, nil)
}
