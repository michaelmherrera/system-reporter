package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

/*
system_profiler -detailLevel mini -json SPHardwareDataType> ~/output.jso
*/
const bytesToGig = 1024 * 1024 * 1024

func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getCPUStats() {
	cpuStat, err := cpu.Info()
	handle(err)
	fmt.Println("CPU Model:", cpuStat[0].ModelName)
}

func getMemStats() {
	memStat, err := mem.VirtualMemory()
	handle(err)
	fmt.Println("RAM:", memStat.Total/bytesToGig, "GB")
}

func getDiskStats() {
	diskStat, err := disk.Usage("/")
	handle(err)
	fmt.Println("Storage:", diskStat.Total/bytesToGig, "GB")
}

func main() {
	getCPUStats()
	getMemStats()
	getDiskStats()
}
