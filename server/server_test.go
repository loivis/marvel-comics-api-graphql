package server

import (
	"fmt"
	"testing"
)

func TestFromTo(t *testing.T) {
	type ft struct{ from, to int }

	for i, tc := range []struct {
		offset int
		length int
		max    int
		out    ft
	}{
		{
			offset: 3,
			length: 7,
			max:    11,
			out:    ft{3, 10},
		},
		{
			offset: 3,
			length: 7,
			max:    7,
			out:    ft{3, 7},
		},
		{
			offset: 3,
			length: 7,
			max:    2,
			out:    ft{2, 2},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			from, to := fromTo(tc.offset, tc.length, tc.max)

			if got, want := (ft{from, to}), tc.out; got != want {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}
