package main

import (
	"database/sql"
	"fmt"
	"time"

	review "github.com/sky0621/fs-mng-review"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	e := loadEnv()

	/*
	 * setup db client
	 */
	var db *sql.DB
	{
		var err error
		// MEMO: ひとまずローカルのコンテナ相手の接続前提なので、べたに書いておく。
		dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%s", e.DBName, e.DBUser, e.DBPassword, e.DBSSLMode, e.DBPort)
		db, err = sql.Open(e.DBDriverName, dataSourceName)
		if err != nil {
			panic(err)
		}
		defer func() {
			if db != nil {
				if err := db.Close(); err != nil {
					panic(err)
				}
			}
		}()

		boil.DebugMode = true

		var loc *time.Location
		loc, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(loc)

		if err := db.Ping(); err != nil {
			panic(err)
		}
	}

	/*
	 * setup gRPC server
	 */
	review.StartServer(e.ServerPort, db)
}
