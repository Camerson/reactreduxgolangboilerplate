package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"reactgoboilerplate/controllers"
)

// Router is exported and used in main.go

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	fmt.Println(path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func Router() *mux.Router {

	router := mux.NewRouter()

	apiRoutes := router.PathPrefix("/api/v1").Subrouter()
	apiRoutes.HandleFunc("/users", controllers.CreateUser).Methods("POST", "OPTIONS")
	apiRoutes.HandleFunc("/users", controllers.GetUsers).Methods("GET", "OPTIONS")
	apiRoutes.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET", "OPTIONS")
	apiRoutes.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT", "OPTIONS")
	apiRoutes.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")
	apiRoutes.Use(CheckAuthorization)

	spa := spaHandler{staticPath: "frontend/build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	return router
}
