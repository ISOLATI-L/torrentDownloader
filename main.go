package main

import (
	"Torrent_downloader/Bencode"
	"Torrent_downloader/torrent"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile(
		`D:\Desktop\T\かぐや様は告らせたい？〜天才たちの恋愛頭脳戦〜.torrent`,
		// `D:\Desktop\T\4B6629BEEFFFB68226D5D84567B686B400C2EDCE.torrent`,
		os.O_RDONLY,
		0666,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bo, err := Bencode.Parse(f)
	err = f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%T\n", bo)
	fmt.Printf("%+v\n", bo)
	// announceList := bo.(*Bencode.Bdict).Get("announce-list").(*Bencode.Blist)
	// for _, announceL := range *announceList {
	// 	announceL := announceL.(*Bencode.Blist)
	// 	fmt.Printf("%+v\n", announceL)
	// 	for _, announce := range *announceL {
	// 		fmt.Printf("%+v\n", announce)
	// 	}
	// }
	torrentFile, err := torrent.FromBobject(bo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", torrentFile)
	torrentJson, err := json.MarshalIndent(torrentFile, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", string(torrentJson))

	wf, err := os.OpenFile(
		`D:\Desktop\T\test.torrent`,
		os.O_RDWR|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	fmt.Printf("%d\n", bo.Bencode(wf))
	err = wf.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	bo2, err := torrentFile.ToBobject()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("bo2:%T\n", bo2)
	fmt.Printf("bo2:%+v\n", bo2)
	wf, err = os.OpenFile(
		`D:\Desktop\T\test2.torrent`,
		os.O_RDWR|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	fmt.Printf("%d\n", bo2.Bencode(wf))
	err = wf.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
