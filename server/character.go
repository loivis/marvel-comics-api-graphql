package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type characterResolver struct{ *Server }

func (r *characterResolver) Comics(ctx context.Context, obj *maco.Character) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Comics)
}

func (r *characterResolver) Events(ctx context.Context, obj *maco.Character) ([]*maco.Event, error) {
	return r.store.EventsByIDs(ctx, obj.Events)
}

func (r *characterResolver) Series(ctx context.Context, obj *maco.Character) ([]*maco.Series, error) {
	return r.store.SeriesByIDs(ctx, obj.Series)
}

func (r *characterResolver) Stories(ctx context.Context, obj *maco.Character) ([]*maco.Story, error) {
	return r.store.StoriesByIDs(ctx, obj.Stories)
}
