package content

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Keys tracks the keys associated with some content.
type Keys struct {
	content  *Content
	uniqueID uint32
}

// Content returns the underlying content the keys were generated for
func (keys Keys) Content() *HarvestedResource {
	return keys.hr
}

// UniqueID returns the unique identifier based on key searching algorithm
func (keys Keys) UniqueID() uint32 {
	return keys.uniqueID
}

// UniqueIDText returns a unique identity key formatted as requested
func (keys Keys) UniqueIDText(format string) string {
	return fmt.Sprintf(format, keys.uniqueID)
}

// Slug returns the title of the content
func (keys Keys) Slug() string {
	return "TODO: Not implemented yet"
}

// KeyExists is a function passed in that checks whether a key already exists
type KeyExists func(random uint32, try int) bool

// GenerateUniqueID generates a unique identifier for this resource
func generateUniqueID(existsFn KeyExists) uint32 {
	nconflict := 0
	for i := 0; i < 10000; i++ {
		nextInt := nextRandomNumber()
		if !existsFn(nextInt, i) {
			return nextInt
		}

		if nconflict++; nconflict > 10 {
			randmu.Lock()
			rand = reseed()
			randmu.Unlock()
		}
	}

	// give up after max tries, not much we can do
	return nextRandomNumber()
}

// CreateKeys returns a new content keys object
func CreateKeys(c *Content, existsFn KeyExists) *Keys {
	result := new(Keys)
	result.content = c
	result.uniqueID = generateUniqueID(existsFn)
	return result
}

// Random number state, approach copied from tempfile.go standard library
var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextRandomNumber() uint32 {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return 1e9 + r%1e9
}
