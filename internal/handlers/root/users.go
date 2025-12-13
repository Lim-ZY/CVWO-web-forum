package root

import (
	"net/http"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Main Page here!"))
}
