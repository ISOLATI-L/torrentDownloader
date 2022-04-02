package Bencode

import (
	"bufio"
	"errors"
	"io"
)

var ErrTyp error
var ErrFmt error

func init() {
	ErrTyp = errors.New("error type")
	ErrFmt = errors.New("error format")
}

// type Btype uint8

// const (
// 	BSTR  Btype = 0x01
// 	BINT  Btype = 0x02
// 	BLIST Btype = 0x03
// 	BDICT Btype = 0x04
// )

type Bobject interface {
	Bencode(w io.Writer) int
}

func Parse(r io.Reader) (Bobject, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	b, err := br.Peek(1)
	if err != nil {
		return nil, err
	}

	var ret Bobject
	switch {
	case b[0] >= '0' && b[0] <= '9':
		// fmt.Println("str")
		val, err := decodeString(br)
		if err != nil {
			return nil, err
		}
		ret = val
		return ret, nil
	case b[0] == 'i':
		// fmt.Println("int")
		val, err := decodeInt(br)
		if err != nil {
			return nil, err
		}
		ret = val
		return ret, nil
	case b[0] == 'l':
		// fmt.Println("list")
		val, err := decodeList(br)
		if err != nil {
			return nil, err
		}
		ret = val
		return ret, nil
	case b[0] == 'd':
		// fmt.Println("dict")
		val, err := decodeDict(br)
		if err != nil {
			return nil, err
		}
		ret = val
		return ret, nil
	default:
		return nil, ErrFmt
	}
}
