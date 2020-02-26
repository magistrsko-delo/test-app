package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main()  {
	fmt.Println("main")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	fmt.Println(reflect.TypeOf(db))
	if err != nil {
		panic(err.Error())
	}
	defer  db.Close()

	results, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	testResult := []Test{}

	for results.Next() {
		var testStruct Test

		err = results.Scan(&testStruct.ID, &testStruct.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(testStruct)
		testResult = append(testResult, testStruct)
	}
	fmt.Println("-----------", testResult)


}

/**
stmt, err := db.Prepare("select id, name from users where id = ?")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()
rows, err := stmt.Query(1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	// ...
}
if err = rows.Err(); err != nil {
	log.Fatal(err)
}


stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
if err != nil {
	log.Fatal(err)
}
res, err := stmt.Exec("Dolly")
if err != nil {
	log.Fatal(err)
}
lastId, err := res.LastInsertId()
if err != nil {
	log.Fatal(err)
}
rowCnt, err := res.RowsAffected()
if err != nil {
	log.Fatal(err)
}
log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)


tx, err := db.Begin()
if err != nil {
	log.Fatal(err)
}
defer tx.Rollback()
stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close() // danger!
for i := 0; i < 10; i++ {
	_, err = stmt.Exec(i)
	if err != nil {
		log.Fatal(err)
	}
}
err = tx.Commit()
if err != nil {
	log.Fatal(err)
}
// stmt.Close() runs here!
*/

