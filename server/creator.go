package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type creatorResolver struct{ *Server }

func (r *creatorResolver) Comics(ctx context.Context, obj *maco.Creator) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *creatorResolver) Events(ctx context.Context, obj *maco.Creator) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *creatorResolver) Series(ctx context.Context, obj *maco.Creator) ([]*maco.Series, error) {
	panic("not implemented")
}

func (r *creatorResolver) Stories(ctx context.Context, obj *maco.Creator) ([]*maco.Story, error) {
	panic("not implemented")
}
