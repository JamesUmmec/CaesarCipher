package local_server

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
)

// AvailableAddr generate an available port to use by calling the system.
// If there's no port available, it will exit the program
// in order to prevent unnecessary computational cost.
func AvailableAddr() string {
	// Randomly generate an available port.
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		cannotOpenPort()
	}

	// Connect to the port to check whether this port is truly available.
	temp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		cannotOpenPort()
	}

	// Close connection to make place for the util connection.
	// if something wrong and cannot close connection,
	// it will then exit the program.
	defer func(temp *net.TCPListener) {
		err := temp.Close()
		if err != nil {
			cannotOpenPort()
		}
	}(temp)

	return temp.Addr().String()
}

func cannotOpenPort() {
	fmt.Println("Error when opening port.")
	os.Exit(404)
}

// OpenInBrowser call system to open url in default browser.
// This url must contain http:// and port when coding local server.
// If it cannot open the url in default browser, which means the ui is unavailable,
// it will directly exit the program, to prevent unnecessary computational cost.
// It only support windows, linux and darwin(macos) now.
func OpenInBrowser(url string) {
	// Generate command to run in terminal of certain os platform.
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	}

	// Execute the commands and if error, exit the program.
	err := cmd.Start()
	if err != nil {
		fmt.Println("Cannot open system default browser.")
		os.Exit(500)
	}
}
