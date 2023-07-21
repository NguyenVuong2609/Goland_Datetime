package Content

import (
	"fmt"
	"time"
)

func Demo3() {
	fmt.Println("-------------------------- TimeSince ----------------------------")
	// Ghi nhận thời điểm bắt đầu
	startTime := time.Now()

	// Mô phỏng một quá trình thực hiện công việc mất thời gian
	time.Sleep(2 * time.Second)

	// Tính thời gian đã trôi qua kể từ thời điểm bắt đầu
	elapsedTime := time.Since(startTime)

	fmt.Printf("Thời gian đã trôi qua: %s\n", elapsedTime)
}
