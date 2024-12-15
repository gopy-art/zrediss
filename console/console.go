package console

import (
	"flag"
	"fmt"
	"os"
)

var (
	URL    string
	Action string
	Key    string
	Set    string
	Value  string
	Logger string
)

var AppVersion string = "v1.0.0"

// init the flags and dependencies throw the flag
func InitFlags() {
	flag.StringVar(&URL, "u", "", "set rediss url")
	flag.StringVar(&Key, "k", "", "set rediss key to get the value")
	flag.StringVar(&Set, "s", "", "set value for the specific key")
	flag.StringVar(&Logger, "l", "stdout", "set app logger type , stdout or file")
	flag.StringVar(&Value, "val", "", "set rediss value for set the value")
	flag.StringVar(&Action, "action", "", "set the action of the module. (all, get, set, delete, flush)")
	version := flag.Bool("v", false, "zrediss version")
	flag.Parse()

	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}
}
