package main

import (
	"encoding/json"
	"fmt"
	"log"

	"os/exec"

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
	cmd := exec.Command("system_profiler", "-detailLevel", "mini", "-json", "SPHardwareDataType")
	response, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var testJSON map[string][](map[string]interface{}) //map[string](map[string][]string)
	err = json.Unmarshal(response, &testJSON)
	handle(err)
	info := make(map[string]interface{})
	hardwareInfo := testJSON["SPHardwareDataType"][0]
	info["Machine Model:"] = hardwareInfo["machine_model"]
	info["Memory:"] = hardwareInfo["physical_memory"]
	info["CPU Type:"] = hardwareInfo["cpu_type"]
	for k := range info {
		fmt.Println(k, info[k])
	}
	// fmt.Printf(testJSON)
	// fmt.Printf("%T", testJSON)
	// getCPUStats()
	// getMemStats()
	// getDiskStats()
}
