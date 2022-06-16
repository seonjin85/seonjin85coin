package main

import (
	"github.com/seonjin85/seonjin85coin/cli"
	"github.com/seonjin85/seonjin85coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
