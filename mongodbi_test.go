package mongodbi_test

import (
	"testing"

	dbi "github.com/khosimorafo/mongodbi"
	"github.com/khosimorafo/mongodbi/mocks"
	"github.com/stretchr/testify/assert"
)

type Thing struct{}

func TestInsert(t *testing.T) {

	mockThing := Thing{}

	mockDAL := &mocks.DAL{}
	mockDAL.On("Insert", "foo", mockThing).Return(nil)

	actual := dbi.Persist("foo", mockThing, mockDAL)

	mockDAL.AssertExpectations(t)

	assert.Equal(t, mockThing, actual, "should return a Thing")
}

/*
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
*/
