package main

import (
	"fmt"
	"truenas-monitor/internal/monitor"
)

func main() {
	fmt.Println("TrueNAS Monitor 启动！")

	// 调用获取磁盘健康信息
	monitor.GetDiskHealth()

	// 调用获取存储池状态
	monitor.GetStoragePoolStatus()
}
