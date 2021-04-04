package main

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/ventu-io/go-shortid"
	"sync"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type UrlRow interface {
	Scan(...interface{}) error
}

type Sql interface {
	Begin() (*sql.Tx, error)
	Close() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) UrlRow
}

type SqlDatabase struct {
	sql Sql
}

func (sqlDb *SqlDatabase) GetSql() Sql {
	return sqlDb.sql
}

func InitDb() (*sql.DB){
	db, err := sql.Open("mysql", "root:Qwerty123@tcp(127.0.0.1:3306)/tinyurl")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return db
}

func (sqlDb *SqlDatabase) InsertIntoDB(key string, value string, date int64) error{
	t := time.Unix(date, 0)
	err := sqlDb.sql.QueryRow("insert into urlMap (original_url, shorten_url, expiry_date, created_at)values (?,?,?, now())", key, value, t)
	if err != nil {
		return errors.New("Error inserting into database")
	}
	return nil
}

// Store contains an LRU Cache
type Store struct {
	mutex *sync.Mutex
	store map[string]*list.Element
	listStore    *list.List
	buffer   int // Zero for unlimited
}

// Node maps a value to a key
type Node struct {
	key     string
	value   string
	expire  int64  // Unix time
}

var s *Store

func NewCache(size int) {
	s = New(size)
}

// Create new cache
func New(buffer int) *Store {
	s := &Store{
		mutex: &sync.Mutex{},
		store: make(map[string]*list.Element),
		listStore:    list.New(),
		buffer:   buffer,
	}
	return s
}

// Get a key from cache
func (s *Store) Get(key string) (string, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	current, exist := s.store[key]
	if exist {
		expire := int64(current.Value.(*Node).expire)
		if expire == 0 || expire > time.Now().Unix() {
			s.listStore.MoveToFront(current)
			return current.Value.(*Node).value, true
		}
	}
	return "", false
}

// Insert key to cache
func (s *Store) Insert(key string, value string, expire int64) {
	current, exist := s.store[key]
	if exist != true {
		s.store[key] = s.listStore.PushFront(&Node{
			key:    key,
			value:  value,
			expire: expire,
		})
		if s.buffer != 0 && s.listStore.Len() > s.buffer {
			s.Delete(s.listStore.Remove(s.listStore.Back()).(*Node).key)
		}
		return
	}
	current.Value.(*Node).value = value
	current.Value.(*Node).expire = expire
	s.listStore.MoveToFront(current)
}

// delete key from cache
func (s *Store) Delete(key string) {
	current, exist := s.store[key]
	if exist != true {
		return
	}
	s.listStore.Remove(current)
	delete(s.store, key)
}

// Flush all keys
func (s *Store) Empty() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store = make(map[string]*list.Element)
	s.listStore = list.New()
}

func main() {
	NewCache(10)
	db := InitDb()

	str := "https://eatigo.com/in/bengaluru/en/r/bamey-s-restro-cafe-kormangala-5006288"


	uniqueID, idError := shortid.Generate()
	if idError != nil {
		fmt.Println("Error")
	}

	fmt.Println(uniqueID)

	curr := time.Now()
	expirationTime := curr.Add(time.Hour * time.Duration(1))

	err := db.QueryRow("insert into urlMap (original_url, shorten_url, expiry_date, created_at)values (?,?,?, now())", str, uniqueID, expirationTime)
	if err != nil {
		fmt.Println("Error while inserting into database: %v", err)
	}

	s.Insert(uniqueID, str, 0)

	var originalURL string

	query := "SELECT original_url from urlMap where shorten_url=?"

	row := db.QueryRow(query, uniqueID)
	switch err := row.Scan(&originalURL); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("Database returned" + originalURL)
	default:
		panic(err)
	}

	if err != nil {
		fmt.Println("Error while fetching from database: %v", err)
	}
	cacheOutput, _ := s.Get(uniqueID)

	fmt.Println("Cache returned " +  cacheOutput)

	//fmt.Println(uniqueID)
}