package dbs

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

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
	update_time timestamp,
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

func getSequenceID(begin int64) func() int64 {
	i := begin
	return func() int64 {
		i += 1
		return i
	}
}

var funcUser func() int64
var funcNote func() int64

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
	funcUser = getSequenceID(maxuserid)
	funcNote = getSequenceID(0)
	fmt.Println("db init ok")
}

func (tx DBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.db.Exec(query, args...)
}

func (tx DBTx) maxUser() int64 {
	rows, err := tx.db.Query("select max(user_id) mid from t_user")
	if err != nil {
		log.Panic("Failed to query max user", err)
	}
	defer rows.Close()
	var id sql.NullInt64
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Panic("Failed to Scan ", err)
		}
	}
	if id.Valid {
		return id.Int64
	} else {
		return 0
	}
}

func (tx DBTx) maxNoteID() int64 {
	rows, err := tx.db.Query("select max(user_id) mid from t_user")
	if err != nil {
		log.Panic("Failed to query max user", err)
	}
	defer rows.Close()
	var id sql.NullInt64
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Panic("Failed to Scan ", err)
		}
	}
	if id.Valid {
		return id.Int64
	} else {
		return 0
	}
}

func (tx DBTx) SaveUser(name, email, passwd string) error {
	userid := funcUser()
	_, err := tx.db.Exec("insert into t_user(user_id,name,email,passwd,role,editor,avatar) values(?,?,?,?,0,?,?)",
		userid, name, email, passwd, "default", "images/logo.png")
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

func (tx DBTx) GetUser(email, passwd string) (User, error) {
	var user User
	rows, err := tx.db.Query("select user_id,name,email,avatar,role,editor from t_user where email=? and passwd=?",
		email, passwd)
	if err != nil {
		fmt.Println("Failed to  Query t_user", err)
		return user, err
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Avatar, &user.Role, &user.Editor)
		return user, err
	}
	return user, errors.New("user not exists")
}

func (tx DBTx) GetUserByEmail(email string) (User, error) {
	var user User
	rows, err := tx.db.Query("select user_id,name,email,avatar,role,editor from t_user where email=?", email)
	if err != nil {
		fmt.Println("Failed to  Query t_user", err)
		return user, err
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Avatar, &user.Role, &user.Editor)
		return user, err
	}
	return user, errors.New("user not exists")
}

func (tx DBTx) GetNoteKey() string {
	id := funcNote()
	year, month, day := time.Now().Date()
	key := fmt.Sprintf("%04d%02d%02d%04d", year, month, day, id)
	return key
}

func (tx DBTx) SaveNoteKey(userid int64, key string) error {

	_, err := tx.db.Exec("insert into t_note(user_id,note_key) values(?,?)",
		userid, key)
	if err != nil {
		fmt.Println("Failed to insert into t_note", err)
		return err
	}
	return nil
}

func (tx DBTx) SaveNoteInfo(n Note) error {

	_, err := tx.db.Exec("update t_note set title=?,summary=?,content=?,editor=?,files=? where note_key=?",
		n.Title, n.Summary, n.Content, n.Editor, n.Files, n.Key)
	if err != nil {
		fmt.Println("Failed to update t_note", err)
		return err
	}
	return nil
}

func (tx DBTx) GetNoteByKey(n *Note) error {

	rows, err := tx.db.Query("select user_id,title,summary,content,editor,files,update_time from t_note where note_key=?", n.Key)
	if err != nil {
		fmt.Println("Failed to  Query t_note", err)
		return err
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&n.UserID, &n.Title, &n.Summary, &n.Content, &n.Editor, &n.Files, &n.UpdatedAt)
		return err
	}
	return nil
}

func (tx DBTx) GetUserByID(id int) (User, error) {
	var user User
	rows, err := tx.db.Query("select user_id,name,email,avatar,role,editor from t_user where user_id=?", id)
	if err != nil {
		fmt.Println("Failed to  Query t_user", err)
		return user, err
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Avatar, &user.Role, &user.Editor)
		return user, err
	}
	return user, errors.New("user not exists")
}
