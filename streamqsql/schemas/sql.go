package schemas

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	steamqsql "github.com/PetkoPetkov/streamq-backend/orm"
	_ "modernc.org/sqlite"
)

var queries *steamqsql.Queries

var ddl []string

func loadDDLEmbed() {
	i := 0
	ddl = make([]string, 2)

	dir := "./streamqsql/schemas"
	ext := ".sql"
	files, err := os.ReadDir(dir)

	if err != nil {
		panic(err.Error())
	}

	for _, file := range files {
		if !file.IsDir() {
			extention := filepath.Ext(file.Name())

			fmt.Println(extention)

			if extention == ext {
				fmt.Println("hello")
				content, err := os.ReadFile(dir + "/" + file.Name())

				if err != nil {
					panic(err.Error())
				}

				ddl[i] = string(content)
				i++
			}
		}
	}
}

func SetUpDBConnection() (*steamqsql.Queries, error) {
	ctx := context.Background()
	db, err := sql.Open("sqlite", "./db.sqlite3")

	if err != nil {
		defer db.Close()
		return nil, err
	}

	defer db.Close()

	loadDDLEmbed()

	for _, dd := range ddl {
		fmt.Println(dd)
		if _, err := db.ExecContext(ctx, dd); err != nil {
			return nil, err
		}
	}

	queries := steamqsql.New(db)

	return queries, nil
}

func SetQueryCaller(q *steamqsql.Queries) {
	queries = q
}

func GetQueryCaller() *steamqsql.Queries {
	return queries
}
