package GoWheels

import (
	"github.com/tecbot/gorocksdb"
	"fmt"
	"log"
)

func RocksDemo() {
	dbOptions := gorocksdb.NewDefaultOptions()
	dbOptions.SetCreateIfMissing(true)

	db, e := gorocksdb.OpenDb(dbOptions, "test_db")
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()

	//simple put
	wopts := gorocksdb.NewDefaultWriteOptions()
	defer wopts.Destroy()
	db.Put(wopts, []byte("a:0"), []byte("b:0"))

	//batch write
	writeBatch := gorocksdb.NewWriteBatch()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("a:%d", i)
		val := fmt.Sprintf("b:%d", i)
		writeBatch.Put([]byte(key), []byte(val))
	}
	db.Write(wopts, writeBatch)

	//get
	ropts := gorocksdb.NewDefaultReadOptions()
	defer ropts.Destroy()

	val, e := db.Get(ropts, []byte("a:0"))
	if e != nil {
		log.Printf("get a:0 error : %s", e)
	}else{
		log.Println(string(val.Data()))
		val.Free()
	}

	//iterator
	log.Println("===================iterator===================")
	ropts.SetFillCache(false)
	iterator := db.NewIterator(ropts)
	for iterator.Seek([]byte("a:1")); iterator.Valid(); iterator.Next() {
		key := string(iterator.Key().Data())
		if key > "a:5" {
			break
		}
		log.Println(fmt.Sprintf("%s -> %s", iterator.Key().Data(), iterator.Value().Data()))
		//释放资源
		iterator.Key().Free()
		iterator.Value().Free()
	}

	e = iterator.Err()
	if e != nil {
		log.Println("iterator error during scan ", e)
	}

	//删除
	e = db.Delete(wopts, []byte("a9"))
	if e != nil{
		log.Println("delete a9 error , ", e)
	}

}
