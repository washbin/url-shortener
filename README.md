# url-shortener
A url shortener created using go

## Current Features
- Simplistic
- Uses only Go stdlib
- Only alphanumerics(_ inclusive) and - available as slug

## TODO
- Peristent storage `May fix soon`

## Available endpoints
- `POST /` to create a new shortened url
	> with body `{"url": "your url", "slug": "your desired slug"}`
- `GET /{slug}` redirect to the original url
- `DELETE /{slug}` remove's the short url
