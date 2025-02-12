package typerace

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Quote struct {
	ID          int    `json:"id"`
	TyperacerID int    `json:"typeracer_id"`
	Text        string `json:"text"`
}

func newQuote(id int, typeracerID int, text string) *Quote {
	return &Quote{
		ID:          id,
		TyperacerID: typeracerID,
		Text:        text,
	}
}

//go:embed typeracerPtBrQuotes.db
var embeddedDB []byte

func connect() (*sql.DB, error) {
	tmpFile, err := os.CreateTemp("", "db.db")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(embeddedDB); err != nil {
		panic(err)
	}
	tmpFile.Close()
	db, err := sql.Open("sqlite3", tmpFile.Name())
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		slog.Error("Error pinging sqlite database", slog.String("error", err.Error()))
	}

	return db, err
}

func GetQuote() *Quote {
	db, err := connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var idStr []byte
	var typeracerIdStr []byte
	var text []byte

	row := db.QueryRow(fmt.Sprintf("SELECT * FROM quotes where id=%v", rand.Intn(1098)))

	err = row.Scan(&idStr, &typeracerIdStr, &text)

	if err != nil {
		panic(err)
	}

	typeracerId, _ := strconv.Atoi(string(typeracerIdStr))
	id, _ := strconv.Atoi(string(idStr))

	quote := newQuote(id, typeracerId, string(text))

	return quote

}
