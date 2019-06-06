package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type storyResolver struct{ *Server }

func (r *storyResolver) Characters(ctx context.Context, obj *maco.Story) ([]*maco.Character, error) {
	panic("not implemented")
}

func (r *storyResolver) Comics(ctx context.Context, obj *maco.Story) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *storyResolver) Creators(ctx context.Context, obj *maco.Story) ([]*maco.Creator, error) {
	panic("not implemented")
}

func (r *storyResolver) Events(ctx context.Context, obj *maco.Story) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *storyResolver) OriginalIssue(ctx context.Context, obj *maco.Story) (*maco.Comic, error) {
	panic("not implemented")
}

func (r *storyResolver) Series(ctx context.Context, obj *maco.Story) ([]*maco.Series, error) {
	panic("not implemented")
}
