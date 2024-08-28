package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"time"
)

// parameters
const listenPort = "8080"
const serveDirectoryPath = "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF"

func main() {
	var files = dirList(serveDirectoryPath)
	if len(files) == 0 {
		panicOnError("No files to serve", fmt.Errorf("len(files) == 0"))
	}
	slices.Sort(files)

	var mapping = buildMapping(files)

	var listenOn = fmt.Sprintf(":%v", listenPort)

	var err = serve(mapping, listenOn)
	panicOnError("Serve failed", err)
}

func serve(mapping map[string]string, listenOn string) error {
	const fnameRequestKey = "fname"

	var handler = func(w http.ResponseWriter, r *http.Request) {
		log("Processing URI ", r.RequestURI)

		fname := r.PathValue(fnameRequestKey)
		fpath, exists := mapping[fname]

		if exists {
			log("Serving file ", fpath)
			http.ServeFile(w, r, fpath)
		} else {
			http.Error(w, fmt.Sprintf("No mapping from name '%v'", fname), http.StatusNotFound)
		}
	}

	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("GET /{%v}", fnameRequestKey), handler)
	http.Handle("/", mux) // set default handler

	log("Starting server at ", listenOn)
	return http.ListenAndServe(listenOn, nil) // nil: default handler
}

func buildMapping(names []string) map[string]string {
	var mapping = make(map[string]string, len(names))
	for idx, name := range names {
		var key = fmt.Sprintf("%v", idx+1)
		mapping[key] = name
		log("added mapping ", key, name)
	}
	return mapping
}

func dirList(path string) []string {
	var cleanPath string = cleanPath(path)

	entries, err := readDir(cleanPath)
	panicOnError("Can't read files from "+path, err)

	var lst = make([]string, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			log("skip directory ", entry.Name())
		} else {
			filePath := joinPathElems(cleanPath, entry.Name())
			lst = append(lst, filePath)
			log("file added: ", filePath)
		}
	}

	return lst
}

func cleanPath(path string) string {
	return filepath.Clean(path)
}

func joinPathElems(e1, e2 string) string {
	var joined = filepath.Join(e1, e2)
	return joined
}

func readDir(path string) ([]fs.DirEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	entries, err := file.ReadDir(0)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func panicOnError(msg string, err error) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}

// log writes message to standard output (fmt.Println). Message combined from prefix msg and slice of arbitrary values
func log(msg string, xs ...any) {
	var line = buildMessage(msg, xs...)
	fmt.Println(line)
}

func buildMessage(msg string, xs ...any) string {
	var line = ts() + ": " + msg // TODO: use buffer, not string concat

	for _, x := range xs {
		// https://pkg.go.dev/fmt
		// line += fmt.Sprintf("%T(%v); ", x, x) // type(value)
		line += fmt.Sprintf("%#v; ", x) // repr
	}

	return line
}

// ts returns current timestamp in RFC3339 with milliseconds
func ts() string {
	/*
		https://pkg.go.dev/time#pkg-constants
		https://stackoverflow.com/questions/35479041/how-to-convert-iso-8601-time-in-golang
	*/
	const (
		RFC3339      = "2006-01-02T15:04:05Z07:00"
		RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"
	)
	return time.Now().UTC().Format(RFC3339Milli)
}
