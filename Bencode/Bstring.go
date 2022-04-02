package Bencode

import (
	"bufio"
	"io"
)

type Bstring string

func (val Bstring) Bencode(w io.Writer) (wLen int) {
	wLen = 1
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	wLen += writeDecimal(bw, len(val))
	bw.WriteByte(':')
	len, err := bw.WriteString(string(val))
	if err != nil {
		return 0
	}
	wLen += len
	err = bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

func decodeString(r io.Reader) (Bstring, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	strLen := 0
	idx := 0
	for {
		b, err := br.ReadByte()
		if err != nil {
			return "", err
		}
		idx++
		if b >= '0' && b <= '9' {
			strLen = strLen*10 + int(b-'0')
		} else if b == ':' {
			break
		} else {
			return "", ErrFmt
		}
	}
	if idx < 2 {
		return "", ErrFmt
	}
	buf := make([]byte, strLen)
	_, err := io.ReadAtLeast(br, buf, strLen)
	if err != nil {
		return "", err
	}
	val := Bstring(buf)
	// fmt.Println(strLen)
	// fmt.Println(val)
	return val, nil
}
