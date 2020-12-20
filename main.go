package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/pkg/browser"
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

func getMacSystemProfile() {
	cmd := exec.Command("system_profiler", "-detailLevel", "mini", "-json", "SPHardwareDataType")
	response, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var testJSON map[string][](map[string]interface{})
	err = json.Unmarshal(response, &testJSON)
	handle(err)
	info := make(map[string]interface{})
	hardwareInfo := testJSON["SPHardwareDataType"][0]
	info["Machine Model:"] = hardwareInfo["machine_model"]
	info["Memory:"] = hardwareInfo["physical_memory"]
	info["CPU Type:"] = hardwareInfo["cpu_type"]
}

func serveUp() {

	// Use npm package inliner to compile the whole project into a single html doc
	browser.OpenURL("http://localhost:8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := Asset("index.html")
	handle(err)
	fmt.Fprintf(w, string(data))
}

func dataEndpoint(w http.ResponseWriter, r *http.Request) {
	macSystemProfiler
	json.NewEncoder(w).Encode(arr)
}

func restCalls() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/data.json", dataEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	restCalls()
	//serveUp()
	// macSystemProfiler()
	// getCPUStats()
	// getMemStats()
	// getDiskStats()
}
