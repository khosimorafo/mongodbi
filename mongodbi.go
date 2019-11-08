package mongodbi

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var a App

// App app
type App struct {
	Session    *mgo.Session
	Collection *mgo.Collection
	Database   *mgo.Database
}

// DB db
func DB(uri string, db string) *App {

	a.Session = AppCollection(uri)
	a.Database = a.Session.DB(db)

	return &a
}

// SetCollection s
func (a *App) SetCollection(coll string) {

	a.Collection = a.Database.C(coll)
}

// AppCollection a
func AppCollection(uri string) *mgo.Session {

	//uri := "mongodb://mastende:mastende@ds115573.mlab.com:15573/mastende-test"
	if uri == "" {
		log.Error("No database url string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {

		log.WithFields(log.Fields{
			"error_message": err,
		}).Error("Failed to connect to mongo server.")
		os.Exit(1)
	}

	//defer sess.Close()

	//sess.SetSafe(&mgo.Safe{})

	return sess
}

// Update u
func Update(db *App, sel interface{}, change interface{}) error {

	if err := db.Collection.Update(sel, change); err != nil {
		return errors.New(err.Error())
	}
	return nil
}
