package Bencode

import (
	"bufio"
	"io"
	"strconv"
)

type Bint int

func (val Bint) Bencode(w io.Writer) (wLen int) {
	wLen = 2
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	bw.WriteByte('i')
	wLen += writeDecimal(bw, int(val))
	bw.WriteByte('e')
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

func decodeInt(r io.Reader) (Bint, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	b, err := br.ReadByte()
	if err != nil {
		return 0, err
	}
	if b != 'i' {
		return 0, ErrFmt
	}
	numStr := ""
	for {
		b, err := br.ReadByte()
		if err != nil {
			return 0, err
		}
		if b != 'e' {
			numStr += string(b)
		} else {
			break
		}
	}
	val, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, ErrFmt
	}
	return Bint(val), nil
}

func writeDecimal(w *bufio.Writer, val int) (len int) {
	len = 0
	if val == 0 {
		w.WriteByte('0')
		len++
		return len
	}
	if val < 0 {
		w.WriteByte('-')
		len++
		val = -val
	}
	numStr := ""
	for val > 0 {
		numStr = string('0'+byte(val%10)) + numStr
		val /= 10
	}
	nn, err := w.WriteString(string(numStr))
	if err != nil {
		return 0
	}
	len += nn
	return len
}
