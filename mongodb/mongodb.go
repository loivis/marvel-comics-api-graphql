package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"github.com/loivis/marvel-comics-api-data-loader/maco"
)

var defaultTimeout time.Duration = 15 * time.Second

type MongoDB struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func New(uri string, database string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	opts := options.Client().
		SetRetryWrites(true).
		SetWriteConcern(
			writeconcern.New(writeconcern.WMajority()),
		)
	client, err := mongo.Connect(ctx, opts.ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %v", err)
	}

	// set context to nil in order to get real error instead of early timeout
	err = client.Ping(nil, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %v", err)
	}

	return &MongoDB{
		client:   client,
		database: database,
		timeout:  defaultTimeout,
	}, nil
}

func (m *MongoDB) CharacterByID(ctx context.Context, id int) (*maco.Character, error) {
	col := m.client.Database(m.database).Collection(maco.TypeCharacters)

	var out maco.Character
	err := col.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (m *MongoDB) Characters(ctx context.Context) ([]*maco.Character, error) {
	col := m.client.Database(m.database).Collection(maco.TypeCharacters)

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return charactersFromCursor(ctx, cur)
}

func (m *MongoDB) CharactersByIDs(ctx context.Context, ids []int) ([]*maco.Character, error) {
	cur, err := m.cursorFilteredByIDs(ctx, maco.TypeCharacters, ids)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return charactersFromCursor(ctx, cur)
}

func charactersFromCursor(ctx context.Context, cur *mongo.Cursor) ([]*maco.Character, error) {
	var out []*maco.Character

	for cur.Next(ctx) {
		var elem maco.Character
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		out = append(out, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (m *MongoDB) ComicByID(ctx context.Context, id int) (*maco.Comic, error) {
	col := m.client.Database(m.database).Collection(maco.TypeComics)

	var out maco.Comic
	err := col.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (m *MongoDB) Comics(ctx context.Context) ([]*maco.Comic, error) {
	col := m.client.Database(m.database).Collection(maco.TypeComics)

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return comicsFromCursor(ctx, cur)
}

func (m *MongoDB) ComicsByIDs(ctx context.Context, ids []int) ([]*maco.Comic, error) {
	cur, err := m.cursorFilteredByIDs(ctx, maco.TypeComics, ids)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return comicsFromCursor(ctx, cur)
}

func comicsFromCursor(ctx context.Context, cur *mongo.Cursor) ([]*maco.Comic, error) {
	var out []*maco.Comic

	for cur.Next(ctx) {
		var elem maco.Comic
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		out = append(out, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (m *MongoDB) CreatorByID(ctx context.Context, id int) (*maco.Creator, error) {
	col := m.client.Database(m.database).Collection(maco.TypeCreators)

	var out maco.Creator
	err := col.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (m *MongoDB) Creators(ctx context.Context) ([]*maco.Creator, error) {
	col := m.client.Database(m.database).Collection(maco.TypeCreators)

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return creatorsFromCursor(ctx, cur)
}

func (m *MongoDB) CreatorsByIDs(ctx context.Context, ids []int) ([]*maco.Creator, error) {
	cur, err := m.cursorFilteredByIDs(ctx, maco.TypeCreators, ids)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return creatorsFromCursor(ctx, cur)
}

func creatorsFromCursor(ctx context.Context, cur *mongo.Cursor) ([]*maco.Creator, error) {
	var out []*maco.Creator

	for cur.Next(ctx) {
		var elem maco.Creator
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		out = append(out, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (m *MongoDB) EventByID(ctx context.Context, id int) (*maco.Event, error) {
	col := m.client.Database(m.database).Collection(maco.TypeEvents)

	var out maco.Event
	err := col.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (m *MongoDB) Events(ctx context.Context) ([]*maco.Event, error) {
	col := m.client.Database(m.database).Collection(maco.TypeEvents)

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return eventsFromCursor(ctx, cur)
}

func (m *MongoDB) EventsByIDs(ctx context.Context, ids []int) ([]*maco.Event, error) {
	cur, err := m.cursorFilteredByIDs(ctx, maco.TypeEvents, ids)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return eventsFromCursor(ctx, cur)
}

func eventsFromCursor(ctx context.Context, cur *mongo.Cursor) ([]*maco.Event, error) {
	var out []*maco.Event

	for cur.Next(ctx) {
		var elem maco.Event
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		out = append(out, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (m *MongoDB) SeriesByID(ctx context.Context, id int) (*maco.Series, error) {
	col := m.client.Database(m.database).Collection(maco.TypeSeries)

	var out maco.Series
	err := col.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (m *MongoDB) Series(ctx context.Context) ([]*maco.Series, error) {
	col := m.client.Database(m.database).Collection(maco.TypeSeries)

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return seriesFromCursor(ctx, cur)
}

func (m *MongoDB) SeriesByIDs(ctx context.Context, ids []int) ([]*maco.Series, error) {
	cur, err := m.cursorFilteredByIDs(ctx, maco.TypeSeries, ids)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return seriesFromCursor(ctx, cur)
}

func seriesFromCursor(ctx context.Context, cur *mongo.Cursor) ([]*maco.Series, error) {
	var out []*maco.Series

	for cur.Next(ctx) {
		var elem maco.Series
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		out = append(out, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (m *MongoDB) StoryByID(ctx context.Context, id int) (*maco.Story, error) {
	col := m.client.Database(m.database).Collection(maco.TypeStories)

	var out maco.Story
	err := col.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (m *MongoDB) Stories(ctx context.Context) ([]*maco.Story, error) {
	col := m.client.Database(m.database).Collection(maco.TypeStories)

	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return storiesFromCursor(ctx, cur)
}

func (m *MongoDB) StoriesByIDs(ctx context.Context, ids []int) ([]*maco.Story, error) {
	cur, err := m.cursorFilteredByIDs(ctx, maco.TypeStories, ids)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	return storiesFromCursor(ctx, cur)
}

func storiesFromCursor(ctx context.Context, cur *mongo.Cursor) ([]*maco.Story, error) {
	var out []*maco.Story

	for cur.Next(ctx) {
		var elem maco.Story
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		out = append(out, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (m *MongoDB) cursorFilteredByIDs(ctx context.Context, collection string, ids []int) (*mongo.Cursor, error) {
	col := m.client.Database(m.database).Collection(collection)

	filter := bson.D{{
		Key: "id",
		Value: bson.D{{
			Key:   "$in",
			Value: ids,
		}},
	}}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return cur, nil
}
