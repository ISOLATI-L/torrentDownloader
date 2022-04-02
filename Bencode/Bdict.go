package Bencode

import (
	"Torrent_downloader/orderedMap"
	"bufio"
	"io"
)

// type Bdict map[string]Bobject

type Bdict struct {
	orderedMap.OrderedMap[string, Bobject]
}

func NewBdict() (ret Bdict) {
	ret.Init()
	return ret
}

func (val *Bdict) Bencode(w io.Writer) (wLen int) {
	wLen = 2
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	bw.WriteByte('d')
	val.Range(func(k string, v Bobject) {
		wLen += Bstring(k).Bencode(bw)
		wLen += v.Bencode(bw)
	})
	// for k, v := range *val {
	// 	wLen += Bstring(k).Bencode(bw)
	// 	wLen += v.Bencode(bw)
	// }
	bw.WriteByte('e')
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

func decodeDict(r io.Reader) (*Bdict, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	if b != 'd' {
		return nil, ErrFmt
	}

	ret := NewBdict()
	for {
		b, err := br.Peek(1)
		if err != nil {
			return nil, err
		}
		if b[0] == 'e' {
			br.ReadByte()
			break
		}
		k, err := decodeString(br)
		if err != nil {
			return nil, err
		}
		v, err := Parse(br)
		if err != nil {
			return nil, err
		}
		ret.Set(string(k), v)
	}
	return &ret, nil
}
