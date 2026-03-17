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
 
 
