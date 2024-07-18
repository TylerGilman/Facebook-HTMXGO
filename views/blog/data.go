package blog

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Article struct {
	ID       int
	Title    string
	Author   string
	Date     time.Time
	Summary  string
	ImageUrl string
	Category string
	Content  string
}

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		return err
	}

	// Create the articles table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS articles (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            author TEXT,
            date DATETIME,
            summary TEXT,
            image_url TEXT,
            category TEXT,
            content TEXT
        )
    `)
	return err
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func SaveArticle(article Article) (int64, error) {
	result, err := db.Exec(`
        INSERT INTO articles (title, author, date, summary, image_url, category, content)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `, article.Title, article.Author, article.Date, article.Summary, article.ImageUrl, article.Category, article.Content)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetAllArticles() ([]Article, error) {
	rows, err := db.Query("SELECT * FROM articles ORDER BY date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		var dateStr string
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &dateStr, &a.Summary, &a.ImageUrl, &a.Category, &a.Content)
		if err != nil {
			return nil, err
		}
		a.Date, _ = time.Parse("2006-01-02 15:04:05", dateStr)
		articles = append(articles, a)
	}
	return articles, nil
}

func SearchArticles(query string, category string) ([]Article, error) {
	sqlQuery := `
        SELECT * FROM articles 
        WHERE (title LIKE ? OR summary LIKE ? OR content LIKE ?) 
        AND (? = '' OR category = ?)
        ORDER BY date DESC
    `
	rows, err := db.Query(sqlQuery, "%"+query+"%", "%"+query+"%", "%"+query+"%", category, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		var dateStr string
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &dateStr, &a.Summary, &a.ImageUrl, &a.Category, &a.Content)
		if err != nil {
			return nil, err
		}
		a.Date, _ = time.Parse("2006-01-02 15:04:05", dateStr)
		articles = append(articles, a)
	}
	return articles, nil
}

func GetRandomArticles(n int) ([]Article, error) {
	rows, err := db.Query("SELECT * FROM articles ORDER BY RANDOM() LIMIT ?", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		var dateStr string
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &dateStr, &a.Summary, &a.ImageUrl, &a.Category, &a.Content)
		if err != nil {
			return nil, err
		}
		a.Date, _ = time.Parse("2006-01-02 15:04:05", dateStr)
		articles = append(articles, a)
	}
	return articles, nil
}
