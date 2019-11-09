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
