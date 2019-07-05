package golib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/x-mod/errors"
)

func SqliteCrud(dir string) error {
	dbpath := filepath.Join(dir, "golib.db")
	os.Remove(dbpath)

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Println(errors.Annotatef(err, "db open: %s", dbpath))
		return err
	}
	defer db.Close()

	sqlStmt := `
	create table IF NOT EXISTS foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Println(errors.Annotatef(err, "db exec: %s", sqlStmt))
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(errors.Annotate(err, "db Begin: "))
		return err
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Println(errors.Annotate(err, "db Prepare: "))
		return err
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Println(errors.Annotate(err, "db exec: "))
			return err
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Println(errors.Annotate(err, "db query: "))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(errors.Annotate(err, "db row scan: "))
			return err
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Println(errors.Annotate(err, "db row err: "))
		return err
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Println(errors.Annotate(err, "db Prepare: "))
		return err
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Println(errors.Annotate(err, "db QueryRow: "))
		return err
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Println(errors.Annotate(err, "db Exec delete: "))
		return err
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Println(errors.Annotate(err, "db Exec insert: "))
		return err
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Println(errors.Annotate(err, "db Exec select: "))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(errors.Annotate(err, "db scan name: "))
			return err
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Println(errors.Annotate(err, "db row err: "))
		return err
	}
	return nil
}
