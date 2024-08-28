package main

import (
	"fmt"
	"net/http"
)

func main() {
	const listenOn = ":8080"
	const fnameKey = "fname"

	var mapping = map[string]string{
		"1": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E01.WEB-DLRip.RGzsRutracker.avi",
		"2": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E02.WEB-DLRip.RGzsRutracker.avi",
		"3": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E03.WEB-DLRip.RGzsRutracker.avi",
		"4": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E04.WEB-DLRip.RGzsRutracker.avi",
		"5": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E05.WEB-DLRip.RGzsRutracker.avi",
		"6": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E06.WEB-DLRip.RGzsRutracker.avi",
		"7": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E07.WEB-DLRip.RGzsRutracker.avi",
		"8": "/mnt/c/Users/valik/Downloads/Reacher.S02.WEB-DLRip.LF/Reacher.S02E08.WEB-DLRip.RGzsRutracker.avi",
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Processing %v ...", r.RequestURI)

		fname := r.PathValue(fnameKey)
		fpath, exists := mapping[fname]

		if exists {
			http.ServeFile(w, r, fpath)
		} else {
			http.Error(w, fmt.Sprintf("No mapping from name '%v'", fname), http.StatusNotFound)
		}
	}

	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("GET /{%v}", fnameKey), handler)
	http.Handle("/", mux)

	fmt.Printf("Starting server at %v\n", listenOn)
	http.ListenAndServe(listenOn, nil)
}
