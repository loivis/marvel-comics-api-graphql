package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type eventResolver struct{ *Server }

func (r *eventResolver) Characters(ctx context.Context, obj *maco.Event) ([]*maco.Character, error) {
	panic("not implemented")
}

func (r *eventResolver) Comics(ctx context.Context, obj *maco.Event) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *eventResolver) Creators(ctx context.Context, obj *maco.Event) ([]*maco.Creator, error) {
	panic("not implemented")
}

func (r *eventResolver) Series(ctx context.Context, obj *maco.Event) ([]*maco.Series, error) {
	panic("not implemented")
}

func (r *eventResolver) Stories(ctx context.Context, obj *maco.Event) ([]*maco.Story, error) {
	panic("not implemented")
}
