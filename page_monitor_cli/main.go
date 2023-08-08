package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoring = 10
	delay      = 3 * time.Second
)

func main() {
	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Displaying logs...")
			printLogs()
		case 0:
			fmt.Println("Closing application...")
			os.Exit(0)
		default:
			fmt.Println("Command not recognized.")
			os.Exit(-1)

		}
	}

}

func showIntroduction() {
	name := "Lucas"
	version := 1.1
	fmt.Println("Hello,", name)
	fmt.Println("This program is running at version", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("0 - Exit program")
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("The chosen comand was", readCommand)

	return readCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readFileSites()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(i, site)
		}
		time.Sleep(delay)
		fmt.Println("")
	}

	fmt.Println("")

}

func testSite(i int, site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("- Site", site, "was loaded successfuly. Status code:", response.StatusCode)
		registerLog(site, true)
	} else {
		fmt.Println("- Site:", site, "has some problem. Status code:", response.StatusCode)
		registerLog(site, true)
	}
}

func readFileSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			fmt.Println("An error has occurred:", err)
			break
		}
	}
	fmt.Println(sites)

	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
