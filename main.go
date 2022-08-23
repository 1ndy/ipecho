package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ensure port is a valid number
func checkPortInRange(p int) bool {
	return 1 <= p && p <= 65535
}

// report the Host used to access the webpage
func handler(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	log.Printf("Got request from %s\n", ip)
	fmt.Fprint(w, ip)
}

// get the IP address as it appears on the internet
func externalHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://checkip.amazonaws.com/")
	if err != nil {
		log.Println("Failed to fetch external IP")
		fmt.Fprint(w, "error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to fetch external IP")
		fmt.Fprint(w, "error")
	}

	sb := strings.TrimSpace(string(body))
	ip := strings.Split(r.RemoteAddr, ":")[0]
	log.Printf("Got external IP %s for client %s\n", sb, ip)
	fmt.Fprint(w, sb)
}

func main() {

	portPtr := flag.Int("port", 5353, "Port on which to run server")

	flag.Parse()

	if !checkPortInRange(*portPtr) {
		log.Fatalf("Port %d is not in range", *portPtr)
	}
	log.Printf("Starting ip echo on port %d\n", *portPtr)

	http.HandleFunc("/", handler)
	http.HandleFunc("/external", externalHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil))
}
