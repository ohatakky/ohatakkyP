package functions

import (
	"net/http"

	"github.com/ohatakky/ohatakkyp/cmd/blog"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	blog.Exec()
}
