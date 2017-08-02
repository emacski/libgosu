// gosu cli utility
package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/emacski/libgosu"
)

func init() {
	// make sure we only have one process and that it runs on the main thread
	// (so that ideally, when we Exec, we keep our user switches and stuff)
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {
	log.SetFlags(0) // no timestamps on our logs

	if len(os.Args) <= 2 {
		self := filepath.Base(os.Args[0])
		log.Printf("Usage: %s user-spec command [args]", self)
		log.Printf("   ie: %s tianon bash", self)
		log.Printf("       %s nobody:root bash -c 'whoami && id'", self)
		log.Printf("       %s 1000:1 id", self)
		log.Println()
		log.Printf("%s version: %s (%s on %s/%s; %s)", self, libgosu.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler)
		log.Printf("%s license: GPL-3 (full text at https://github.com/emacski/libgosu)\n", strings.Repeat(" ", len(self)))
		log.Println()
		os.Exit(1)
	}

	err := libgosu.Exec(os.Args[1], os.Args[2:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
