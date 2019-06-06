package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type queryResolver struct{ *Server }

func (r *queryResolver) Character(ctx context.Context, input *int) (*maco.Character, error) {
	panic("not implemented")
}

func (r *queryResolver) Characters(ctx context.Context) ([]*maco.Character, error) {
	panic("not implemented")
}

func (r *queryResolver) Comic(ctx context.Context, input *int) (*maco.Comic, error) {
	panic("not implemented")
}

func (r *queryResolver) Comics(ctx context.Context) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *queryResolver) Creator(ctx context.Context, input *int) (*maco.Creator, error) {
	panic("not implemented")
}

func (r *queryResolver) Creators(ctx context.Context) ([]*maco.Creator, error) {
	panic("not implemented")
}

func (r *queryResolver) Event(ctx context.Context, input *int) (*maco.Event, error) {
	panic("not implemented")
}

func (r *queryResolver) Events(ctx context.Context) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *queryResolver) Serie(ctx context.Context, input *int) (*maco.Series, error) {
	panic("not implemented")
}

func (r *queryResolver) Series(ctx context.Context) ([]*maco.Series, error) {
	panic("not implemented")
}

func (r *queryResolver) Story(ctx context.Context, input *int) (*maco.Story, error) {
	panic("not implemented")
}

func (r *queryResolver) Stories(ctx context.Context) ([]*maco.Story, error) {
	panic("not implemented")
}
