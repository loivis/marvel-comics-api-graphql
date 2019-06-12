package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

type eventResolver struct{ *Server }

func (r *eventResolver) Characters(ctx context.Context, obj *maco.Event, first *int, after *int) (*macogql.CharactersResult, error) {
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
		Characters: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *eventResolver) Comics(ctx context.Context, obj *maco.Event, first *int, after *int) (*macogql.ComicsResult, error) {
	if len(obj.Comics) == 0 {
		return nil, nil
	}

	res, err := r.store.ComicsByIDs(ctx, obj.Comics)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.ComicsResult{
		Comics: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *eventResolver) Creators(ctx context.Context, obj *maco.Event, first *int, after *int) (*macogql.CreatorsResult, error) {
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
		Creators: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *eventResolver) Series(ctx context.Context, obj *maco.Event, first *int, after *int) (*macogql.SeriesResult, error) {
	if len(obj.Series) == 0 {
		return nil, nil
	}

	res, err := r.store.SeriesByIDs(ctx, obj.Series)
	if err != nil {
		return nil, err
	}

	length := len(res)
	if length == 0 {
		return nil, nil
	}

	from, to := fromTo(*after, *after+*first, length)

	return &macogql.SeriesResult{
		Series: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *eventResolver) Stories(ctx context.Context, obj *maco.Event, first *int, after *int) (*macogql.StoriesResult, error) {
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
		Stories: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}
