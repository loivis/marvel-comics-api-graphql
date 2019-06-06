//go:generate go run github.com/99designs/gqlgen

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/tabwriter"

	"github.com/TV4/env"

	"github.com/loivis/marvel-comics-api-graphql/mongodb"
	"github.com/loivis/marvel-comics-api-graphql/server"
)

func main() {
	conf := readConfig()

	mux := http.NewServeMux()

	store := mongodb.New()

	server.New(store).Register(mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.port)

	log.Fatal(http.ListenAndServe(":"+conf.port, mux))
}

type config struct {
	mongodbURI      string
	mongodbDatabase string
	port            string
}

func readConfig() *config {
	return &config{
		mongodbURI:      env.String("MONGODB_URI", ""),
		mongodbDatabase: env.String("MONGODB_DATABASE", "maco"),
		port:            env.String("PORT", "8080"),
	}
}

func (c *config) String() string {
	hideIfSet := func(v interface{}) string {
		s := ""

		switch typedV := v.(type) {
		case string:
			s = typedV
		case []string:
			s = strings.Join(typedV, ",")
		case fmt.Stringer:
			if typedV != nil {
				s = typedV.String()
			}
		}

		if s != "" {
			return "<hidden>"
		}
		return ""
	}

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 1, 4, ' ', 0)
	for _, e := range []struct {
		k string
		v interface{}
	}{
		{"MONGODB_URI", hideIfSet(c.mongodbURI)},
		{"MONGODB_DATABASE", c.mongodbDatabase},
		{"PORT", c.port},
	} {
		fmt.Fprintf(w, "%s\t%v\n", e.k, e.v)
	}
	w.Flush()
	return buf.String()
}
