package pkg1

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	time.Sleep(5 * time.Second)
	fmt.Println("pkg1")
}
