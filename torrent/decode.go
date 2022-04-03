package torrent

import (
	"reflect"
	"torrentDownloader/Bencode"
)

func FromBobject(bo Bencode.Bobject) (ret TorrentFile, err error) {
	rv := reflect.ValueOf(&ret).Elem()
	err = fromBobject(bo, rv)
	return ret, err
}

func fromBobject(bo Bencode.Bobject, rv reflect.Value) error {
	rvType := rv.Type()
	switch rvType.Kind() {
	case reflect.String:
		return fromBstring(bo, rv)
	case reflect.Int:
		return fromBint(bo, rv)
	case reflect.Array, reflect.Slice:
		return fromBlist(bo, rv)
	case reflect.Struct:
		return fromBdict(bo, rv)
	case reflect.Interface:
		bd, ok := bo.(*Bencode.Bdict)
		if !ok {
			return ErrFmt
		}
		_, isSingle := bd.GetWithOK("length")
		_, isMulti := bd.GetWithOK("files")
		if isSingle == isMulti {
			return ErrFmt
		} else if isSingle {
			rSingleFileInfo := reflect.ValueOf(
				&SingleFileInfo{},
			).Elem()
			err := fromBobject(bo, rSingleFileInfo)
			if err != nil {
				return err
			}
			rv.Set(rSingleFileInfo)
		} else if isMulti {
			rMultiFileInfo := reflect.ValueOf(
				&MultiFileInfo{},
			).Elem()
			err := fromBobject(bo, rMultiFileInfo)
			if err != nil {
				return err
			}
			rv.Set(rMultiFileInfo)
		}
		return nil
	default:
		// fmt.Println(rvType)
		return ErrFmt
	}
}

func fromBstring(bo Bencode.Bobject, rv reflect.Value) error {
	bs, ok := bo.(Bencode.Bstring)
	if !ok {
		return ErrFmt
	}
	rv.SetString(string(bs))
	return nil
}

func fromBint(bo Bencode.Bobject, rv reflect.Value) error {
	bi, ok := bo.(Bencode.Bint)
	if !ok {
		return ErrFmt
	}
	rv.SetInt(int64(bi))
	return nil
}

func fromBlist(bo Bencode.Bobject, rv reflect.Value) error {
	bl, ok := bo.(*Bencode.Blist)
	if !ok {
		return ErrFmt
	}
	rvType := rv.Type()
	rv.Set(reflect.MakeSlice(rvType, 0, 0))
	elemType := rvType.Elem()
	for _, boi := range *bl {
		elem := reflect.New(elemType).Elem()
		err := fromBobject(boi, elem)
		if err != nil {
			return err
		}
		rv.Set(reflect.Append(rv, elem))
	}
	return nil
}

func fromBdict(bo Bencode.Bobject, rv reflect.Value) error {
	bd, ok := bo.(*Bencode.Bdict)
	if !ok {
		return ErrFmt
	}
	rvType := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fieldType := rvType.Field(i)
		key := fieldType.Tag.Get("bencode")
		if len(key) == 0 {
			continue
		}
		value, ok := bd.GetWithOK(key)
		if !ok {
			continue
		}
		field := rv.Field(i)
		elemType := field.Type()
		elem := reflect.New(elemType).Elem()
		err := fromBobject(value, elem)
		if err != nil {
			return err
		}
		field.Set(elem)
	}
	return nil
}
