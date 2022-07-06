package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

// Fetch data from %appdata%
func UsersDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Sample code from goquery
// Edited for getting data from itch.io
func ItchIoScrape() {
	// Request the HTML page.
	res, err := http.Get("https://kay-yu.itch.io/holocure")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".info_panel_wrapper .game_info_panel_widget").Find("td").Each(func(i int, s *goquery.Selection) {
		if i < 2 {
			prefix := s.Find("b").Text()
			result := strings.TrimPrefix(s.Text(), prefix)
			println(result)

		}
	})
}

// Check current holocure version
// From version.ini
func LocalVersion() {
	fmt.Println("Cheking current version...")
	dirAppData := UsersDir()

	//get contents of version.ini
	content, err := os.ReadFile(dirAppData + "\\AppData\\Local\\holocure\\version.ini")
	if err != nil {
		log.Fatal(err)
	}

	//print version.ini
	fmt.Println(string(content))
}

// Main app
func main() {

	println("===Holocure Update checker===\n")

	color.Set(color.FgCyan)
	println("# Menu #")
	color.Unset()

	color.Set(color.FgYellow)
	println(" 1 -> Update Status : Check update status from itch.io")
	println(" 2 -> Check Version : Display installed version of Holocure")
	println(" 3 -> Download : Open itch.io site directly in your browser")
	println(" 4 -> Exit : Close the app\n")
	color.Unset()

	for {

		// Initialize variable
		var choice string

		// Handle user input
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		fmt.Print("\n")

		// Choose a statement
		switch {

		case choice == "1":
			color.Set(color.FgBlue)
			ItchIoScrape()
			color.Unset()
			fmt.Print("\n")

		case choice == "2":
			color.Set(color.FgGreen)
			LocalVersion()
			color.Unset()
			fmt.Print("\n")

		case choice == "3":
			color.Set(color.FgMagenta)
			println("Opening link from your default browser")
			color.Unset()
			exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://kay-yu.itch.io/holocure").Start()
			fmt.Print("\n")

		case choice == "4":
			break

		case choice == "":
			color.Set(color.FgRed)
			println("ERROR: Missing Command")
			color.Unset()
			fmt.Print("\n")

		default:
			color.Set(color.FgRed)
			println("ERROR: Unknown command = ", choice)
			color.Unset()
			fmt.Print("\n")

		}

		if choice == "4" {
			break
		}

	}

}
