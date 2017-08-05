package controllers

import (
	"github.com/revel/revel"
	"podfeed"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) ReturnError(err error, errString string) revel.Result {
			c.Flash.Error(errString)
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(App.Index)
}

func (c App) Getepisodes(podcasturl string) revel.Result {
			feed, err := podfeed.OpenFeed(podcasturl)
			if err != nil {
				c.Flash.Error("Unable to open podcast")
				c.Validation.Keep()
				c.FlashParams()
				return c.Redirect(App.Index)
		  }
			podcastTitle := podfeed.GetTitle(feed)
			podcastEpisodes := podfeed.GetEpisodes(feed)
			return c.Render(podcastTitle, podcastEpisodes)
}

func (c App) Downloadepisodes() revel.Result {
	var urls []string
	var downloadLocation string

	c.Params.Bind(&urls, "urls")
	c.Params.Bind(&downloadLocation, "downloadLocation")
	return c.Render(urls, downloadLocation)
}
