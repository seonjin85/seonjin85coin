package main

import (
	"github.com/seonjin85/seonjin85coin/explorer"
	"github.com/seonjin85/seonjin85coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
