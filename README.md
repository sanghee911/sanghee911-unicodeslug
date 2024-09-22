# Unicode Slug
Unicode Slug provides functions to create Unicode slugs and verify slugs with Unicode characters in it.

## Install
```shell
go get -u github.com/sanghee911/unicodeslug
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/sanghee911/unicodeslug"
)

func main() {
	unicodeslug.Slugify(" 파이썬 기초강좌 #1 ")
	// returns "파이썬-기초강좌-1"
	
	unicodeslug.Slugify("😄 이모티콘")
	// returns "이모티콘"

	unicodeslug.ValidateSlug("파이썬 기초 강좌 #1")
	// returns false
	
	unicodeslug.ValidateSlug("파이썬-기초-강좌-#1")
	// returns false
	
	unicodeslug.ValidateSlug("파이썬-기초-강좌-1")
	// returns true
}
```