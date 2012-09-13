package fileloader

import (
	"appengine"
	"appengine/memcache"
	"io/ioutil"
)

func GetFile(c appengine.Context, p string) ([]byte, error) {
	item, e := memcache.Get(c, p)

	if e == memcache.ErrCacheMiss {
		return storeInCache(c, p)
	} else if e != nil {
		return nil, e
	}

	return item.Value, nil
}

// Read file from disk and cache it.
func storeInCache(c appengine.Context, p string) ([]byte, error) {
	b, e := ioutil.ReadFile(p)

	if e != nil {
		c.Errorf("Reading file %v caused %v", p, e)
		return nil, e
	}

	newItem := &memcache.Item{
		Key:   p,
		Value: b,
	}

	return b, memcache.Set(c, newItem)
}
