package params

import (
	"errors"
	"fmt"
	"strings"
)

// Implement flag.Value interface
type Urls []string

func (u *Urls) String() string {
	return fmt.Sprint(*u)
}

func (u *Urls) Set(value string) error {
	if len(*u) > 0 {
		return errors.New("urls flag already set")
	}
	for _, url := range strings.Split(value, ",") {
		if url != "" {
			*u = append(*u, url)
		}
	}
	return nil
}
