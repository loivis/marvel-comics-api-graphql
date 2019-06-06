package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type comicResolver struct{ *Server }

func (r *comicResolver) Characters(ctx context.Context, obj *maco.Comic) ([]*maco.Character, error) {
	panic("not implemented")
}

func (r *comicResolver) CollectedIssues(ctx context.Context, obj *maco.Comic) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *comicResolver) Collections(ctx context.Context, obj *maco.Comic) ([]*maco.Comic, error) {
	panic("not implemented")
}

func (r *comicResolver) Creators(ctx context.Context, obj *maco.Comic) ([]*maco.Creator, error) {
	panic("not implemented")
}

func (r *comicResolver) Events(ctx context.Context, obj *maco.Comic) ([]*maco.Event, error) {
	panic("not implemented")
}

func (r *comicResolver) Stories(ctx context.Context, obj *maco.Comic) ([]*maco.Story, error) {
	panic("not implemented")
}

func (r *comicResolver) Variants(ctx context.Context, obj *maco.Comic) ([]*maco.Comic, error) {
	panic("not implemented")
}

type comicPriceResolver struct{ *Server }

func (r *comicPriceResolver) Price(ctx context.Context, obj *maco.ComicPrice) (float64, error) {
	panic("not implemented")
}
