package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type storyResolver struct{ *Server }

func (r *storyResolver) Characters(ctx context.Context, obj *maco.Story) ([]*maco.Character, error) {
	return r.store.CharactersByIDs(ctx, obj.Characters)
}

func (r *storyResolver) Comics(ctx context.Context, obj *maco.Story) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Comics)
}

func (r *storyResolver) Creators(ctx context.Context, obj *maco.Story) ([]*maco.Creator, error) {
	return r.store.CreatorsByIDs(ctx, obj.Creators)
}

func (r *storyResolver) Events(ctx context.Context, obj *maco.Story) ([]*maco.Event, error) {
	return r.store.EventsByIDs(ctx, obj.Events)
}

func (r *storyResolver) OriginalIssue(ctx context.Context, obj *maco.Story) (*maco.Comic, error) {
	return r.store.ComicByID(ctx, obj.OriginalIssue)
}

func (r *storyResolver) Series(ctx context.Context, obj *maco.Story) ([]*maco.Series, error) {
	return r.store.SeriesByIDs(ctx, obj.Series)
}
