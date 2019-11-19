package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbfile = "blog.db"
const create_user = `create table t_user (
	user_id int PRIMARY key,
	name varchar(30) not null,
	email varchar(100) ,
	avatar varchar(200),
	passwd varchar(20) not null,
	role int,
	editor varchar(30)
)`

const create_note = `create table t_note(
	note_key varchar(30),
	user_id int,
	title varchar(100),
	summary varchar(200),
	content text,
	source varchar(200),
	editor varchar(400),
	files varchar(400),
	visit int,
	praise int
)`

const create_msg = `create table t_message(
	msg_key    varchar(30),
	note_key    varchar(30),
	user_id int,
	content varchar(200),
	Praise  int
)`

type DBTx struct {
	db *sql.DB
}

var Dbx *DBTx

func getUserID(begin int) func() int {
	i := begin
	return func() int {
		i += 1
		return i
	}
}

var funcUser func() int

func init() {
	fd, err := os.Open(dbfile)
	flag := false
	if err != nil {
		fmt.Printf("%s %s, will create it\n", dbfile, err)
		flag = true
	} else {
		fd.Close()
	}

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Panic("Failed to Open dbfile", err)
	}
	if flag {
		//需要初始化
		_, err = db.Exec(create_user)
		if err != nil {
			log.Panic("Failed to create_user", err)
		}
		_, err = db.Exec(create_note)
		if err != nil {
			log.Panic("Failed to create_note", err)
		}
		_, err = db.Exec(create_msg)
		if err != nil {
			log.Panic("Failed to create_msg", err)
		}
	}

	Dbx = &DBTx{db}
	maxuserid := Dbx.maxUser()
	fmt.Println("maxuserid===", maxuserid)
	funcUser = getUserID(maxuserid)
	fmt.Println("db init ok")
}

func (tx DBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.db.Exec(query, args...)
}

func (tx DBTx) maxUser() int {
	rows, err := tx.db.Query("select max(user_id) mid from t_user")
	if err != nil {
		log.Panic("Failed to query max user", err)
	}
	defer rows.Close()
	var id sql.NullInt32
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Panic("Failed to Scan ", err)
		}
	}
	if id.Valid {
		return int(id.Int32)
	} else {
		return 0
	}
}

func (tx DBTx) SaveUser(name, email, passwd string) error {
	userid := funcUser()
	_, err := tx.db.Exec("insert into t_user(user_id,name,email,passwd,role) values(?,?,?,?,0)",
		userid, name, email, passwd)
	if err != nil {
		fmt.Println("Failed to insert into t_user", err)
		return err
	}
	return nil
}

func (tx DBTx) UserLogin(email, passwd string) (bool, error) {
	rows, err := tx.db.Query("select 1 from t_user where email=? and passwd=?",
		email, passwd)
	if err != nil {
		fmt.Println("Failed to  Query t_user", err)
		return false, err
	}
	return rows.Next(), nil
}
