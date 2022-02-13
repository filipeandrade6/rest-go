package middleware

// import (
// 	"context"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-chi/chi/v5"
// )

// type (
// 	RowsPerPageKey string
// 	PageNumberKey  string
// )

// const (
// 	RowsPerPage RowsPerPageKey
// 	PageNumber  PageNumberKey
// )

// func Paginate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var rowsPerPage, pageNumber int

// 		// Limite
// 		l := chi.URLParam(r, "limit")
// 		rowsPerPage, err := strconv.Atoi(l)
// 		if err != nil {
// 			rowsPerPage = 10
// 		}
// 		ctx := context.WithValue(r.Context(), RowsPerPage, rowsPerPage)

// 		// Pagina
// 		p := chi.URLParam(r, "page")
// 		pageNumber, err = strconv.Atoi(p)
// 		if err != nil {
// 			pageNumber = 1
// 		}
// 		ctx = context.WithValue(ctx, PageNumber, pageNumber)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
