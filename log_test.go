package log

import (
	"fmt"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	for i := 0; i < 10; i++ {
		l := NewDefaultLog("test" + fmt.Sprint(i))
		for i := 0; i < 10; i++ {
			l.Info.Println("hello", i)
			l.Warn.Println("hello", i)
			l.Error.Println("hello", i)
			l.Debug.Println("hello", i)
		}
		time.Sleep(time.Second)
	}
}
