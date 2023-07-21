package Content

import (
	"fmt"
	"time"
)

func DemoTimeParse() {
	// Chuỗi thời gian có định dạng
	timeStr := "2023-07-14 15:30:00"

	// Định dạng thời gian tương ứng với chuỗi trên
	layout := "2010-01-01 15:04:05"

	// Chuyển đổi chuỗi thời gian thành giá trị kiểu time.Time
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Không thể chuyển đổi chuỗi thời gian.")
		return
	}

	// In ra giá trị kiểu time.Time đã chuyển đổi
	fmt.Println("Thời gian đã chuyển đổi:", parsedTime)
}
