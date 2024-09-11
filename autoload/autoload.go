package autoload

import "github.com/brendanjcarlson/genv"

func init() {
	genv.LoadOrPanic(".env")
}
