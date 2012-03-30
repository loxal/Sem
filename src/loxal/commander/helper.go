// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commander

import (
    "net/http"
    "encoding/json"
    "appengine"
    "appengine/memcache"
)

func addCacheItem(r *http.Request, cmd *Cmd) {
    c := appengine.NewContext(r)
    cmdJson, _ := json.Marshal(cmd)
    cmdItem := &memcache.Item {
        Key: cmd.Name,
        Value: []byte(cmdJson),
    }

    // Add the item to the memcache, if the key does not already exist
    if err := memcache.Add(c, cmdItem); err == memcache.ErrNotStored {
        c.Debugf("item with key %q already exists", cmdItem.Key)
    } else if err != nil {
        c.Debugf("error adding item: %v", err)
    }
    c.Debugf("CACHED")
}

func updateCacheItem(r *http.Request) {
//    c := appengine.NewContext(r)
//    // Change the Value of the item
//    item.Value = []byte("Where the buffalo roam")
//    // Set the item, unconditionally
//    if err := memcache.Set(c, item); err != nil {
//        c.Debugf("error setting item: %v", err)
//    }
}

func getCacheItem(r *http.Request, cmdName string) {
    c := appengine.NewContext(r)

    if item, err := memcache.Get(c, cmdName); err == memcache.ErrCacheMiss {
        c.Debugf("item not in cache: %q", cmdName)
    } else if err != nil {
        c.Debugf("error getting item: %v", err)
    } else {
        c.Debugf("cache item found: %q", item.Value)
    }
}