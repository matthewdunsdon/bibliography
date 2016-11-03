# matthewdunsdon/bibliography
[![GoDoc](https://godoc.org/github.com/matthewdunsdon/bibliography?status.svg)](https://godoc.org/github.com/matthewdunsdon/bibliography)
[![Build Status](https://travis-ci.org/matthewdunsdon/bibliography.svg?branch=master)](https://travis-ci.org/matthewdunsdon/bibliography)

Package `matthewdunsdon/bibliography` implements a request router and dispatcher for matching incoming requests to
their respective handler.

The name mux stands for "HTTP request multiplexer". Like the standard `http.ServeMux`, `mux.Router` matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions. The main features are:

* It implements the `http.Handler` interface so it is compatible with the standard `http.ServeMux`.
* Requests can be matched based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.
* URL hosts and paths can have variables with an optional regular expression.
* Registered URLs can be built, or "reversed", which helps maintaining references to resources.
* Routes can be used as subrouters: nested routes are only tested if the parent route matches. This is useful to define groups of routes that share common conditions like a host, a path prefix or other repeated attributes. As a bonus, this optimizes request matching.

---

* [Install](#install)
* [Examples](#examples)
* [Matching Routes](#matching-routes)
* [Static Files](#static-files)
* [Registered URLs](#registered-urls)
* [Full Example](#full-example)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/matthewdunsdon/bibliography
```

## Examples

Let's start registering a couple of URL paths and handlers:

```go
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
```

Here we register three routes mapping URL paths to handlers. This is equivalent to how `http.HandleFunc()` works: if an incoming request URL matches one of the paths, the corresponding handler is called passing (`http.ResponseWriter`, `*http.Request`) as parameters.

Paths can have variables. They are defined using the format `{name}` or `{name:pattern}`. If a regular expression pattern is not defined, the matched variable will be anything until the next slash. For example:

...

### Matching Routes

Routes can also be restricted to a domain or subdomain. Just define a host pattern to be matched. They can also have variables:

```go
r := mux.NewRouter()
// Only matches if domain is "www.example.com".
r.Host("www.example.com")
// Matches a dynamic subdomain.
r.Host("{subdomain:[a-z]+}.domain.com")
```

## Full Example

Here's a complete, runnable example of a small `mux` based server:

```go
package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
```

## License

MIT licensed. See the LICENSE file for details.
