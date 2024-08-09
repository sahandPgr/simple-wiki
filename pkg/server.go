package mhttp

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type IServer interface {
	initializeHandlerFunctions() error
	ListenAndServe()
}

type ServerConfig struct {
	staticDirectory string
	url             string
	port            string
}

// Define server
func NewServer(staticDir string, url string, port string) ServerConfig {
	var serverConfig = ServerConfig{
		staticDirectory: staticDir,
		url:             url,
		port:            port,
	}
	return serverConfig
}

var templates *template.Template
var validPath *regexp.Regexp

// Initializing Handlers
func (fs ServerConfig) InitializeHandlerFunctions() error {
	staticDirectory = fs.staticDirectory
	fmt.Println("Initializing the handler functions...")

	if err := validateFolder(staticDirectory); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
	}
	// Initialize the wiki variables
	templates = template.Must(template.ParseFiles(staticDirectory+"/edit.html", staticDirectory+"/view.html"))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

	http.HandleFunc("/", FrontPageHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	fmt.Println("Handler functions have been initialized")
	return nil
}

// Listening the port
func (fs ServerConfig) ListenAndServe() {
	serverUrl := fs.url + ":" + fs.port
	fmt.Printf("Listening on %s...\n", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, nil))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// Check Folder Path
func validateFolder(folderPath string) error {
	fileInfo, err := os.Stat(folderPath)

	if err != nil {
		return err
	}

	if fileInfo != nil && !fileInfo.IsDir() {
		return fmt.Errorf("%s : is not a directory", folderPath)
	}
	return nil
}
