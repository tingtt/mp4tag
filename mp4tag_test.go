package mp4tag

import (
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases []testCase = []testCase{{
	path: "Cartoon, Jéja - On & On (feat. Daniel Levi) [NCS Release].m4a",
	tag: MP4Tags{
		Title:       "On & On",
		Artist:      "Cartoon, Jéja, Daniel Levi",
		Album:       "On & On",
		AlbumArtist: "Cartoon, Jéja, Daniel Levi",
		Year:        2015,
		TrackNumber: 1,
	},
}, {
	path: "Ghostnaps - Grow Apart [NCS Release].m4a",
	tag: MP4Tags{
		Title:       "grow apart",
		Artist:      "Ghostnaps",
		Album:       "grow apart",
		AlbumArtist: "Ghostnaps",
		Year:        2024,
		TrackNumber: 1,
	},
}, {
	path: "Spicyverse - Vibe [NCS Release].m4a",
	tag: MP4Tags{
		Title:       "Vibe",
		Artist:      "Spicyverse",
		Album:       "Vibe",
		AlbumArtist: "Spicyverse",
		Year:        2024,
		TrackNumber: 1,
	},
}}

type testCase struct {
	path string
	tag  MP4Tags
}

type argsRead struct {
	f *os.File
}

type testCallRead struct {
	name string
	args argsRead
	want MP4Tags
}

func TestReader_Read(t *testing.T) {
	t.Parallel()

	tests := Map(cases, func(c testCase) testCallRead {
		return testCallRead{
			name: c.path,
			args: argsRead{
				f: func() *os.File {
					file, _ := os.Open(path.Join("./test/data/", c.path))
					return file
				}(),
			},
			want: MP4Tags{
				Title:       c.tag.Title,
				Artist:      c.tag.Artist,
				Album:       c.tag.Album,
				AlbumArtist: c.tag.AlbumArtist,
				Year:        c.tag.Year,
				TrackNumber: c.tag.TrackNumber,
			},
		}
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reader(tt.args.f)
			assert.NoError(t, err)

			tag, err := got.Read()
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Title, tag.Title)
			assert.Equal(t, tt.want.Artist, tag.Artist)
			assert.Equal(t, tt.want.Album, tag.Album)
			assert.Equal(t, tt.want.AlbumArtist, tag.AlbumArtist)
			assert.Equal(t, tt.want.Year, tag.Year)
			assert.Equal(t, tt.want.TrackNumber, tag.TrackNumber)
		})
	}
}

type argsWrite struct {
	tag MP4Tags
}
type testCallWrite struct {
	name string
	args struct {
		argsRead
		argsWrite
	}
}

func copySampleAudioFile(filename string) *os.File {
	src, err := os.Open(path.Join("./test/data/", filename))
	if err != nil {
		panic(err)
	}
	dst, err := os.Create(path.Join("./test/data/tmp/", filename))
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := dst.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	copied, err := os.Open(path.Join("./test/data/tmp/", filename))
	if err != nil {
		panic(err)
	}
	return copied
}

func clearCopiedSampleAudioFiles() {
	err := os.RemoveAll("./test/data/tmp/")
	if err != nil {
		panic(err)
	}
}

func TestParseReadWriter_Write(t *testing.T) {
	t.Parallel()

	if err := os.MkdirAll("./test/data/tmp/", 0755); err != nil {
		panic(err)
	}
	defer clearCopiedSampleAudioFiles()

	tests := Map(cases, func(c testCase) testCallWrite {
		return testCallWrite{
			name: c.path,
			args: struct {
				argsRead
				argsWrite
			}{
				argsRead: argsRead{
					f: copySampleAudioFile(c.path),
				},
				argsWrite: argsWrite{
					tag: MP4Tags{
						Title:       fmt.Sprintf("%s-modified", c.tag.Title),
						Artist:      fmt.Sprintf("%s-modified", c.tag.Artist),
						Album:       fmt.Sprintf("%s-modified", c.tag.Album),
						AlbumArtist: fmt.Sprintf("%s-modified", c.tag.AlbumArtist),
						Year:        c.tag.Year + 1,
						TrackNumber: c.tag.TrackNumber + 1,
					},
				},
			},
		}
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseReadWriter(tt.args.f)
			assert.NoError(t, err)

			err = got.Write(&tt.args.tag, nil)
			assert.NoError(t, err)

			tag, err := got.Read()
			assert.NoError(t, err)
			assert.Equal(t, tt.args.tag.Title, tag.Title)
			assert.Equal(t, tt.args.tag.Artist, tag.Artist)
			assert.Equal(t, tt.args.tag.Album, tag.Album)
			assert.Equal(t, tt.args.tag.AlbumArtist, tag.AlbumArtist)
			assert.Equal(t, tt.args.tag.Year, tag.Year)
			assert.Equal(t, tt.args.tag.TrackNumber, tag.TrackNumber)
		})
	}
}

func Map[T1, T2 any](list []T1, yield func(T1) T2) []T2 {
	newList := make([]T2, 0, len(list))
	for _, t := range list {
		newList = append(newList, yield(t))
	}
	return newList
}
