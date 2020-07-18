package renderers

import (
	"fmt"
	"net/http"
)

func writeErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)                                     // Status Code
	w.Header().Set("Content-Type", "application/json")        // Response Type - JSON
	w.Write([]byte("{\"message\": \"" + err.Error() + "\"}")) // Error Message
}

func renderErr(w http.ResponseWriter, status int, err string) {
	writeErr(w, status, fmt.Errorf("%v", err))
}
