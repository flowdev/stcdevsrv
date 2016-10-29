package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port `number` of the web server")
}

func main() {
	// first parse command line flags and mandatory argument:
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("ERROR: Exactly one directory to serve content from is needed as last argument!")
	}
	dir := flag.Args()[0]

	// now validate port and directory:
	if port < 1 || port > 65535 {
		log.Fatalln("ERROR: The web server port number must be in the range 1 to 65535!")
	}
	finf, err := os.Stat(dir)
	if err != nil {
		log.Fatalln("ERROR: The given content directory", dir, " doesn't exist!")
	}
	if !finf.IsDir() {
		log.Fatalln("ERROR: The given content directory", dir, " isn't a directory!")
	}

	// give some information to the user:
	log.Println("Serving content from:", dir)
	log.Println("Port:", port)
	log.Println("Please stop the web server by pressing: ^C")

	// finally start the static webserver:
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(dir))))
}
