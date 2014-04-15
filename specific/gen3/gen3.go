package gen3

import (
	"github.com/boomshroom/PokemonROMAnalyzer/base"
	"strconv"
)

var rom []byte

func Analyze(romPar []byte)base.RomData{
	rom=romPar
	data:=&base.RomData{name: "emerald", make([]Pokemon,387)}
	initCharMap()

	pokemonNamePointer:=uint32(rom[324]) | uint32(rom[325])<<8 | uint32(rom[326])<<16 | uint32(rom[327])<<24

	for i,_ := range data.pokemon{
		data.pokemon[i]=getString(pokemonNamePointer+i*11, 11)

	}

	return data
}

var charMap []byte = make([]byte,256)
const (
	alphabetUpper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLower string = "abcdefghijklmnopqrstuvwxyz"
)

func initCharMap() {
	for i:=0; i<10; i++{
		charMap[i+0xA1]=strconv.Itoa(i)[0]
	}
	for i,char:= range alphabetUpper{
		charMap[i+0xBB]=char
		charMap[i+0Xd5]=alphabetLower[i]
	}
	charMap[0xFE]='\n'
}


func getString(index int,length int)string {
	temp:=make([]byte,length)
	for i:=0;i<length;i++{
		temp[i]=rom[index+ i*length]
	}
	return string(temp)
}

