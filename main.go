package main

import (
	"github.com/seonjin85/seonjin85coin/explorer"
	"github.com/seonjin85/seonjin85coin/rest"
)

func main() {
	go explorer.Start(4000)
	rest.Start(3000)
}
