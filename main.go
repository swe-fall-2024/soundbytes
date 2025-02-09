package main

import (
	//"encoding/json"
	"fmt"
	"log"

	//"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	//"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


func main() {
	host := "127.0.0.1:4201" // may be 4201
    fmt.Println("Starting server on " + host)  // Debugging line

	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		fmt.Print("Failed to listen to " + host)
		log.Fatalf("Failed to listen on %s: %v", host, err)
	} else {
		fmt.Print("Listening to " + host)
	}
}

// httpHandler creates the backend HTTP router for queries, types,
// and serving the Angular frontend.
func httpHandler() http.Handler {
    fmt.Print("inside of httpHandler in Go")
    
    router := mux.NewRouter()

    // ✅ Define /api/message endpoint
    router.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "Hello from Go!"}`))
    }).Methods("GET")

    return handlers.LoggingHandler(os.Stdout,
        handlers.CORS(
            handlers.AllowCredentials(),
            handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
                "DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
                "Cache-Control", "Content-Range", "Range"}),
            handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
            handlers.AllowedOrigins([]string{"http://localhost:4200"}),
            handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
                "X-Requested-With", "If-Modified-Since", "Cache-Control",
                "Content-Type", "Content-Range", "Range", "Content-Disposition"}),
            handlers.MaxAge(86400),
        )(router))
}


func getOrigin() *url.URL {
    origin, err := url.Parse("http://localhost:4200")
    if err != nil {
        log.Fatalf("Failed to parse origin URL: %v", err)
    }
    return origin
}

var origin = getOrigin()

var director = func(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.Header.Add("X-Origin-Host", origin.Host)
	req.URL.Scheme = "http"
	req.URL.Host = origin.Host
}

// AngularHandler loads angular assets
var AngularHandler = &httputil.ReverseProxy{Director: director}