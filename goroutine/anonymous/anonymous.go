package anonymous

import (
	"fmt"
	"time"
)

func Anonymous() {
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Hello from anonymous goroutine")
		}
	}()
	time.Sleep(600 * time.Millisecond) // 메인 함수가 종료되지 않도록 잠시 대기
	fmt.Println("Anonymous function done")
}
