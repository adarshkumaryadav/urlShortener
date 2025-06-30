// core logic to short the URL and map
package utils

import (
	"adarshkumaryadav/urlShortener/model"
)

func GenerateShortURL(url string) (err error) {
	// dummy code here for the first iteration
	const dummy string = "localhost:8080/first"
	model.Mu.Lock()
	defer model.Mu.Unlock()
	model.URLToShort[url] = dummy

	model.RMu.Lock()
	defer model.RMu.Unlock()
	model.ShortToURL[dummy] = url

	return nil
}
