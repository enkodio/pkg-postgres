package postgres

import (
	"fmt"
	"strings"
)

type Placeholder int

func (p *Placeholder) Next() string {
	*p = *p + 1
	return fmt.Sprintf("$%v", *p)
}

type Placeholders []string

func (p Placeholders) GenerateByNumber(number int) string {
	placeholders := strings.Builder{}
	for n := 1; n < number; n++ {
		placeholders.WriteString(fmt.Sprintf("$%v,", n))
	}
	placeholders.WriteString(fmt.Sprintf("$%v", number))
	return placeholders.String()
}
