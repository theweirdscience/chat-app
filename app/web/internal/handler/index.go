package handler

import "net/http"

func IndexHander() http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		//do something
	})
}