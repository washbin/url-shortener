# url-shortener
A url shortener created using go

## Current Features
- Simplistic
- ~~Uses only Go stdlib~~ Go-redis for redis interaction
- Only alphanumerics(_ inclusive) and - available as slug
- Peristent storage with redis(as only simple key value pair is required at the moment)

## TODO
- Add at GET / route as a simple web frontend to the api

## Available endpoints
- `POST /` to create a new shortened url
	> with body `{"url": "your url", "slug": "your desired slug"}`
- `GET /{slug}` redirect to the original url
- `DELETE /{slug}` removes the short url

> Note: Credits to [Samrid Pandit](https://github.com/caffeineduck) and do check out his superior python implementation at [shotcut](https://github.com/caffeineduck/shotcut)
