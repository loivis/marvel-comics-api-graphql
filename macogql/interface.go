package macogql

import (
	"context"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

type Store interface {
	CharacterByID(ctx context.Context, id int) (*maco.Character, error)
	Characters(ctx context.Context) ([]*maco.Character, error)
	CharactersByIDs(ctx context.Context, ids []int) ([]*maco.Character, error)
	ComicByID(ctx context.Context, id int) (*maco.Comic, error)
	Comics(ctx context.Context) ([]*maco.Comic, error)
	ComicsByIDs(ctx context.Context, ids []int) ([]*maco.Comic, error)
	CreatorByID(ctx context.Context, id int) (*maco.Creator, error)
	Creators(ctx context.Context) ([]*maco.Creator, error)
	CreatorsByIDs(ctx context.Context, ids []int) ([]*maco.Creator, error)
	EventByID(ctx context.Context, id int) (*maco.Event, error)
	Events(ctx context.Context) ([]*maco.Event, error)
	EventsByIDs(ctx context.Context, ids []int) ([]*maco.Event, error)
	SeriesByID(ctx context.Context, id int) (*maco.Series, error)
	Series(ctx context.Context) ([]*maco.Series, error)
	SeriesByIDs(ctx context.Context, ids []int) ([]*maco.Series, error)
	StoryByID(ctx context.Context, id int) (*maco.Story, error)
	Stories(ctx context.Context) ([]*maco.Story, error)
	StoriesByIDs(ctx context.Context, ids []int) ([]*maco.Story, error)
}
