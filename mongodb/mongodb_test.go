package mongodb

import (
	"context"
	"fmt"
	"testing"
)

func TestMongoDB_CharacterByID(t *testing.T) {
	m, _, err := setupDatabase("marvel_test", "characters")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer m.client.Database("marvel_test").Drop(context.Background())

	char, err := m.CharacterByID(context.Background(), 3)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	if got, want := char.ID, 3; got != want {
		t.Errorf("got character ID %d, want %d", got, want)
	}

	if got, want := char.Name, "foo-3"; got != want {
		t.Errorf("got character name %q, want %q", got, want)
	}
}

func TestMongoDB_Characters(t *testing.T) {
	m, _, err := setupDatabase("marvel_test", "characters")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer m.client.Database("marvel_test").Drop(context.Background())

	chars, err := m.Characters(context.Background())
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	wantChars := []struct {
		ID   int
		Name string
	}{
		{ID: 0, Name: "foo-0"},
		{ID: 1, Name: "foo-1"},
		{ID: 2, Name: "foo-2"},
		{ID: 3, Name: "foo-3"},
		{ID: 4, Name: "foo-4"},
		{ID: 5, Name: "foo-5"},
		{ID: 6, Name: "foo-6"},
	}

	if got, want := len(chars), len(wantChars); got != want {
		t.Fatalf("got %d characters, want %d", got, want)
	}

	for i, char := range wantChars {
		if got, want := chars[i].ID, char.ID; got != want {
			t.Errorf("got character[%d].ID %d, want %d", i, got, want)
		}

		if got, want := chars[i].Name, char.Name; got != want {
			t.Errorf("got character[%d].name %q, want %q", i, got, want)
		}
	}

}

func TestMongoDB_CharactersByIDs(t *testing.T) {
	m, _, err := setupDatabase("marvel_test", "characters")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer m.client.Database("marvel_test").Drop(context.Background())

	chars, err := m.CharactersByIDs(context.Background(), []int{1, 2, 4})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	wantChars := []struct {
		ID   int
		Name string
	}{
		{
			ID:   1,
			Name: "foo-1",
		},
		{
			ID:   2,
			Name: "foo-2",
		},
		{
			ID:   4,
			Name: "foo-4",
		},
	}

	if got, want := len(chars), len(wantChars); got != want {
		t.Fatalf("got %d characters, want %d", got, want)
	}

	for i, char := range wantChars {
		if got, want := chars[i].ID, char.ID; got != want {
			t.Errorf("got character[%d].ID %d, want %d", i, got, want)
		}

		if got, want := chars[i].Name, char.Name; got != want {
			t.Errorf("got character[%d].name %q, want %q", i, got, want)
		}
	}

}

func setupDatabase(database, collection string) (*MongoDB, []interface{}, error) {
	m, err := New("mongodb://localhost:27017", database)
	if err != nil {
		return nil, nil, err
	}

	if err := m.client.Database(database).Drop(context.Background()); err != nil {
		return nil, nil, err
	}

	docs := []interface{}{}
	for i := 0; i < 7; i++ {

		docs = append(docs, struct {
			ID   int
			Name string
		}{
			ID:   i,
			Name: fmt.Sprintf("foo-%d", i),
		})
	}

	m.client.Database(database).Collection(collection).InsertMany(
		context.Background(),
		docs,
	)

	return m, docs, nil
}
