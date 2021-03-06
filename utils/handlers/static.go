package handlers

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/spaceuptech/space-cloud/modules/static"
)

// HandleStaticRequest creates a static request endpoint
func HandleStaticRequest(static *static.Module) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Path
		host := strings.Split(r.Host, ":")[0]

		route, ok := static.SelectRoute(host, url)
		if !ok {
			http.Error(w, "Path not found", http.StatusNotFound)
			return
		}

		path := strings.TrimPrefix(url, route.URLPrefix)
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		path = route.Path + path

		// Its a proxy request
		if route.Proxy != "" {
			addr := route.Proxy + path
			req, err := http.NewRequest(r.Method, addr, r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			// Set the http headers
			req.Header = make(http.Header)
			if contentType, p := r.Header["Content-Type"]; p {
				req.Header["Content-Type"] = contentType
			}

			// Make the http client request
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			defer res.Body.Close()

			reader := bufio.NewReader(res.Body)

			w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
			w.WriteHeader(res.StatusCode)
			reader.WriteTo(w)
			return
		}

		// Check if path exists
		if fileInfo, err := os.Stat(path); !os.IsNotExist(err) {
			// If path exists and is of type file then serve that file
			if !fileInfo.IsDir() {
				http.ServeFile(w, r, path)
				return
			}
			// Else if a index file exists within that folder serve that index file
			path = strings.TrimSuffix(path, "/")
			if _, err := os.Stat(path + "/index.html"); !os.IsNotExist(err) {
				http.ServeFile(w, r, path+"/index.html")
				return
			}
		}

		// If path does not exists serve the root index file
		http.ServeFile(w, r, strings.TrimSuffix(route.Path, "/")+"/index.html")
	}
}
