package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type eventResolver struct{ *Server }

func (r *eventResolver) Characters(ctx context.Context, obj *maco.Event) ([]*maco.Character, error) {
	return r.store.CharactersByIDs(ctx, obj.Characters)
}

func (r *eventResolver) Comics(ctx context.Context, obj *maco.Event) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Comics)
}

func (r *eventResolver) Creators(ctx context.Context, obj *maco.Event) ([]*maco.Creator, error) {
	return r.store.CreatorsByIDs(ctx, obj.Creators)
}

func (r *eventResolver) Series(ctx context.Context, obj *maco.Event) ([]*maco.Series, error) {
	return r.store.SeriesByIDs(ctx, obj.Series)
}

func (r *eventResolver) Stories(ctx context.Context, obj *maco.Event) ([]*maco.Story, error) {
	return r.store.StoriesByIDs(ctx, obj.Stories)
}
