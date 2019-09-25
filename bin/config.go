package config

import "time"

//go:generate ./vc
type Config struct {
	A    int `default:"2" desc:"一个整数a"`
	Nest struct {
		B time.Duration
	}
}
