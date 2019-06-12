package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

type characterResolver struct{ *Server }

func (r *characterResolver) Comics(ctx context.Context, obj *maco.Character, first *int, after *int) (*macogql.ComicsResult, error) {
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

func (r *characterResolver) Events(ctx context.Context, obj *maco.Character, first *int, after *int) (*macogql.EventsResult, error) {
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

func (r *characterResolver) Series(ctx context.Context, obj *maco.Character, first *int, after *int) (*macogql.SeriesResult, error) {
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

func (r *characterResolver) Stories(ctx context.Context, obj *maco.Character, first *int, after *int) (*macogql.StoriesResult, error) {
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
