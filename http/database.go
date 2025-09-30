package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type serverControl struct {
	db *pgx.Conn
}

func Execute4() {
	var sc serverControl
	{
		var err error
		sc.db, err = pgx.Connect(context.Background(), "postgres://localhost:5432")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n",
				err)
			os.Exit(1)
		}
	}

	http.HandleFunc("GET /database",
		sc.databaseHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

type comment struct {
	UserID  int    `json:"userID"`
	Comment string `json:"comment"`
}

func (sc serverControl) databaseHandler(w http.ResponseWriter, r *http.Request) {
	var comments []comment
	rows, err := sc.db.Query(r.Context(), "select user_id, comment from comments limit $1", 5)
	defer rows.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var c comment
		if err := rows.Scan(&c.UserID, &c.Comment); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		comments = append(comments, c)
	}

	output, err := json.Marshal(comments)
	if err != nil {

	}
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
}
