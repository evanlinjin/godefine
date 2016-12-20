package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/evanlinjin/wordsapi_go"
)

const confDirName string = ".godefine"
const keyFileName string = "api.key"

func main() {
	// Check configuration file.
	key, err := config()
	if err != nil {
		printQuery(fmt.Sprint(err))
		return
	}
	wordsapi.SetApiKey(key)

	// Check command line arguments.
	switch len(os.Args) {
	case 1:
		printErr1("No word entered.")
	case 2:
		displayGetWord(os.Args[1])
	}

	//wordPtr := flag.String("", "", "word to define")
	//flag.Parse()

	//fmt.Println("[ARGS]", args)
	//fmt.Println("[ WRD]", *wordPtr)
}

func config() (key string, err error) {
	// Grab current user and location of confing file..
	usr, _ := user.Current()
	confDir := fmt.Sprintf("%s/%s/", usr.HomeDir, confDirName)

	// Read configuration file for godefine.
	confFd, err := ioutil.ReadFile(confDir + keyFileName)
	if err != nil {
		// Tell user that api.key is non-existant.
		printErr1(fmt.Sprint(err))

		// Query user to enter a key.
		printQuery("Enter your WordsAPI Key")
		fmt.Scanln(&key)
		key = strings.TrimSpace(key)

		// Store key in file.
		os.MkdirAll(confDir, 0700)
		ioutil.WriteFile(confDir+keyFileName, []byte(key), 0700)
		err = nil

	} else {
		key = string(confFd)
	}
	return
}

func displayGetWord(word string) {
	obj, _ := wordsapi.GetWord(word)
	switch obj.Word {
	case "":
		printErr1("Word not found.")
	default:
		printH1("word", obj.Word)
		printH2("definitions", len(obj.Results))
		for i := 0; i < len(obj.Results); i++ {
			printDef1(i+1, obj.Results[i].Definition, obj.Results[i].PartOfSpeech)
		}
	}
}
