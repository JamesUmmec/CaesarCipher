package local_server

import (
	"fmt"
	"net/http"
	"testing"
)

// It will exit the program if no available port,
// so unnecessary to deal with the error here.
func TestAvailablePort(t *testing.T) {
	port := AvailableAddr()
	fmt.Println(port)
}

// It will call default browser in your computer to visit github.
// now only support Windows, Linux and Darwin(MacOS)
func TestOpenInBrowser(t *testing.T) {
	OpenInBrowser("https://github.com")
}

// Attention: This test function will not exit automatically,
// because there is a loop to hold the server.
// Don't forget to use Ctrl+C in terminal to exit the process
// to prevent dead loop and endless test.
func TestWebUI(t *testing.T) {
	// get available address
	addr := AvailableAddr()

	// don't use https here...
	OpenInBrowser(fmt.Sprintf("http://%qhello", addr))

	// create simple webserver for test
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic("serve error")
	}
}

// helloHandler handle a server for the test,
// which when the page opened, you'll see "Hello, it works!"
// with the address of "/hello" on its bottom.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(
		w, "<h1>Hello, it works!</h1> you are visiting: %q\n", r.URL.Path)
	if err != nil {
		panic("http connect error")
	}
}
