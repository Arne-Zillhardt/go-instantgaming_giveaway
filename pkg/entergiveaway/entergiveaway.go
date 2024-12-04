package entergiveaway

import (
	"log"
	"strconv"

	"github.com/arne-zillhardt/instantgaming_giveaway/pkg/dataprovider"
	"github.com/go-vgo/robotgo"
)

func setUp() {
	robotgo.SetDelay(100)

	browserPositionX, _ := strconv.Atoi(dataprovider.GetBrowserPositionX())
	browserPositionY, _ := strconv.Atoi(dataprovider.GetBrowserPositionY())

	robotgo.Move(browserPositionX, browserPositionY)
	robotgo.Click("left", false)
	robotgo.Sleep(1)
}

func EnterAllGiveaways() {
	setUp()

	urlsToParticipate := dataprovider.GetUrls()
	for _, v := range urlsToParticipate {
		enterGiveaway(v)
	}

	robotgo.KeyTap("w", robotgo.CmdCtrl())
}

func enterGiveaway(urlToVisit dataprovider.ParticipatingUrl) {
	if !urlToVisit.ActiveGiveaway {
		return
	}

	//remove "coockies"

	robotgo.KeyTap("t", robotgo.CmdCtrl())
	typeUrl(urlToVisit.UrlToVisit)
	robotgo.KeyPress(robotgo.Enter)
	robotgo.Sleep(6)
	clickParticipate()
	robotgo.Sleep(2)
	scrollToExtras()
	robotgo.Sleep(1)

	for {
		if clickExtraPoints() == "0c0d0e" {
			break
		}

		robotgo.ScrollDir(3, "down")
		robotgo.KeyTap("r", robotgo.CmdCtrl())
		robotgo.Sleep(2)
	}

	robotgo.KeyTap("w", robotgo.CmdCtrl())
}

func clickExtraPoints() string {
	extrasXPosition, _ := strconv.Atoi(dataprovider.GetExtrasXPosition())
	extrasYPositionStart, _ := strconv.Atoi(dataprovider.GetExtrasYPositionStart())
	extrasYPositionEnd, _ := strconv.Atoi(dataprovider.GetExtrasYPositionEnd())

	for i := extrasYPositionStart; i < extrasYPositionEnd; i++ {
		colorAsString := robotgo.GetPixelColor(extrasXPosition, i)
		if colorAsString == "0c0d0e" {
			return "0c0d0e"
		}

		greenValueInColor, _ := strconv.ParseUint(colorAsString[2:4], 16, 8)
		redValueInColor, _ := strconv.ParseUint(colorAsString[0:2], 16, 8)

		if colorAsString == "1d2021" || (greenValueInColor > 180 && redValueInColor > 180) {
			continue
		}

		if greenValueInColor > redValueInColor {
			i += 55
			log.Println("Already claimed")
			continue
		}

		log.Println("Claim bonus")
		robotgo.Move(extrasXPosition, i+10)
		robotgo.Click("left", false)
		robotgo.Sleep(1)
		robotgo.KeyTap("w", robotgo.CmdCtrl())

		i += 60
		robotgo.Sleep(1)
	}

	return robotgo.GetPixelColor(extrasXPosition, 1025)
}

func typeUrl(urlToType string) {
	for _, v := range urlToType {
		if v == '/' {
			robotgo.KeyTap(robotgo.Key7, robotgo.Shift)
			continue
		}
		if v == ':' {
			robotgo.KeyTap(".", robotgo.Shift)
			continue
		}

		robotgo.KeyPress(string(v))
	}
}

func clickParticipate() {
	participationCheckPositionX, _ := strconv.Atoi(dataprovider.GetParticipationCheckXPosition())
	participationCheckPositionY, _ := strconv.Atoi(dataprovider.GetParticipationCheckYPosition())
	participationPositionX, _ := strconv.Atoi(dataprovider.GetParticipationXPosition())
	participationPositionY, _ := strconv.Atoi(dataprovider.GetParticipationYPosition())

	if robotgo.GetPixelColor(participationCheckPositionX, participationCheckPositionY) != "1d2021" {
		robotgo.Move(participationPositionX, participationPositionY + 145)
		robotgo.Click("left", false)
		return
	}

	robotgo.Move(participationPositionX, participationPositionY)

	robotgo.Click("left", false)
}

func scrollToExtras() {
	participationCheckPositionX, _ := strconv.Atoi(dataprovider.GetParticipationCheckXPosition())
	participationCheckPositionY, _ := strconv.Atoi(dataprovider.GetParticipationCheckYPosition())

	if robotgo.GetPixelColor(participationCheckPositionX, participationCheckPositionY) != "1d2021" {
		robotgo.ScrollDir(8, "down")
		return
	}

	robotgo.ScrollDir(6, "down")
}
