package sqlite

import (
	"api/internal/config"
	"api/internal/logger"
	"database/sql"
	"fmt"
	"log/slog"
	"math/rand"

	_ "modernc.org/sqlite"
)

type Database struct {
	Connect *sql.DB
	Logs    *logger.Logs
}

func New(cfg *config.Config, loggger *logger.Logs) (*Database, error) { // Loading new database
	conn, err := sql.Open("sqlite", "internal/sqlite/database.db")
	if err != nil {
		return nil, err
	}

	err = conn.Ping() // Testing connection to database
	if err != nil {
		return nil, err
	}

	conn.Exec(` 
	CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		link TEXT UNIQUE NOT NULL,
		alias TEXT UNIQUE NOT NULL
	)`) // Create table

	return &Database{Connect: conn, Logs: loggger}, nil
}

func (d *Database) GetLink(alias string) (string, error) { //Search link on database
	var link string
	err := d.Connect.QueryRow("SELECT link FROM links WHERE alias=?", alias).Scan(&link)

	//Handler error
	switch {
	case err == sql.ErrNoRows:
		return "", fmt.Errorf("G3")
	case err != nil:
		d.Logs.Sqlite.Warn("GetLink",
			slog.String("Status", "Error"),
			slog.String("Error", err.Error()),
		)
		return "", fmt.Errorf("G1")
	default:
		return link, nil
	}
}

func (d *Database) SetLink(link string) (string, error) { // Write link on database
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Charset for random alias
	b := make([]byte, 6)                                                        // Variable for random alias
	var alias string

	// View database on copy link
	err := d.Connect.QueryRow("SELECT alias FROM links WHERE link=?", link).Scan(&alias)
	if err == sql.ErrNoRows { // If not copy
		for i := 0; i < 10; i++ {
			for i := range b {
				b[i] = charset[rand.Intn(len(charset))] // Generate random alias
			}

			// View database on copy alias
			err = d.Connect.QueryRow("SELECT id FROM links WHERE alias=?", string(b)).Scan(&alias)
			if err == sql.ErrNoRows { // If not copy
				d.Connect.Exec("INSERT INTO links (link, alias) VALUES (?, ?)", link, string(b))
				d.Logs.Sqlite.Info("New alias",
					slog.String("Link", link),
					slog.String("Alias", string(b)),
				)
				alias = string(b)
				break
			} else if err != nil { // If copy alias is have - continue cycle
				continue
			}
		}
	} else if err != nil { // Handler error
		d.Logs.Sqlite.Warn("New SetLink",
			slog.String("Status", "Error"),
			slog.String("Error", err.Error()),
		)

		return "", fmt.Errorf("S1")
	}

	return alias, nil
}
