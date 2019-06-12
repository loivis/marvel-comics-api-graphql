package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

type storyResolver struct{ *Server }

func (r *storyResolver) Characters(ctx context.Context, obj *maco.Story, first *int, after *int) (*macogql.CharactersResult, error) {
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

func (r *storyResolver) Comics(ctx context.Context, obj *maco.Story, first *int, after *int) (*macogql.ComicsResult, error) {
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

func (r *storyResolver) Creators(ctx context.Context, obj *maco.Story, first *int, after *int) (*macogql.CreatorsResult, error) {
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

func (r *storyResolver) Events(ctx context.Context, obj *maco.Story, first *int, after *int) (*macogql.EventsResult, error) {
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

func (r *storyResolver) OriginalIssue(ctx context.Context, obj *maco.Story) (*maco.Comic, error) {
	return r.store.ComicByID(ctx, obj.OriginalIssue)
}

func (r *storyResolver) Series(ctx context.Context, obj *maco.Story, first *int, after *int) (*macogql.SeriesResult, error) {
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
