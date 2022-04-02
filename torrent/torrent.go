package torrent

import (
	"errors"
)

type TorrentFile struct {
	Announce     string     `bencode:"announce" json:"announce,omitempty"`
	AnnounceList [][]string `bencode:"announce-list" json:"announce-list,omitempty"`
	Comment      string     `bencode:"comment" json:"comment,omitempty"`
	CommentUtf8  string     `bencode:"comment.utf-8" json:"comment.utf-8,omitempty"`
	CreatedBy    string     `bencode:"created by" json:"created by,omitempty"`
	CreationDate int        `bencode:"creation date" json:"creation date,omitempty"`
	Encoding     string     `bencode:"encoding" json:"encoding,omitempty"`
	Info         Info       `bencode:"info" json:"info"`
}

type Info interface {
	fileInfo()
}

type SingleFileInfo struct {
	Length           int    `bencode:"length" json:"length"`
	MD5Sum           string `bencode:"md5sum" json:"md5sum,omitempty"`
	Name             string `bencode:"name" json:"name,omitempty"`
	NameUtf8         string `bencode:"name.utf-8" json:"name.utf-8,omitempty"`
	PieceLength      int    `bencode:"piece length" json:"piece length"`
	Pieces           string `bencode:"pieces" json:"pieces,omitempty"`
	Private          int    `bencode:"private" json:"private,omitempty"`
	Publisher        string `bencode:"publisher" json:"publisher,omitempty"`
	PublisherUtf8    string `bencode:"publisher.utf-8" json:"publisher.utf-8,omitempty"`
	PublisherUrl     string `bencode:"publisher-url" json:"publisher-url,omitempty"`
	PublisherUrlUtf8 string `bencode:"publisher-url.utf-8" json:"publisher-url.utf-8,omitempty"`
}

func (SingleFileInfo) fileInfo() {}

type MultiFileInfo struct {
	Files            []FileInfo `bencode:"files" json:"files"`
	Name             string     `bencode:"name" json:"name,omitempty"`
	NameUtf8         string     `bencode:"name.utf-8" json:"name.utf-8,omitempty"`
	PieceLength      int        `bencode:"piece length" json:"piece length"`
	Pieces           string     `bencode:"pieces" json:"pieces,omitempty"`
	Private          int        `bencode:"private" json:"private,omitempty"`
	Publisher        string     `bencode:"publisher" json:"publisher,omitempty"`
	PublisherUtf8    string     `bencode:"publisher.utf-8" json:"publisher.utf-8,omitempty"`
	PublisherUrl     string     `bencode:"publisher-url" json:"publisher-url,omitempty"`
	PublisherUrlUtf8 string     `bencode:"publisher-url.utf-8" json:"publisher-url.utf-8,omitempty"`
}

type FileInfo struct {
	Length   int      `bencode:"length" json:"length"`
	MD5Sum   string   `bencode:"md5sum" json:"md5sum,omitempty"`
	Path     []string `bencode:"path" json:"path,omitempty"`
	PathUtf8 []string `bencode:"path.utf-8" json:"path.utf-8,omitempty"`
}

func (MultiFileInfo) fileInfo() {}

var ErrFmt error

func init() {
	ErrFmt = errors.New("error format")
}
