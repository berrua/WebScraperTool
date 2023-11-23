package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

// a data structure to store the scraped data
type News struct {
	name        string
	description string
	date        string
}

func main() {
	displayAsciiArt()

	// flags
	firstSite := flag.Bool("1", false, "It is used to scrape news from the first website")
	secondSite := flag.Bool("2", false, "It is used to scrape news from the second website")
	thirdSite := flag.Bool("3", false, "It is used to scrape news from the third website")
	hideDate := flag.Bool("date", false, "It is used to hide the date part")
	hideDesc := flag.Bool("description", false, "It is used to hide the description part")
	flag.Parse()

	var wg sync.WaitGroup
	if *firstSite {
		wg.Add(1)
		go func() {
			c := color.New(color.FgGreen, color.Bold)
			defer c.Println("↳ Data scraped from: \"https://thehackernews.com/\"\n")
			defer wg.Done()
			collectData(1, hideDate, hideDesc)
		}()
	}

	if *secondSite {
		wg.Add(1)
		go func() {
			c := color.New(color.FgGreen, color.Bold)
			defer c.Println("↳ Data scraped from: \"https://cybersecuritynews.com/\"\n")
			defer wg.Done()
			collectData(2, hideDate, hideDesc)
		}()
	}

	if *thirdSite {
		wg.Add(1)
		go func() {
			c := color.New(color.FgGreen, color.Bold)
			defer c.Println("↳ Data scraped from: \"https://thecyberwire.com/\"\n")
			defer wg.Done()
			collectData(3, hideDate, hideDesc)
		}()
	}

	wg.Wait()
}

func collectData(value int, hideDate, hideDesc *bool) {
	if value == 1 {

		//that will contain the scraped data
		var news []News

		c := colly.NewCollector()

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
		})

		c.OnHTML("a.story-link", func(e *colly.HTMLElement) {
			// check if the element has unwanted text
			if strings.Contains(e.Attr("class"), "clear home-right") ||
				strings.Contains(e.Text, "Easily discover your employees' SaaS usage") {
				return // skip
			}

			// initializing a new News instance
			new := News{}

			// scraping the data of interest
			new.name = e.ChildText(".home-title")
			new.description = e.ChildText(".home-desc")
			new.date = e.ChildText(".item-label .h-datetime")

			// adding the news instance with scraped data to the list of news
			news = append(news, new)

			// printing the scraped news to the terminal
			blue := color.New(color.FgBlue, color.Bold).SprintFunc()
			fmt.Printf("%s%s", blue(fmt.Sprintf("%d - ", len(news))), blue("New\n")+new.name)

			// check if the -description flag is used or not
			if !*hideDesc {
				yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
				fmt.Printf("%s", yellow("\nDescription :\n")+new.description)
			}

			// check if the -date flag is used or not
			if !*hideDate {
				magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()
				fmt.Printf("%s", magenta("\nDate :\n")+new.date)
			}

			fmt.Print("\n\n")
		})

		err := c.Visit("https://thehackernews.com/")
		if err != nil {
			log.Fatal(err)
		}

	} else if value == 2 {

		//that will contain the scraped data
		var news []News

		c := colly.NewCollector()

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
		})

		c.OnHTML(".td_module_10", func(e *colly.HTMLElement) {
			// initializing a new News instance
			new := News{}

			// scraping the data of interest
			new.name = e.ChildText(".td-module-title")
			new.description = e.ChildText(".td-excerpt")
			new.date = e.ChildText(".td-post-date time")

			// adding the news instance with scraped data to the list of news
			news = append(news, new)

			// printing the scraped news to the terminal
			blue := color.New(color.FgBlue, color.Bold).SprintFunc()
			fmt.Printf("%s%s", blue(fmt.Sprintf("%d - ", len(news))), blue("New\n")+new.name)

			// check if the -description flag is used or not
			if !*hideDesc {
				yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
				fmt.Printf("%s", yellow("\nDescription :\n")+new.description)
			}

			// check if the -date flag is used or not
			if !*hideDate {
				magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()
				fmt.Printf("%s", magenta("\nDate :\n")+new.date)
			}

			fmt.Print("\n\n")
		})

		err := c.Visit("https://cybersecuritynews.com/")
		if err != nil {
			log.Fatal(err)
		}

	} else if value == 3 {

		//that will contain the scraped data
		var news []News

		c := colly.NewCollector()

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
		})

		c.OnHTML(".card", func(e *colly.HTMLElement) {
			// initializing a new News instance
			new := News{}

			// scraping the data of interest
			new.name = e.ChildText(".card .title")
			new.description = e.ChildText(".card .description")
			new.date = e.ChildText(".card .meta .meta-text")

			// adding the news instance with scraped data to the list of news
			news = append(news, new)

			// printing the scraped news to the terminal
			blue := color.New(color.FgBlue, color.Bold).SprintFunc()
			fmt.Printf("%s%s", blue(fmt.Sprintf("%d - ", len(news))), blue("New\n")+new.name)

			// check if the -description flag is used or not
			if !*hideDesc {
				yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
				fmt.Printf("%s", yellow("\nDescription :\n")+new.description)
			}

			// check if the -date flag is used or not
			if !*hideDate {
				magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()
				fmt.Printf("%s", magenta("\nDate :\n")+new.date)
			}

			fmt.Print("\n\n")
		})

		err := c.Visit("https://thecyberwire.com/")
		if err != nil {
			log.Fatal(err)
		}

	}

}

func displayAsciiArt() {
	data, err := ioutil.ReadFile("ascii_art.txt")
	if err != nil {
		fmt.Println("ASCII art could not be displayed:", err)
		return
	}

	fmt.Println(string(data))
}
