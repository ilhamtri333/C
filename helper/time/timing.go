package timing

import (
	"fmt"
	"time"
)

func TimingCheck(str string) bool {
	_, err := time.Parse("2006-01-02", str)
	fmt.Println(err, str)
	if err != nil {
		return false
	}
	return true
}
