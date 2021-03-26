package main

import (
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `portfor %s

portfor returns a deterministic hash for a given argument suitable for use as an 
unpriviledged port number. As a special case, it trims "www." or "local." prefixes 
from its argument.

Options:

`, getVersion())
		flag.PrintDefaults()
	}
	l := flag.Bool("l", false, "output localhost address (change string format with $PORTFOR_LOCALHOST)")
	flag.Parse()
	port := PortFor(flag.Arg(0))
	format := "%d"
	if *l {
		if format = os.Getenv("PORTFOR_LOCALHOST"); format == "" {
			format = "http://localhost:%d"
		}
	}
	s := fmt.Sprintf(format, port)
	fmt.Println(s)
}

// PortFor returns a deterministic port number for a given string.
// As a special case, it trims "www." or "local." prefixes from its argument.
// Hashing is based on MD5 and not suitable for adversarial environments.
func PortFor(name string) int {
	const minport = 1024

	name = strings.TrimPrefix(name, "local.")
	name = strings.TrimPrefix(name, "www.")

	hash := md5.New()
	hash.Write([]byte(name))
	port := binary.BigEndian.Uint16(hash.Sum(nil))
	port %= 1<<16 - minport
	return int(port) + minport
}

func getVersion() string {
	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "(unknown)"
	}
	return i.Main.Version
}
