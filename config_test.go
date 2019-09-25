package main

import (
	"fmt"
	"testing"

	"main/bin"
)

func TestBinConfig(t *testing.T) {
	config.Init(nil)
	fmt.Println(config.C)
}
