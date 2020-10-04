package functions

import (
	"log"
	"net/http"

	"github.com/ohatakky/ohatakkyp/cmd/blog"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	err := blog.Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
