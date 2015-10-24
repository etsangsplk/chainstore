package logmgr

import (
	"log"
	"os"
	"testing"

	"github.com/pressly/chainstore"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestLogMgrStore(t *testing.T) {
	var store chainstore.Store
	var err error

	ctx, _ := context.WithCancel(context.Background())

	assert := assert.New(t)

	logger := log.New(os.Stdout, "", 0)

	store = chainstore.New(New(logger, "test"))
	err = store.Open()
	assert.Nil(err)
	defer store.Close()

	// Put a bunch of objects
	e1 := store.Put(ctx, "hi", []byte{1, 2, 3})
	e2 := store.Put(ctx, "bye", []byte{4, 5, 6})
	assert.Nil(e1)
	assert.Nil(e2)

	// Get those objects
	_, e1 = store.Get(ctx, "hi")
	_, e2 = store.Get(ctx, "bye")
	assert.Equal(e1, nil)
	assert.Equal(e2, nil)

	// Delete those objects
	e1 = store.Del(ctx, "hi")
	e2 = store.Del(ctx, "bye")
	assert.Equal(e1, nil)
	assert.Equal(e2, nil)
}
