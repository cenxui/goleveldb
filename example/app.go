package main

import (
	"github.com/cenxui/goleveldb/leveldb"
	"github.com/cenxui/goleveldb/leveldb/errors"
	"fmt"
)

func main()  {
	db, err := leveldb.OpenFile("./db", nil)

	if err != nil {
		errors.New(err.Error())
	}

	err = db.Put([]byte("key"), []byte("value"), nil)

	err = db.Put([]byte("john"), []byte("12"), nil)

	err = db.Put([]byte("Mia"), []byte("123拗好ㄏㄠ"), nil)

	err = db.Put([]byte("Lisa"), []byte("42"), nil)

  	for i:=1; i<10 ; i++ {
		err = db.Put([]byte(fmt.Sprintf("peter%d", i)), []byte("42"), nil)
	}



	defer db.Close()
}
