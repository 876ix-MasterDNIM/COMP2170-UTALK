package htmlserver

import (
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)

// ServeHTML serves html file
func ServeHTML(path string, response http.ResponseWriter) {
	page, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	} else {
		response.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
		response.Write(page)
	}
}
