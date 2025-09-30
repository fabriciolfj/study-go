package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type comment2 struct {
	text       string
	dateString string
}

var comments []comment2

func getComments(w http.ResponseWriter, r *http.Request) {
	commentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if commentID == 0 || len(comments) < commentID {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("%s (%s)\n", comments[i].text, comments[i].dateString)
	}
	fmt.Fprintln(w, fmt.Sprintf("Comments %d: \n%s", commentID, commentBody))
}

func postComments(w http.ResponseWriter, r *http.Request) {
	commentText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	comments = append(comments, comment2{text: string(commentText), dateString: time.Now().String()})
	w.WriteHeader(http.StatusOK)
}

func Execute() {

	http.HandleFunc("GET /comments", getComments)
	http.HandleFunc("GET /comments/{id}", getComments)
	http.HandleFunc("POST /comments", postComments)

	if err := http.ListenAndServe(":8004", nil); err != nil {
		panic(err)
	}
}
