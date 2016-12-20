package main

import (
	"flag"
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
	// Check command line flags.
	flags := Flags{
		Key: flag.String("key", "", "usage"),
	}

	// Check command line arguments.
	if len(os.Args) == 1 || strings.HasPrefix(os.Args[1], "-") {
		printErr1("No word entered.")
		return
	}

	// Check configuration file.
	keyFromConfig := config(*flags.Key)
	wordsapi.SetApiKey(keyFromConfig)

	// Display data from WordsAPI.
	displayGetWord(os.Args[1])
}

func config(keyFromFlag string) (key string) {
	// Grab current user and location of confing file..
	usr, _ := user.Current()
	configDir := fmt.Sprintf("%s/%s/", usr.HomeDir, confDirName)

	// Declare local functions.
	saveKey := func(kStr string) {
		os.MkdirAll(configDir, 0700)
		ioutil.WriteFile(configDir+keyFileName, []byte(kStr), 0700)
	}

	// Use key from flag if flag defined.
	if keyFromFlag != "" {
		key = keyFromFlag
		saveKey(keyFromFlag)
		return
	}

	// Use key from user input.
	configFileData, err := ioutil.ReadFile(configDir + keyFileName)
	if err != nil {
		printQuery("Enter your WordsAPI Key")
		fmt.Scanln(&key)
		saveKey(key)
		return
	}

	// Use key from file.
	key = string(configFileData)
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
