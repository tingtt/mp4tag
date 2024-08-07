package mp4tag

import (
	"io"
	"os"
)

type ErrBoxNotPresent struct {
	Msg string
}

type ErrUnsupportedFtyp struct {
	Msg string
}

type ErrInvalidStcoSize struct{}

type ErrInvalidMagic struct{}

func (e *ErrBoxNotPresent) Error() string {
	return e.Msg
}

func (e *ErrUnsupportedFtyp) Error() string {
	return e.Msg
}

func (e *ErrInvalidStcoSize) Error() string {
	return "stco size is invalid"
}

func (e *ErrInvalidMagic) Error() string {
	return "file header is corrupted or not an mp4 file"
}

var ftyps = [8][]byte{
	{0x4D, 0x34, 0x41, 0x20}, // M4A
	{0x4D, 0x34, 0x42, 0x20}, // M4B
	{0x64, 0x61, 0x73, 0x68}, // dash
	{0x6D, 0x70, 0x34, 0x31}, // mp41
	{0x6D, 0x70, 0x34, 0x32}, // mp42
	{0x69, 0x73, 0x6F, 0x6D}, // isom
	{0x69, 0x73, 0x6F, 0x32}, // iso2
	{0x61, 0x76, 0x63, 0x31}, // avc1
}

var containers = []string{
	"moov", "udta", "meta", "ilst", "----", "(c)alb",
	"aART", "(c)art", "(c)nam", "(c)cmt", "(c)gen", "gnre",
	"(c)wrt", "(c)con", "cprt", "desc", "(c)lyr", "(c)nrt",
	"(c)pub", "trkn", "covr", "(c)day", "disk", "(c)too",
	"trak", "mdia", "minf", "stbl", "rtng", "plID",
	"atID", "tmpo", "sonm", "soal", "soar", "soco",
	"soaa",
}

// 0-9
var numbers = []rune{
	0x30, 0x31, 0x32, 0x33, 0x34,
	0x35, 0x36, 0x37, 0x38, 0x39,
}

type MP4RW struct {
	MP4R
	f *os.File
}
type MP4R struct {
	f           io.ReadSeeker
	size        int64
	upperCustom bool
}

type MP4Box struct {
	StartOffset int64
	EndOffset   int64
	BoxSize     int64
	Path        string
}

type MP4Boxes struct {
	Boxes []*MP4Box
}

type ImageType int8

const (
	ImageTypeJPEG ImageType = iota + 13
	ImageTypePNG
	ImageTypeAuto
)

var resolveImageType = map[uint8]ImageType{
	13: ImageTypeJPEG,
	14: ImageTypePNG,
}

type ItunesAdvisory int8

const (
	ItunesAdvisoryNone ItunesAdvisory = iota
	ItunesAdvisoryExplicit
	ItunesAdvisoryClean
)

var resolveItunesAdvisory = map[uint8]ItunesAdvisory{
	1: ItunesAdvisoryExplicit,
	2: ItunesAdvisoryClean,
}

type MP4Picture struct {
	Format ImageType
	Data   []byte
}

type MP4Tags struct {
	Album           string
	AlbumSort       string
	AlbumArtist     string
	AlbumArtistSort string
	Artist          string
	ArtistSort      string
	BPM             int16
	Comment         string
	Composer        string
	ComposerSort    string
	Conductor       string
	Copyright       string
	Custom          map[string]string
	CustomGenre     string
	Date            string
	Description     string
	Director        string
	DiscNumber      int16
	DiscTotal       int16
	Genre           Genre
	ItunesAdvisory  ItunesAdvisory
	ItunesAlbumID   int32
	ItunesArtistID  int32
	Lyrics          string
	Narrator        string
	OtherCustom     map[string][]string
	Pictures        []*MP4Picture
	Publisher       string
	Title           string
	TitleSort       string
	TrackNumber     int16
	TrackTotal      int16
	Year            int32
}
