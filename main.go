package main

import "github.com/khosimorafo/mongodbi/mongodbi"

var url = "mongodb://mastende:mastende@ds115573.mlab.com:15573/mastende-test"
var dbname = "mastende-test"

type Thing struct{}

func main() {
	mongoDAL, _ := mongodbi.NewMongoDAL(url, dbname)

	mongodbi.Persist("testcollection", Thing{}, mongoDAL)
}
