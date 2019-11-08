package mongodbi_test

import (
	"os"
	"testing"

	dbi "github.com/khosimorafo/mongodbi"

	assert "github.com/stretchr/testify/assert"
)

var a dbi.App

var url = "mongodb://mastende:mastende@ds115573.mlab.com:15573/mastende-test"
var dbname = "mastende-test"

// TestMain wraps all tests with the needed initialized mock DB and fixtures
func TestMain(m *testing.M) {

	// Init test session/db/collection
	//configor.Load(&mongodbi.Config, "config.yml")
	a = *dbi.DB(url, dbname)
	a.Database.C("testcollection").DropCollection()

	// Run the test suite
	retCode := m.Run()

	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session
	a.Session.Close()

	// call with result of m.Run()
	os.Exit(retCode)
}

func TestCollectionCreate(t *testing.T) {

	a.SetCollection("testcollection")

	c := a.Collection.Name

	assert.Equal(t, c, "testcollection", "the should be equal")
}
