package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type UrlData struct {
	CreateTime  time.Time
	Description string
	Link        string
	Tags        string
}

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "file:./_files/urls.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	qs := []string{
		`CREATE TABLE IF NOT EXISTS Urls (
			id INTEGER PRIMARY KEY,
			create_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			url TEXT UNIQUE DEFAULT 'no_url',
			description TEXT DEFAULT 'no_description',
			tags TEXT DEFAULT 'no_tags'
		)`,

		`CREATE VIRTUAL TABLE IF NOT EXISTS UrlsFTS USING fts5(url, tags)`,

		`CREATE TRIGGER IF NOT EXISTS Urls_After_Insert AFTER INSERT ON Urls
		 BEGIN
			 INSERT INTO UrlsFTS(rowid, url, tags) VALUES (NEW.rowid, NEW.url, NEW.tags);
		 END;`,

		`CREATE TRIGGER IF NOT EXISTS Urls_After_Update AFTER UPDATE ON Urls
		 BEGIN
			 UPDATE UrlsFTS SET url = NEW.url, tags = NEW.tags WHERE rowid = OLD.rowid;
		 END;`,

		`CREATE TRIGGER IF NOT EXISTS Urls_After_Delete AFTER DELETE ON Urls
		 BEGIN
			 DELETE FROM UrlsFTS WHERE rowid = OLD.rowid;
		 END;`,
	}

	for _, q := range qs {
		if err := sendQuery(q); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("DB initiate")
}

func sendQuery(query string, args ...interface{}) error {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args...)

	return err
}

func SaveLink(u UrlData) error {
	q := `INSERT INTO Urls (url, description, tags) VALUES (?, ?, ?)
                      ON CONFLICT(url) DO UPDATE SET 
                      description=excluded.description,
                      tags=excluded.tags`

	return sendQuery(q, u.Link, u.Description, u.Tags)
}

func GetUrlList() []UrlData {
	var list []UrlData

	rows, err := DB.Query(`SELECT create_at, url, description, tags FROM Urls`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var url UrlData
		if err := rows.Scan(&url.CreateTime, &url.Link, &url.Description, &url.Tags); err != nil {
			log.Println("Error scanning UrlData:", err)

			continue
		}

		list = append(list, url)
	}

	return list
}

func DeleteUrl(link string) error {
	q := `DELETE FROM Urls WHERE url = ?`
	return sendQuery(q, link)
}

func SearchUrls(matchField, query string) ([]UrlData, error) {
	var results []UrlData

	// Формирование запроса с использованием выбранного поля для MATCH
	searchQuery := fmt.Sprintf(`SELECT u.url, u.description, u.tags 
                                FROM UrlsFTS AS fts
                                JOIN Urls AS u ON fts.rowid = u.rowid
                                WHERE fts.%s MATCH ?`, matchField)

	// Формирование запроса для полнотекстового поиска
	formattedQuery := query + "*"

	rows, err := DB.Query(searchQuery, formattedQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url UrlData
		if err := rows.Scan(&url.Link, &url.Description, &url.Tags); err != nil {
			return nil, err
		}
		results = append(results, url)
	}

	return results, nil
}
