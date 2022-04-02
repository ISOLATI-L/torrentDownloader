package torrent

import (
	"Torrent_downloader/Bencode"
	"bytes"
	"crypto/sha1"
	"errors"
)

const SHALEN int = 20

var ErrLen error

func init() {
	ErrLen = errors.New("error length")
}

func hash(bo Bencode.Bobject) [SHALEN]byte {
	var buf bytes.Buffer
	bo.Bencode(&buf)
	return sha1.Sum(buf.Bytes())
}

func splitPiece(pieces string) ([][SHALEN]byte, error) {
	return splitHash([]byte(pieces))
}

func splitHash(hashes []byte) ([][SHALEN]byte, error) {
	if len(hashes)%SHALEN != 0 {
		return nil, ErrLen
	}
	numPieces := len(hashes) / SHALEN
	ret := make([][SHALEN]byte, numPieces)
	for i := 0; i < numPieces; i++ {
		copy(ret[i][:], hashes[i*SHALEN:(i+1)*SHALEN])
	}
	return ret, nil
}
