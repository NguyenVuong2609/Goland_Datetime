package Content

import (
	"fmt"
	"time"
)

func DemoTimeTick() {
	// Tạo một kênh (`chan time.Time`) bằng cách sử dụng time.Tick
	interval := 3 * time.Second
	ticker := time.Tick(interval)

	// Sử dụng vòng lặp for và select để lắng nghe từng giá trị được gửi qua kênh
	for {
		select {
		case currentTime := <-ticker:
			// Khi nhận được giá trị từ kênh, in ra thời gian hiện tại
			fmt.Println("Thời gian hiện tại:", currentTime)
		}
	}
}
