// Time-stamp: <2025-02-12 19:05:16 krylon>

package main

import (
	"flag"
	"fmt"
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
	)

	flag.StringVar(&addr, "addr", "[::]:8086", "The address to listen on")
	flag.StringVar(&logpath, "log", "httpot.log", "The path of the log file")
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

	l = log.New(fh, "http ", log.Ldate|log.Ltime|log.Lshortfile)

	router = mux.NewRouter()
	srv.Addr = addr
	srv.Handler = router
	srv.ErrorLog = l

	router.HandleFunc(".*", handleHTTP)

	srv.ListenAndServe()
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	l.Printf("Handle request for %s from %s\n",
		r.URL.EscapedPath(),
		r.RemoteAddr)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(403)
	w.Write([]byte("Aloha\r\n\r\n"))
}
