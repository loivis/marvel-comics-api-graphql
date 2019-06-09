package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type comicResolver struct{ *Server }

func (r *comicResolver) Characters(ctx context.Context, obj *maco.Comic) ([]*maco.Character, error) {
	return r.store.CharactersByIDs(ctx, obj.Characters)
}

func (r *comicResolver) CollectedIssues(ctx context.Context, obj *maco.Comic) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.CollectedIssues)
}

func (r *comicResolver) Collections(ctx context.Context, obj *maco.Comic) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Collections)
}

func (r *comicResolver) Creators(ctx context.Context, obj *maco.Comic) ([]*maco.Creator, error) {
	return r.store.CreatorsByIDs(ctx, obj.Creators)
}

func (r *comicResolver) Events(ctx context.Context, obj *maco.Comic) ([]*maco.Event, error) {
	return r.store.EventsByIDs(ctx, obj.Events)
}

func (r *comicResolver) Stories(ctx context.Context, obj *maco.Comic) ([]*maco.Story, error) {
	return r.store.StoriesByIDs(ctx, obj.Stories)
}

func (r *comicResolver) Variants(ctx context.Context, obj *maco.Comic) ([]*maco.Comic, error) {
	return r.store.ComicsByIDs(ctx, obj.Variants)
}

type comicPriceResolver struct{ *Server }

func (r *comicPriceResolver) Price(ctx context.Context, obj *maco.ComicPrice) (float64, error) {
	return float64(obj.Price), nil
}
