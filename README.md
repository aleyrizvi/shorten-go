# shorten - an url shortening library

Shorten helps you easily convert long urls to short ones with multiple service providers. 


## Install
`go get -u github.com/go-shorten/shorten`

## Example usage

```go
package main

import (
	"fmt"
	"os"

	bitly "github.com/go-shorten/bit.ly"
	"github.com/go-shorten/shorten"
)

func main() {
	btly := bitly.New("MY_BITLY_KEY")
	sh := shorten.New()

	sh.AddProvider("bitly", btly)

	url, err := sh.Shorten(&shorten.UrlParams{
		Destination: "https://github.com/go-shorten/shorten",
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	fmt.Println(url)
}

```
---
## Usage with multiple providers
```go
package main

import (
	"fmt"
	"os"

	bitly "github.com/go-shorten/bit.ly"
	"github.com/go-shorten/shorten"
	...
)

func main() {
	btly := bitly.New("MY_BITLY_KEY")
	rebrandly := rebrandly.New(...)
	sh := shorten.New()

	sh.AddProvider("bitly", btly)
	sh.AddProvider("rebrandly", btly)

	url, err := sh.With("rebrandly").Shorten(&shorten.UrlParams{
		Destination: "https://github.com/go-shorten/shorten",
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	fmt.Println(url)
}

```

---

## Features
- Good test coverage
- Multiple providers
- Custom URL shortener with redis (coming soon)

---

## Providers
Currently we are supporting bit.ly at the moment but other providers will be available really soon.

Following is the list of providers in pipeline. Feel free to create a pull request if you like to contribute or reach out if your required provider is not in list.

- [X] [bit.ly](https://bit.ly)
- [ ] [linkb.ee](https://linkb.ee/)
- [ ] [rebrandly](https://www.rebrandly.com/)
- [ ] [short.io](https://short.io)
- [ ] [tiny.cc](https://tiny.cc)
- [ ] custom-redis (redis provider to generate short url locally)

---
## Contributions
We would love to accept contributions on this project. You can help us by:
- Writing godocs 
- Write github actions to run go tests. I may be writing it pretty soon anyway.
- Add other providers.
- Refactor code and send pull requests.
- Improve tutorial on this readme.

---
## Maintainers

- [Asad Rizvi](https://github.com/aleyrizvi) 