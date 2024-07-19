package internal

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

func InitDB() {
	GetTableQuery := "SHOW TABLES"
	var table string
	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = db.QueryRow(GetTableQuery).Scan(&table)
	if table != "" {
		return
	} else {
		CreateTableQuery := "CREATE TABLE url_list (url varchar(255), randStr varchar(255))"
		_, err = db.Exec(CreateTableQuery)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func OpenDB() (*sql.DB, error) {
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
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}
	return db, nil
}

func AddDomain(url, randStr string) error {
	query := "INSERT INTO url_list (url, randStr) VALUES (?, ?)"
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(query, url, randStr)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func GetDomainByStr(randStr string) (string, error) {
	var url string
	query := "SELECT url FROM url_list WHERE randStr = ?"
	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow(query, randStr).Scan(&url)
	return url, nil
}

func GetStrByDomain(url string) (string, error) {
	var randStr string
	query := "SELECT url FROM url_list WHERE url = ?"
	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = db.QueryRow(query, url).Scan(&randStr)
	return randStr, nil
}

func GetAllDomains() []string {
	var ListOfURLs []string
	query := "SELECT randStr FROM url_list"
	db, err := OpenDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(query)
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
