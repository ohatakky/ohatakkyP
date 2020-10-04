package functions

import (
	"log"
	"net/http"

	"github.com/ohatakky/ohatakkyp/functions/blog"
	"github.com/ohatakky/ohatakkyp/functions/trending"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	err := blog.Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func TrendingHandler(w http.ResponseWriter, r *http.Request) {
	err := trending.Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
