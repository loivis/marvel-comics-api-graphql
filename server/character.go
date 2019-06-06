package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type characterResolver struct{ *Server }

func (r *characterResolver) Comics(ctx context.Context, obj *maco.Character) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *characterResolver) Events(ctx context.Context, obj *maco.Character) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *characterResolver) Series(ctx context.Context, obj *maco.Character) ([]*maco.Series, error) {
	panic("not implemented")
}

func (r *characterResolver) Stories(ctx context.Context, obj *maco.Character) ([]*maco.Story, error) {
	panic("not implemented")
}
