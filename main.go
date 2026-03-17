package main
 
import (
    "fmt"
    "log"
    "net/http"
    "os"
)
 
func newHandler() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = fmt.Fprintln(w, "github-actions-demo is running")
    })
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = fmt.Fprintln(w, "Hello, this go application is healthy!")
    })
    return mux
}
 
func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
 
    addr := ":" + port
    log.Printf("listening on %s", addr)
    if err := http.ListenAndServe(addr, newHandler()); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
