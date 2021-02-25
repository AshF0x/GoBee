package main

import (
	"encoding/json"
	"fmt"
	"gobee/fetcher"
	"gobee/mailer"
	"io/ioutil"
	"log"
	"os"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
)

func versionAlert(name string, tested string, current string) {
	fmt.Printf("[%v] Different Versions detected for %v!\n", emoji.Fire, name)
	color.Red("Tested Version: %v\n", tested)
	color.Green("Newest Version: %v\n\n", current)
}

func main() {
	// Struct for tested versions
	type Tested struct {
		Firefox      string
		FirefoxESR   string
		Thunderbird  string
		Chrome       string
		EdgeMajor    string
		Edge         string
		Office2013   string
		Office2010   string
		Office201619 string
		Ubuntu       string
		MacOS        string
		Windows10    string
	}
	file := "tested.json"
	jsonFile, err := os.Open("./" + file)
	if err != nil {
		fmt.Println(err)
	}
	color.Green("File loaded successfully") // Replace this with a checkmark emoji
	defer jsonFile.Close()
	// Read the opened file
	data, _ := ioutil.ReadAll(jsonFile)
	var tested Tested
	err = json.Unmarshal(data, &tested)
	if err != nil {
		log.Fatal(err)
	}

	updateCounter := 0
	updateList := []byte("Outdated:\n")
	// Firefox
	currentFirefox := fetcher.Slicer(fetcher.Fetch("https://www.mozilla.org/en-US/firefox/releases/"), "data-latest-firefox=\"", "\" data-esr-versions=")
	if tested.Firefox != currentFirefox {
		updateList = []byte(fmt.Sprintf("%s---Firefox---\nTested: %v\nNew: %v\n", updateList, tested.Firefox, currentFirefox))
		updateCounter++
		versionAlert("Firefox", tested.Firefox, currentFirefox)
	} else {
		fmt.Printf("[%v] Firefox: %v\n\n", emoji.Frog, currentFirefox)
	}
	// Thunderbird
	currentThunderbird := fetcher.Slicer(fetcher.Fetch("https://www.mozilla.org/en-US/thunderbird/releases/"), "en-US/thunderbird/", "/releasenotes")
	if tested.Thunderbird != currentThunderbird {
		updateList = []byte(fmt.Sprintf("%s---Thunderbird---\nTested: %v\nNew: %v\n", updateList, tested.Thunderbird, currentThunderbird))
		updateCounter++
		versionAlert("Thunderbird", tested.Thunderbird, currentThunderbird)
	} else {
		fmt.Printf("[%v] Thunderbird: %v\n\n", emoji.Frog, currentThunderbird)
	}
	// Chrome
	currentChrome := fetcher.Slicer(fetcher.Fetch("https://en.wikipedia.org/wiki/Google_Chrome"), "Windows, macOS, Linux</th><td>", "<")
	if tested.Windows10 != currentChrome {
		updateList = []byte(fmt.Sprintf("%s---Chrome---\nTested: %v\nNew: %v\n", updateList, tested.Chrome, currentChrome))
		updateCounter++
		versionAlert("Chrome", tested.Chrome, currentChrome)
	} else {
		fmt.Printf("[%v] Windows10: %v\n\n", emoji.Frog, currentChrome)
	}
	// Ubuntu
	currentUbuntu := fetcher.Slicer(fetcher.Fetch("https://distrowatch.com/news/distro/ubuntu.xml"), "<title>Distribution Release: Ubuntu ", "</title>")
	if tested.Ubuntu != currentUbuntu {
		updateList = []byte(fmt.Sprintf("%s---Ubuntu---\nTested: %v\nNew: %v\n", updateList, tested.Ubuntu, currentUbuntu))
		updateCounter++
		versionAlert("Ubuntu", tested.Ubuntu, currentUbuntu)
	} else {
		fmt.Printf("[%v] Ubuntu: %v\n\n", emoji.Frog, currentUbuntu)
	}
	// MacOS
	currentMacOS := fetcher.Slicer(fetcher.Fetch("https://en.wikipedia.org/wiki/MacOS"), "Latest release</a></th><td>", "<")
	if tested.MacOS != currentMacOS {
		updateList = []byte(fmt.Sprintf("%s---MacOS---\nTested: %v\nNew: %v\n", updateList, tested.MacOS, currentMacOS))
		updateCounter++
		versionAlert("MacOS", tested.MacOS, currentMacOS)
	} else {
		fmt.Printf("[%v] MacOS: %v\n\n", emoji.Frog, currentMacOS)
	}
	// Windows 10
	currentWindows10 := fetcher.Slicer(fetcher.Fetch("https://en.wikipedia.org/wiki/Windows_10"), "Latest release</a></th><td>", "<")
	if tested.Windows10 != currentWindows10 {
		updateList = []byte(fmt.Sprintf("%s---Windows10---\nTested: %v\nNew: %v\n", updateList, tested.Windows10, currentWindows10))
		updateCounter++
		versionAlert("Windows10", tested.Windows10, currentWindows10)
	} else {
		fmt.Printf("[%v] Windows10: %v\n\n", emoji.Frog, currentWindows10)
	}
	if updateCounter != 0 {
		fmt.Printf("%v updates in total\n", updateCounter)
		fmt.Printf("Please update %v after testing and updating\n", file)
		mailer.Mail(updateList)
	}
}
