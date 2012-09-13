package server

import (
	"appengine"
	"mime"
	"net/http"
	"server/fileloader"
	"strings"
)

var ROOT_DIR string = "homepage"

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var p string = "homepage"
	if r.URL.Path == "/" {
		p += "/index.html"
	} else {
		p += r.URL.Path
	}

	b, e := fileloader.GetFile(c, p)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	writeContentType(w, p)
	w.Write(b)
}

func writeContentType(w http.ResponseWriter, p string) {
	i := strings.LastIndex(p, ".")
	if i == -1 {
		return
	}

	mType := mime.TypeByExtension(p[i:])

	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec7.html#sec7
	if mType != "" {
		w.Header().Add("Content-Type", mType)
	}
}
