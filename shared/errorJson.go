package shared

import (
	"fmt"
	"net/http"
)

func SendError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(
		w,
		fmt.Sprintf("{\"error\":\"%s\"}", msg),
		code,
	)
}
