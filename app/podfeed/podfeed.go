package podfeed

import (
  "github.com/mmcdole/gofeed"
  "strconv"
)

type Episode struct {
  Title string
  Url string
  Length int
  MediaType string
  PubDate string
}

var episode Episode
var episodes []Episode

func OpenFeed(URL string) (f *gofeed.Feed, e error) {
  fp := gofeed.NewParser()
  f, e = fp.ParseURL(URL)
  return f, e
}

func GetTitle(f *gofeed.Feed) (t string) {
  t = f.Title
  return t
}

func GetEpisodes(f *gofeed.Feed) (episodes []Episode) {
  for _, item := range f.Items {
    var url, mediatype string
    var length int
    for _, enclosure := range item.Enclosures {
      url = enclosure.URL
      length, _ = strconv.Atoi(enclosure.Length)
      mediatype = enclosure.Type
    }
    title := item.Title
    pubdate := item.Published
    episodes = append(episodes, Episode{
      Title: title,
      Url: url,
      Length: length,
      MediaType: mediatype,
      PubDate: pubdate,
    })
  }
  return episodes
}
