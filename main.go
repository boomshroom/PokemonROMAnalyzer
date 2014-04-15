/**
 * Created by angelo on 14/04/14.
 */
package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/boomshroom/PokemonROMAnalyzer/specific/gen3"
)

func main() {
	fmt.Printf("Enter ROM Location:")
	path:=new(*string)
	_,err:=fmt.Scanln(path)
	if err!=nil{
		fmt.Fprint(os.Stderr,"Error reading path:\n",err)
		return
	}
	file, err:=os.Open(path)
	if err!=nil{
		fmt.Fprint(os.Stderr,"Error opening file:\n",err)
		return
	}
	info, err:=file.Stat()
	if err!=nil{
		fmt.Fprint(os.Stderr,"Error reading file:\n",err)
		return
	}else if info.IsDir(){
		fmt.Fprint(os.Stderr,"Path Given is not file")
		return
	}else if !checkExt(info.Name()){
		fmt.Fprint(os.Stderr,"Path Given is not ROM")
		return
	}

	rom:=file.Read(make([]byte,info.Size()))
	data:=gen3.Analyze(rom)
	for _,poke := range data.pokemon{
		fmt.Println(poke.name)
	}
}

func checkExt(name string) bool{
	if strings.HasSuffix(name,".gb"){ return true}
	if strings.HasSuffix(name,".sgb"){ return true}
	if strings.HasSuffix(name,".gbc"){ return true}
	if strings.HasSuffix(name,".gba"){ return true}
	if strings.HasSuffix(name,".nds"){ return true}
	return false
}

