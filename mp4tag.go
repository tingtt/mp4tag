package mp4tag

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"

	"gitlab.com/osaki-lab/iowrapper"
)

type MP4TagReader interface {
	Read() (*MP4Tags, error)
	UpperCustom(b bool)
}

type MP4TagWriter interface {
	Close() error
	Write(tags *MP4Tags, delStrings []string) error
	UpperCustom(b bool)
}

type MP4TagReadWriter interface {
	MP4TagReader
	MP4TagWriter
}

func (mp4 *MP4R) UpperCustom(b bool) {
	mp4.upperCustom = b
}

func (mp4 *MP4RW) Close() error {
	return mp4.f.Close()
}

func (mp4 *MP4R) Read() (*MP4Tags, error) {
	tags, _, err := mp4.actualRead()
	return tags, err
}

func (mp4 *MP4RW) Write(tags *MP4Tags, delStrings []string) error {
	if tags == nil && len(delStrings) == 0 {
		return nil
	}
	err := mp4.actualWrite(tags, delStrings)
	return err
}

func (mp4 *MP4R) checkHeader() error {
	_, err := mp4.f.Seek(4, io.SeekStart)
	if err != nil {
		return err
	}
	buf := make([]byte, 8)
	_, err = io.ReadFull(mp4.f, buf)
	if err != nil {
		return err
	}

	if !bytes.Equal(buf[:4], []byte{0x66, 0x74, 0x79, 0x70}) {
		return &ErrInvalidMagic{}
	}
	for _, ftyp := range ftyps {
		if bytes.Equal(buf[4:], ftyp) {
			return nil
		}
	}
	return &ErrUnsupportedFtyp{
		Msg: "unsupported ftyp: " + fmt.Sprintf("%x", buf[4:]),
	}
}

func Open(trackPath string) (MP4TagReadWriter, error) {
	return open(trackPath)
}

func open(trackPath string) (*MP4RW, error) {
	f, err := os.Open(trackPath)
	if err != nil {
		return nil, err
	}
	return parse(f)
}

func ParseReadWriter(f *os.File) (MP4TagReadWriter, error) {
	return parse(f)
}

func parse(f *os.File) (*MP4RW, error) {
	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	mp4 := &MP4RW{
		MP4R: MP4R{
			f:           f,
			size:        stat.Size(),
			upperCustom: true,
		},
		f: f,
	}
	err = mp4.checkHeader()
	if err != nil {
		f.Close()
		return nil, err
	}
	return mp4, nil
}

func Reader(f fs.File) (MP4TagReader, error) {
	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	mp4 := &MP4R{
		f:           iowrapper.NewSeeker(f, iowrapper.MaxBufferSize(uint64(stat.Size()))),
		size:        stat.Size(),
		upperCustom: true,
	}
	err = mp4.checkHeader()
	if err != nil {
		f.Close()
		return nil, err
	}
	return mp4, nil
}
