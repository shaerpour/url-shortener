package internal

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type DBConfig struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

func InitDB() {
	// Open new connetion to database
	db := new(sql.DB)
	var err error
	var c = DBConfig{
		Host:   os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		User:   os.Getenv("MYSQL_USER"),
		Pass:   os.Getenv("MYSQL_PASS"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	db, err = sql.Open("mysql", c.User+":"+c.Pass+"@tcp("+c.Host+")/"+c.DBName)
	if err != nil {
		log.Fatal(err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(err)
	}
	DB = db // Create db connection pointer
	// Create table if needed
	GetTableQuery := "SHOW TABLES"
	var table string
	_ = DB.QueryRow(GetTableQuery).Scan(&table)
	if table != "" {
		return
	} else {
		CreateTableQuery := "CREATE TABLE IF NOT EXISTS url_list (url varchar(255), randStr varchar(255))"
		_, err := DB.Exec(CreateTableQuery)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func AddDomain(url, randStr string) error {
	query := "INSERT INTO url_list (url, randStr) VALUES (?, ?)"
	_, err := DB.Exec(query, url, randStr)
	if err != nil {
		return err
	}
	return nil
}

func GetDomainByStr(randStr string) (string, error) {
	var url string
	query := "SELECT url FROM url_list WHERE randStr = ?"
	_ = DB.QueryRow(query, randStr).Scan(&url)
	return url, nil
}

func GetStrByDomain(url string) (string, error) {
	var randStr string
	query := "SELECT randStr FROM url_list WHERE url = ?"
	_ = DB.QueryRow(query, url).Scan(&randStr)
	return randStr, nil
}

func GetAllDomains() []string {
	var ListOfURLs []string
	query := "SELECT randStr FROM url_list"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var UrlString string
		rows.Scan(&UrlString)
		ListOfURLs = append(ListOfURLs, UrlString)
	}
	return ListOfURLs
}
