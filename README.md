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
H1 status: The H1 is too short, aim for 20 characters minimum.
H2 count: 18
H3 count: 7
H4 count: 0
```

## Improvements
- Add more analyser options and useful advice
- Add ability to scan entire domains concurrently gathering information on a per webpage basis
- Add ability to validate inbound / outbound links
- Persist results to a database
