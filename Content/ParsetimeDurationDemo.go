package Content

import (
	"fmt"
	"time"
)

func DemoParsetimeDuration() {
	// Chuỗi thời gian có đơn vị
	durationStr := "1h30m15s"

	// Chuyển đổi chuỗi thời gian thành giá trị kiểu time.Duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("Không thể chuyển đổi chuỗi thời gian.")
		return
	}

	// In ra giá trị kiểu time.Duration đã chuyển đổi
	fmt.Println("Thời gian đã chuyển đổi:", duration)
	fmt.Println("Giờ:", duration.Hours())
	fmt.Println("Phút:", duration.Minutes())
	fmt.Println("Giây:", duration.Seconds())
}
