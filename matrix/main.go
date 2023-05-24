package matrix

import (
	"fmt"
	"strconv"
)

type Matrix struct {
	Start string
	End   string
}

func New(coordinates [2]string) *Matrix {
	return &Matrix{
		Start: coordinates[0],
		End:   coordinates[1],
	}
}

func (m *Matrix) Generate() {
	for row := 0; row <= m.height(); row++ {
		m.print(row)
	}
}

func (m *Matrix) prefix() string {
	return string(m.Start[0])
}

func (m *Matrix) height() int {
	var (
		start, _ = strconv.ParseInt(m.Start[2:], 10, 32)
		end, _   = strconv.ParseInt(m.End[2:], 10, 32)
	)
	return int(end - start)
}

func (m *Matrix) width() int {
	return int(m.End[1] - m.Start[1])
}

func (m *Matrix) print(row int) {
	for column := 0; column <= m.width(); column++ {
		fmt.Printf("%v%v%v ", m.prefix(), string(m.Start[0]+byte(column)), row+1)
	}
	fmt.Print("\n")
}
