package szamlazzhu

import "fmt"

type SzamlazzhuError struct {
	c int
	e string
}

func (e *SzamlazzhuError) Error() string {
	return fmt.Sprintf("szamlazz.hu error %d: %s", e.c, e.e)
}
