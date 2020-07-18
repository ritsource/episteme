package renderers

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderErr(w, 404, fmt.Sprintf("Page Not Found"))
}
