package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type seriesResolver struct{ *Server }

func (r *seriesResolver) Characters(ctx context.Context, obj *maco.Series) ([]*maco.Character, error) {
	panic("not implemented")
}

func (r *seriesResolver) Comics(ctx context.Context, obj *maco.Series) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *seriesResolver) Creators(ctx context.Context, obj *maco.Series) ([]*maco.Creator, error) {
	panic("not implemented")
}

func (r *seriesResolver) Events(ctx context.Context, obj *maco.Series) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *seriesResolver) Stories(ctx context.Context, obj *maco.Series) ([]*maco.Story, error) {
	panic("not implemented")
}
