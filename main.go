package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"

	"github.com/pkg/browser"
	"github.com/shirou/gopsutil/cpu"
)

/*
system_profiler -detailLevel mini -json SPHardwareDataType> ~/output.jso
*/
const bytesToGig = 1024 * 1024 * 1024

var macInfo *MacOSInfo

// MacOSInfo comment to stop go from complaining
type MacOSInfo struct {
	CPUModel      string
	CPUSpecs      string
	Memory        string
	BatteryHealth string
	HDCapacity    string
	SerialNumber  string
	DeviceModel   string
}

func handle(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getCPUSpecs(profileJSON *map[string][](map[string]interface{}), info *MacOSInfo) {
	cpuStat, err := cpu.Info()
	handle(err)
	info.CPUModel = cpuStat[0].ModelName
	info.CPUSpecs = (*profileJSON)["SPHardwareDataType"][0]["cpu_type"].(string)

}

func getMemorySpecs(profileJSON *map[string][](map[string]interface{}), info *MacOSInfo) {
	info.Memory = (*profileJSON)["SPHardwareDataType"][0]["physical_memory"].(string)
}

func getBatterySpecs(profileJSON *map[string][](map[string]interface{}), info *MacOSInfo) {
	health := (*profileJSON)["SPPowerDataType"][0]["sppower_battery_health_info"].(map[string]interface{})
	info.BatteryHealth = health["sppower_battery_health"].(string)
}

func getStorageSpecs(profileJSON *map[string][](map[string]interface{}), info *MacOSInfo) {
	for _, storageDevice := range (*profileJSON)["SPStorageDataType"] {
		if storageDevice["_name"] == "Macintosh HD" {
			// Apples uses 1 billion bytes to a gigabyte, rather than 2^30
			info.HDCapacity = fmt.Sprintf("%d GB", int(storageDevice["size_in_bytes"].(float64)/1e9))
		}
	}
}

func getProductSpecs(profileJSON *map[string][](map[string]interface{}), info *MacOSInfo) {
	hardwareInfo := (*profileJSON)["SPHardwareDataType"][0]
	info.SerialNumber = hardwareInfo["serial_number"].(string)
	info.DeviceModel = hardwareInfo["machine_model"].(string)
}

func profileMac() *MacOSInfo {
	info := new(MacOSInfo)
	cmd := exec.Command("system_profiler", "-json", "SPHardwareDataType", "SPStorageDataType", "SPPowerDataType")
	response, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var profileJSON map[string][](map[string]interface{})

	err = json.Unmarshal(response, &profileJSON)
	handle(err)

	getCPUSpecs(&profileJSON, info)
	getBatterySpecs(&profileJSON, info)
	getStorageSpecs(&profileJSON, info)
	getMemorySpecs(&profileJSON, info)
	getProductSpecs(&profileJSON, info)
	return info
}

func serveUp(info *MacOSInfo) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset("index.html")
		handle(err)
		tmpl, err := template.New("test").Parse(string(data))
		handle(err)

		tmpl.Execute(w, info)
	}

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
	if macInfo == nil {
		// If this is the first time the data endpoint has been
		// hit, retrieve the info
		macInfo = profileMac()
	}
	json.NewEncoder(w).Encode(*macInfo)
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
