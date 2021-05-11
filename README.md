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
Link count: 82
```

An example of finding all internal links in a single page can be performed by running:
```
> go run example_recursive/example_recursive.go https://dant.blog
```
Sample output:
```
Internal links found:
> 58 links
> with a search depth of 2
> taking 1.694995831s
Internal and external links found:
> 517 links
> with a search depth of 2
> taking 1.757099677s
```

## Improvements
- Add more analyser options and useful advice
- Add ability to validate inbound / outbound links
- Persist results to a database
- Improve memory handling of results
