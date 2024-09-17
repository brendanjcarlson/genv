// genv is a Go package for loading and accessing environment variables.
//
// It simplifies the processes of lookup, casting, and populating configuration structs, using generics, struct tags, and reflection for mapping.
//
// Use:
//
//	package main
//
//	import "github.com/brendanjcarlson/genv"
//
//	func main() {
//	    err := genv.Load(".env")
//	    if err != nil {
//	        log.Fatalf("load env: %v", err)
//	    }
//
//	    superSecretKey, err := genv.Get[string]("SUPER_SECRET_KEY")
//	    if err != nil {
//	        log.Fatalln(err)
//	    }
//
//	    timeoutSeconds, err := genv.Get[int]("TIMEOUT_SECONDS")
//	    if errors.Is(err, genv.ErrCannotCast) {
//	        log.Fatalln(err)
//	    } else {
//	        timeoutSeconds = 5
//	    }
//
//	    mustHaveValue := genv.GetOrPanic[string]("GOTTA_HAVE_THIS_ONE")
//
//	    type ServerConfig struct {
//	        Host string `genv:"SERVER_HOST"`
//	        Port string `genv:"SERVER_PORT"`
//	    }
//
//	    var serverConfig ServerConfig
//	    if err := genv.GetStruct(&serverConfig); err !=nil {
//	        log.Fatalln(err)
//	    }
//	}
package genv
