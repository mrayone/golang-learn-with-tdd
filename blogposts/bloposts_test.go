package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/mrayone/learn-go/blogposts"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	testCases := []struct {
		desc        string
		fs          fs.FS
		want        []blogposts.Post
		expectedErr string
	}{
		{
			desc: "open valid fs",
			want: []blogposts.Post{
				{
					Title:       "Post 1",
					Description: "Description 1",
					Tags:        []string{"tdd", "go"},
					Body: `Hello
World`,
				},
				{
					Title:       "Post 2",
					Description: "Description 2",
					Tags:        []string{"rust", "borrow-checker"},
					Body: `B
L
M`,
				},
			},
			fs: fstest.MapFS{
				"hello world.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`)},
				"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`)},
			},
		},
		{
			desc:        "open invalid fs",
			fs:          StubFailingFS{},
			expectedErr: "oh no, i always fail",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := blogposts.NewPostsFromFS(tc.fs)

			assertError(t, err, tc.expectedErr)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %+v posts, wanted %+v posts", got, tc.want)
			}
		})
	}
}

func assertError(t testing.TB, err error, expectedErr string) {
	t.Helper()
	message := ""
	if err != nil {
		message = err.Error()
	}

	if message != expectedErr {
		t.Errorf("got %s error, wanted %s error", message, expectedErr)
	}
}
