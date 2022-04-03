package torrent

import (
	"reflect"
	"torrentDownloader/Bencode"
)

func (tf *TorrentFile) ToBobject() (Bencode.Bobject, error) {
	return toBobject(reflect.ValueOf(tf))
}

func toBobject(rv reflect.Value) (Bencode.Bobject, error) {
	rvType := rv.Type()
	switch rvType.Kind() {
	case reflect.Interface, reflect.Pointer:
		return toBobject(rv.Elem())
	case reflect.String:
		return Bencode.Bstring(rv.String()), nil
	case reflect.Int:
		return Bencode.Bint(rv.Int()), nil
	case reflect.Array, reflect.Slice:
		list := make(Bencode.Blist, 0)
		for i := 0; i < rv.Len(); i++ {
			index := rv.Index(i)
			value, err := toBobject(index)
			if err != nil {
				return &list, nil
			}
			list = append(list, value)
		}
		return &list, nil
	case reflect.Struct:
		dict := Bencode.NewBdict()
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if !field.IsValid() || field.IsZero() {
				continue
			}
			fieldType := rvType.Field(i)
			key := fieldType.Tag.Get("bencode")
			if len(key) == 0 {
				continue
			}
			value, err := toBobject(field)
			if err != nil {
				return &dict, nil
			}
			dict.Set(key, value)
		}
		return &dict, nil
	default:
		return nil, ErrFmt
	}
}
