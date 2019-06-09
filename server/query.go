package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type queryResolver struct{ *Server }

func (r *queryResolver) Character(ctx context.Context, id *int) (*maco.Character, error) {
	return r.store.CharacterByID(ctx, *id)
}

func (r *queryResolver) Characters(ctx context.Context) ([]*maco.Character, error) {
	panic("not implemented")
}

func (r *queryResolver) Comic(ctx context.Context, id *int) (*maco.Comic, error) {
	return r.store.ComicByID(ctx, *id)
}

func (r *queryResolver) Comics(ctx context.Context) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *queryResolver) Creator(ctx context.Context, id *int) (*maco.Creator, error) {
	return r.store.CreatorByID(ctx, *id)
}

func (r *queryResolver) Creators(ctx context.Context) ([]*maco.Creator, error) {
	panic("not implemented")
}

func (r *queryResolver) Event(ctx context.Context, id *int) (*maco.Event, error) {
	return r.store.EventByID(ctx, *id)
}

func (r *queryResolver) Events(ctx context.Context) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *queryResolver) Serie(ctx context.Context, id *int) (*maco.Series, error) {
	return r.store.SeriesByID(ctx, *id)
}

func (r *queryResolver) Series(ctx context.Context) ([]*maco.Series, error) {
	panic("not implemented")
}

func (r *queryResolver) Story(ctx context.Context, id *int) (*maco.Story, error) {
	return r.store.StoryByID(ctx, *id)
}

func (r *queryResolver) Stories(ctx context.Context) ([]*maco.Story, error) {
	panic("not implemented")
}
