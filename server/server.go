package server

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

// Server implements macogql.ResolverRoot
type Server struct {
	store macogql.Store
}

func New(store macogql.Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Character() macogql.CharacterResolver {
	return &characterResolver{s}
}

func (s *Server) Comic() macogql.ComicResolver {
	return &comicResolver{s}
}

func (s *Server) ComicPrice() macogql.ComicPriceResolver {
	return &comicPriceResolver{s}
}

func (s *Server) Creator() macogql.CreatorResolver {
	return &creatorResolver{s}
}

func (s *Server) Event() macogql.EventResolver {
	return &eventResolver{s}
}

func (s *Server) Query() macogql.QueryResolver {
	return &queryResolver{s}
}

func (s *Server) Series() macogql.SeriesResolver {
	return &seriesResolver{s}
}

func (s *Server) Story() macogql.StoryResolver {
	return &storyResolver{s}
}

func (s *Server) Register(mux *http.ServeMux) {
	mux.Handle("/", handler.Playground("GraphQL playground", "/query"))
	mux.Handle("/query", handler.GraphQL(macogql.NewExecutableSchema(macogql.Config{Resolvers: s})))
}

func fromTo(offset, length, max int) (int, int) {
	from, to := offset, offset+length

	if from > max {
		from = max
	}

	if to > max {
		to = max
	}

	return from, to
}
