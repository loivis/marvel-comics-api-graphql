package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

type seriesResolver struct{ *Server }

func (r *seriesResolver) Characters(ctx context.Context, obj *maco.Series, first *int, after *int) (*macogql.CharactersResult, error) {
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

func (r *seriesResolver) Comics(ctx context.Context, obj *maco.Series, first *int, after *int) (*macogql.ComicsResult, error) {
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

func (r *seriesResolver) Creators(ctx context.Context, obj *maco.Series, first *int, after *int) (*macogql.CreatorsResult, error) {
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

func (r *seriesResolver) Events(ctx context.Context, obj *maco.Series, first *int, after *int) (*macogql.EventsResult, error) {
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
		Events: res[from:to],
		PageInfo: &macogql.PageInfo{
			End:     to,
			HasNext: to < length,
		},
		TotalCount: length,
	}, nil
}

func (r *seriesResolver) Stories(ctx context.Context, obj *maco.Series, first *int, after *int) (*macogql.StoriesResult, error) {
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
