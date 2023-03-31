package main

import (
	"context"
	"fmt"
)

func main() {

	contextParent := context.Background()

	ctxSatu := context.WithValue(contextParent, "key1", "Hello World")
	ctxDua  := context.WithValue(ctxSatu, "key2", "Hello girls")
	ctxTiga := context.WithValue(ctxDua, "key3", "Hello boys")
	ctxEmpat := context.WithValue(contextParent, "key4", "Hello Children")
	ctxLima := context.WithValue(ctxSatu, "key5", "Hello Later")

	fmt.Println(ctxLima.Value("key5"))
	fmt.Println(ctxLima.Value("key4"))
	fmt.Println(ctxLima.Value("key3"))
	fmt.Println(ctxLima.Value("key2"))
	fmt.Println(ctxLima.Value("key1"))
	fmt.Println(ctxTiga.Value("key1"))
	fmt.Println(ctxEmpat.Value("key1"))
}