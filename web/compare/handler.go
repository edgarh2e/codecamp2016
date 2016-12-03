package compare

import (
	"encoding/json"
	"github.com/pressly/chi"
	"log"
	"net/http"
)

func imageHandler(w http.ResponseWriter, r *http.Request) {
	file := chi.URLParam(r, "file")

	buf, err := readImage(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buf)
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	users := make([]string, 0, 2)
	for _, u := range r.Form["user"] {
		users = append(users, u)
	}
	log.Printf("user: %v", users)

	res, err := compare(users...)
	if err != nil {
		log.Printf("xxx: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("res: %v", res)

	out, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}
