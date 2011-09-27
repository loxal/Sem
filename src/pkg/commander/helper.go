// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commander

import (
    "http"
    "json"
    "appengine"
    "appengine/memcache"
)

func addToCache(r *http.Request, cmd *Cmd) {
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
}

func testExecCmdFromCache(r *http.Request, cmdName string) {
//c := appengine.NewContext(r)

}