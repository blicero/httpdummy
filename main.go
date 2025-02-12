// Time-stamp: <2025-02-12 19:20:59 krylon>

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var l *log.Logger

func main() {
	var (
		err           error
		addr, logpath string
		router        *mux.Router
		srv           http.Server
		fh            *os.File
		multi         io.Writer
	)

	flag.StringVar(&addr, "addr", "[::]:8086", "The address to listen on")
	flag.StringVar(&logpath, "log", "httpdummy.log", "The path of the log file")
	flag.Parse()

	if fh, err = os.Create(logpath); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Cannot open log at %s: %s\n",
			logpath,
			err.Error())
		os.Exit(1)
	}

	defer fh.Close()

	multi = io.MultiWriter(fh, os.Stdout)

	l = log.New(multi, "http ", log.Ldate|log.Ltime|log.Lshortfile)

	router = mux.NewRouter()
	srv.Addr = addr
	srv.Handler = router
	srv.ErrorLog = l

	router.HandleFunc("/{req:(?:.*)}", handleHTTP)

	if err = srv.ListenAndServe(); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to start web server: %s\n",
			err.Error())
		os.Exit(1)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	l.Printf("Handle request for %s from %s\n",
		r.URL.EscapedPath(),
		r.RemoteAddr)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(403)
	w.Write([]byte("Aloha\r\n\r\n"))
}
