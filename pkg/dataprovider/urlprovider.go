package dataprovider

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

type ParticipatingUrl struct {
	UrlToVisit     string
	ActiveGiveaway bool
}

var c *colly.Collector

func GetUrls() []ParticipatingUrl {
	configData := getConfigData()
	c = colly.NewCollector(colly.AllowedDomains("www.instant-gaming.com", "instant-gaming.com"))
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s\n", r.URL)
	})

	var urlsToVisit []ParticipatingUrl
	i := 0
	for _, v := range configData.Alive {
		url := "https://www.instant-gaming.com/fr/giveaway/" + v
		urlsToVisit = append(urlsToVisit, ParticipatingUrl{
			UrlToVisit:     url,
			ActiveGiveaway: checkIfGiveawayIsActive(url),
		})
		i++
	}

	log.Println()
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	return urlsToVisit
}

func checkIfGiveawayIsActive(urlToCheck string) bool {
	activeGiveaway := false
	c.OnHTML("div button.button.validate", func(h *colly.HTMLElement) {
		if strings.Contains(h.Text, "Participer") {
			activeGiveaway = true
		}
	})

	if err := c.Visit(urlToCheck); err != nil {
		log.Fatal("Error while scraping: ", err)
	}

	return activeGiveaway
}
