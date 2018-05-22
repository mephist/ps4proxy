/*
Inspired by https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c

Yet another ugly proxy to support PSN and hide from dumb proxy scanner
*/

package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	var bind string
	flag.StringVar(&bind, "bind", "127.0.0.1:8888", "Listen on the ip.add.re.ss:port")
	var proto string
	flag.StringVar(&proto, "proto", "http", "Proxy protocol (http or https)")
	var certPath string
	flag.StringVar(&certPath, "pem", "server.crt", "path to cert file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "server.key", "path to key file")
	//	var debug bool
	//	flag.BoolVar(&debug, "debug", false, "Debug requests to logfile (not in use)")
	var logFile string
	flag.StringVar(&logFile, "logfile", "", "path to log file")
	flag.Parse()

	if len(logFile) > 0 {
		f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//defer to close when you're done with it, not because you think it's idiomatic!
		defer f.Close()
		//set output of logs to f
		log.SetOutput(f)
	}

	log.Printf("Starting proxy on %s, proto: %s\n", bind, proto)

	StartProxy(bind, proto, certPath, keyPath)

}
