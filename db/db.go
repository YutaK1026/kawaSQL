package db

import (
	"github.com/yutak1026/kawasql/query"
	"fmt"
	// "log"
	// "os"
	// "os/signal"
	// "strings"
)

func Execute(q string) {
	tokenizer := query.NewTokenizer(q)
	fmt.Print(tokenizer)
}