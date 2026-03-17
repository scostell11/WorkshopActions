main.go
 
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
 
main_test.go
 
package main
 
import (
    "net/http"
    "net/http/httptest"
    "testing"
)
 
func TestHealth(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rr := httptest.NewRecorder()
 
    newHandler().ServeHTTP(rr, req)
 
    if rr.Code != http.StatusOK {
        t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
    }
}
 
 
