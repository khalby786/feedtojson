# feedtojson
a simple API that parses RSS, Atom and JSON feeds using [gofeed](https://github.com/mmcdole/gofeed).

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fkhalby786%2Ffeedtojson)

## usage


`GET https://feedtojson.vercel.app/your-feed-url-here`

[`https://feedtojson.vercel.app/https%3A%2F%2Fxkcd.com%2Frss.xml`](https://feedtojson.vercel.app/https%3A%2F%2Fxkcd.com%2Frss.xml)
[`https://feedtojson.vercel.app/https%3A%2F%2Fxkcd.com%2Fatom.xml`](https://feedtojson.vercel.app/https%3A%2F%2Fxkcd.com%2Fatom.xml)
[`https://feedtojson.vercel.app/https%3A%2F%2Fwww.jsonfeed.org%2Ffeed.json`](https://feedtojson.vercel.app/https%3A%2F%2Fwww.jsonfeed.org%2Ffeed.json)


### RSS

[`https://xkcd.com/rss.xml`](https://xkcd.com/rss.xml)

```json
{
  "title": "xkcd.com",
  "description": "xkcd.com: A webcomic of romance and math humor.",
  "link": "https://xkcd.com/",
  "links": [
    "https://xkcd.com/"
  ],
  "language": "en",
  "items": [...],
  "feedType": "rss",
  "feedVersion": "2.0"
}
```

### Atom

[`https://xkcd.com/atom.xml`](https://xkcd.com/atom.xml)

```json
{
  "title": "xkcd.com",
  "link": "https://xkcd.com/",
  "links": [
    "https://xkcd.com/"
  ],
  "updated": "2024-02-09T00:00:00Z",
  "updatedParsed": "2024-02-09T00:00:00Z",
  "language": "en",
  "items": [...],
  "feedType": "atom",
  "feedVersion": "1.0"
}
```

### JSON

[`https://www.jsonfeed.org/feed.json`](https://www.jsonfeed.org/feed.json)

```json
{
  "title": "JSON Feed",
  "link": "https://www.jsonfeed.org/",
  "feedLink": "https://www.jsonfeed.org/feed.json",
  "links": [
    "https://www.jsonfeed.org/",
    "https://www.jsonfeed.org/feed.json"
  ],
  "published": "2020-08-07T11:44:36-05:00",
  "publishedParsed": "2020-08-07T11:44:36-05:00",
  "image": {
    "url": "https://micro.blog/jsonfeed/avatar.jpg"
  },
  "items": [...],
  "feedType": "json",
  "feedVersion": "https://jsonfeed.org/version/1"
}
```

<br>

### why

i wanted to add [my blog](https://blog.khaleelgibran.com) posts to my [website](https://khaleelgibran.com), and after trying three different packages from npm, none of them worked. additionally, my blog was giving me cors issues although it shouldn't because of my netlify configuration. i then thought it'd be a pretty quick and easy challenge to build a simple Go function on Vercel that'd parse my feed as JSON.

### how

it uses [gofeed](https://github.com/mmcdole/gofeed), which is  a powerful and flexible Go library designed for parsing RSS, Atom, and JSON feeds. Vercel handles the routing using my custom configuration in `vercel.json`.


## license

this project is licensed under the [MIT License](https://raw.githubusercontent.com/khalby786/feedtojson/master/LICENSE.md).

<br><br>

made by [khaleel](https://khaleelgibran.com)