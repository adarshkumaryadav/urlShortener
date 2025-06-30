// this is in-memory model
package model

import "sync"

// using map to store the original URL -> shorten URL
var URLToShort map[string]string

// using map to store the shorten URL -> Original URL
// so redirecting will be done by looking into this map
var ShortToURL map[string]string

var Mu sync.Mutex

// using RWMutex for locking shortToURL map as this will be frequently read wrt writing
var RMu sync.RWMutex
