package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type seriesResolver struct{ *Server }

func (r *seriesResolver) Characters(ctx context.Context, obj *maco.Series) ([]*maco.Character, error) {
	return r.store.CharactersByIDs(ctx, obj.Characters)
}

func (r *seriesResolver) Comics(ctx context.Context, obj *maco.Series) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Comics)
}

func (r *seriesResolver) Creators(ctx context.Context, obj *maco.Series) ([]*maco.Creator, error) {
	return r.store.CreatorsByIDs(ctx, obj.Creators)
}

func (r *seriesResolver) Events(ctx context.Context, obj *maco.Series) ([]*maco.Event, error) {
	return r.store.EventsByIDs(ctx, obj.Events)
}

func (r *seriesResolver) Stories(ctx context.Context, obj *maco.Series) ([]*maco.Story, error) {
	return r.store.StoriesByIDs(ctx, obj.Stories)
}
