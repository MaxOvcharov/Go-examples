package main

import "fmt"
import "net"
import "net/url"

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println("SCHEME: ", u.Scheme)

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.
	fmt.Println("USER: ",u.User)
	fmt.Println("USER_NAME: ",u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("PASSWORD: ",p)

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.
	fmt.Println("HOST: ",u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("HOST_NAME: ",host)
	fmt.Println("PORT: ",port)

	// Here we extract the `path` and the fragment after
	// the `#`.
	fmt.Println("PATH: ",u.Path)
	fmt.Println("FRAGMENT_#: ",u.Fragment)

	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.
	fmt.Println("RAW_QUERY: ",u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("PARSED_QUERY: ",m)
	fmt.Println("SOME_QUERY_PARAMETR: ",m["k"][0])
}
