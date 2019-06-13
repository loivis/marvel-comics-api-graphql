package server

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
	"github.com/loivis/marvel-comics-api-graphql/macogql"
)

type queryResolver struct{ *Server }

func (r *queryResolver) Character(ctx context.Context, id int) (*maco.Character, error) {
	return r.store.CharacterByID(ctx, id)
}

func (r *queryResolver) Characters(ctx context.Context, first *int, after *int) (*macogql.CharactersResult, error) {
	res, err := r.store.Characters(ctx)
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

func (r *queryResolver) Comic(ctx context.Context, id int) (*maco.Comic, error) {
	return r.store.ComicByID(ctx, id)
}

func (r *queryResolver) Comics(ctx context.Context, first *int, after *int) (*macogql.ComicsResult, error) {
	res, err := r.store.Comics(ctx)
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

func (r *queryResolver) Creator(ctx context.Context, id int) (*maco.Creator, error) {
	return r.store.CreatorByID(ctx, id)
}

func (r *queryResolver) Creators(ctx context.Context, first *int, after *int) (*macogql.CreatorsResult, error) {
	res, err := r.store.Creators(ctx)
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

func (r *queryResolver) Event(ctx context.Context, id int) (*maco.Event, error) {
	return r.store.EventByID(ctx, id)
}

func (r *queryResolver) Events(ctx context.Context, first *int, after *int) (*macogql.EventsResult, error) {
	res, err := r.store.Events(ctx)
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

func (r *queryResolver) Serie(ctx context.Context, id int) (*maco.Series, error) {
	return r.store.SeriesByID(ctx, id)
}

func (r *queryResolver) Series(ctx context.Context, first *int, after *int) (*macogql.SeriesResult, error) {
	res, err := r.store.Series(ctx)
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

func (r *queryResolver) Story(ctx context.Context, id int) (*maco.Story, error) {
	return r.store.StoryByID(ctx, id)
}

func (r *queryResolver) Stories(ctx context.Context, first *int, after *int) (*macogql.StoriesResult, error) {
	res, err := r.store.Stories(ctx)
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
