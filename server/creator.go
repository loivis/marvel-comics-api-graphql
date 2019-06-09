package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type creatorResolver struct{ *Server }

func (r *creatorResolver) Comics(ctx context.Context, obj *maco.Creator) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Comics)
}

func (r *creatorResolver) Events(ctx context.Context, obj *maco.Creator) ([]*maco.Event, error) {
	return r.store.EventsByIDs(ctx, obj.Events)
}

func (r *creatorResolver) Series(ctx context.Context, obj *maco.Creator) ([]*maco.Series, error) {
	return r.store.SeriesByIDs(ctx, obj.Series)
}

func (r *creatorResolver) Stories(ctx context.Context, obj *maco.Creator) ([]*maco.Story, error) {
	return r.store.StoriesByIDs(ctx, obj.Stories)
}
