package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, web!")
	})
	router.HandleFunc("/name/{PARAM}", handleGetParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleGetParam(w http.ResponseWriter, r *http.Request) {
	p, err := mux.Vars(r)["PARAM"]
	if err {
		fmt.Println(w, err)
	}
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w,"Hello, %v", p)
	w.Write([]byte(fmt.Sprintf("Hello, %v!", p)))

}
func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

}

func handleData(w http.ResponseWriter, r *http.Request) {
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, err)
	}
	s := "I got message:\n" + string(p)
	w.Write([]byte(s))

}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	x, _ := strconv.Atoi(r.Header.Get("a"))
	y, _ := strconv.Atoi(r.Header.Get("b"))

	w.Header().Add("a+b", strconv.Itoa(x+y))

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)

}
