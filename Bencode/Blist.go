package Bencode

import (
	"bufio"
	"io"
)

type Blist []Bobject

func (val *Blist) Bencode(w io.Writer) (wLen int) {
	wLen = 2
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	bw.WriteByte('l')
	for _, elem := range *val {
		wLen += elem.Bencode(bw)
	}
	bw.WriteByte('e')
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

func decodeList(r io.Reader) (*Blist, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	if b != 'l' {
		return nil, ErrFmt
	}

	ret := make(Blist, 0)
	for {
		b, err := br.Peek(1)
		if err != nil {
			return nil, err
		}
		if b[0] == 'e' {
			br.ReadByte()
			break
		}
		elem, err := Parse(br)
		if err != nil {
			return nil, err
		}
		ret = append(ret, elem)
	}
	return &ret, nil
}
