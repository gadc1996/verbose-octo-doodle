package main

import (
	"github.com/gadc1996/verbose-octo-doodle/matrix"
)

var Items = [][2]string{
	{"AA1", "AM6"},
	{"BA1", "BJ6"},
	{"CA1", "CD6"},
	{"CE1", "CH9"},
	{"DA1", "DG9"},
	{"EA1", "EF9"},
	{"FA1", "FF9"},
	{"GA1", "GG10"},
	{"HA1", "HI4"},
	{"IA1", "II4"},
	{"JA1", "JK4"},
}

func main() {
	for _, item := range Items {
		m := matrix.New(item)
		m.Generate()
	}
}
