package main

import (
	"fmt"
	"os"

	"github.com/evanlinjin/wordsapi_go"
	"github.com/fatih/color"
)

func main() {
	switch len(os.Args) {
	case 1:
		color.Yellow("No word entered.")
		return
	case 2:
		wordsapi.SetApiKey("")
		obj, err := wordsapi.GetWord(os.Args[1])
		if err != nil {
			color.Red(fmt.Sprint(err))
		}
		printH1("word", obj.Word)
		printH2("definitions", len(obj.Results))
		for i := 0; i < len(obj.Results); i++ {
			printDef1(i+1, obj.Results[i].Definition, obj.Results[i].PartOfSpeech)
		}
		//fmt.Println(obj.Results)
	}

	//wordPtr := flag.String("", "", "word to define")
	//flag.Parse()

	//fmt.Println("[ARGS]", args)
	//fmt.Println("[ WRD]", *wordPtr)
}
