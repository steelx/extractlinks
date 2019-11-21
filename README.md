[![Actions Status](https://github.com/steelx/extractlinks/workflows/go/badge.svg)](https://github.com/steelx/extractlinks/actions)

# extractlinks
extractlinks GO package for extracting anchor links from HTML


Extracts all anchor links from a HTML page into an Array of `[]Link`
```
type Link struct {
	Href string
	Text string
}
```

## Install
`go get -u github.com/steelx/extractlinks`


## Example

```
package main

import (
  "fmt"
  "net/http"
  
  "github.com/steelx/extractlinks"
)

func main() {
  resp, _ := http.Get("http://www.youtube.com/JsFunc")
  links, err := extractlinks.All(resp.Body)
  checkErr(err)
  
  fmt.Println(links)
}

```

Output: (... is just to suppress rest of the result)

   `[{/ IN} {//www.youtube.com/upload } {/channel/UCuB4FSBjofpagXnBlHQUocA } ...]`
