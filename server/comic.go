package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

type comicResolver struct{ *Server }

func (r *comicResolver) Characters(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.CharactersResult, error) {
	if len(obj.Characters) == 0 {
		return nil, nil
	}

	res, err := r.store.CharactersByIDs(ctx, obj.Characters)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.CharactersResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *comicResolver) CollectedIssues(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.ComicsResult, error) {
	if len(obj.CollectedIssues) == 0 {
		return nil, nil
	}

	res, err := r.store.ComicsByIDs(ctx, obj.CollectedIssues)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.ComicsResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *comicResolver) Collections(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.ComicsResult, error) {
	if len(obj.Collections) == 0 {
		return nil, nil
	}

	res, err := r.store.ComicsByIDs(ctx, obj.Collections)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.ComicsResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *comicResolver) Creators(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.CreatorsResult, error) {
	if len(obj.Creators) == 0 {
		return nil, nil
	}

	res, err := r.store.CreatorsByIDs(ctx, obj.Creators)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.CreatorsResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *comicResolver) Events(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.EventsResult, error) {
	if len(obj.Events) == 0 {
		return nil, nil
	}

	res, err := r.store.EventsByIDs(ctx, obj.Events)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.EventsResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *comicResolver) Stories(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.StoriesResult, error) {
	if len(obj.Stories) == 0 {
		return nil, nil
	}

	res, err := r.store.StoriesByIDs(ctx, obj.Stories)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.StoriesResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *comicResolver) Variants(ctx context.Context, obj *maco.Comic, first *int, after *int) (*macogql.ComicsResult, error) {
	res, err := r.store.ComicsByIDs(ctx, obj.Variants)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.ComicsResult{
		Items: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

type comicPriceResolver struct{ *Server }

func (r *comicPriceResolver) Price(ctx context.Context, obj *maco.ComicPrice) (float64, error) {
	return float64(obj.Price), nil
}
