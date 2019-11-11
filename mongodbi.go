package mongodbi

import (
	"errors"

	"gopkg.in/mgo.v2"

	log "github.com/sirupsen/logrus"
)

// DAL is the data access layer for them mongo database
type DAL interface {
	Insert(collectionName string, docs ...interface{}) error
}

// MongoDAL is an implementation of DataAccessLayer for a defined above
type MongoDAL struct {
	session *mgo.Session
	db      *mgo.Database
	dbName  string
}

// NewMongoDAL creates a MongoDAL
func NewMongoDAL(dbURI string, dbName string) (DAL, error) {

	if dbURI == "" {
		log.WithFields(log.Fields{
			"URI": dbURI,
		}).Error("Invalid URI string provided.")

		return nil, errors.New("Invalid URI provided")
	}

	session, err := mgo.Dial(dbURI)
	if err != nil {
		log.WithFields(log.Fields{
			"error_message": err.Error(),
		}).Info("Database onnection error.")
	}

	mongo := &MongoDAL{
		session: session,
		db:      session.DB(dbName),
		dbName:  dbName,
	}

	return mongo, err
}

// c is a helper method to get a collection from the session
func (m *MongoDAL) c(collection string) *mgo.Collection {
	return m.session.DB(m.dbName).C(collection)
}

// Insert stores documents in mongo
func (m *MongoDAL) Insert(collectionName string, docs ...interface{}) error {

	log.WithFields(log.Fields{
		"collection": collectionName,
		"documents":  docs,
	}).Info("Inserting record into db...")

	return m.c(collectionName).Insert(docs)
}

// Persist new record into database
func Persist(collection string, doc interface{}, data DAL) interface{} {

	data.Insert(collection, doc)
	return doc
}
