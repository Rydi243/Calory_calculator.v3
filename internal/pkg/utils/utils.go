package utils

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func WriteInternalError(w http.ResponseWriter, err error, message string) {
	fmt.Fprintln(w, errors.Wrap(err, message))
	w.WriteHeader(http.StatusInternalServerError)
}
