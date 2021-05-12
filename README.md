# go-seo
A Go tool to provide SEO advice for a specified domain.

## Usage
An example of validating a H1 length can be performed by running:
```
> go run example/example.go https://dant.blog
```
Sample output:
```
SEO advice for: https://dant.blog

H1 currently: 'Dan&#39;s Blog'
H1 status: {false []}
H2 count: 18
H3 count: 7
H4 count: 0

Internal links found: 20
External links found: 15
```

## Improvements
- Add more analyser options and useful advice
- Add ability to validate inbound / outbound links
- Persist results to a database
- Improve memory handling of results
