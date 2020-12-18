package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
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

// func getCPUStats() {
// 	cpuStat, err := cpu.Info()
// 	handle(err)
// 	fmt.Println("CPU Model:", cpuStat[0].ModelName)
// }

// func getMemStats() {
// 	memStat, err := mem.VirtualMemory()
// 	handle(err)
// 	fmt.Println("RAM:", memStat.Total/bytesToGig, "GB")
// }

// func getDiskStats() {
// 	diskStat, err := disk.Usage("/")
// 	handle(err)
// 	fmt.Println("Storage:", diskStat.Total/bytesToGig, "GB")
// }

// func macSystemProfiler() {
// 	cmd := exec.Command("system_profiler", "-detailLevel", "mini", "-json", "SPHardwareDataType")
// 	response, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var testJSON map[string][](map[string]interface{})
// 	err = json.Unmarshal(response, &testJSON)
// 	handle(err)
// 	info := make(map[string]interface{})
// 	hardwareInfo := testJSON["SPHardwareDataType"][0]
// 	info["Machine Model:"] = hardwareInfo["machine_model"]
// 	info["Memory:"] = hardwareInfo["physical_memory"]
// 	info["CPU Type:"] = hardwareInfo["cpu_type"]
// 	for k := range info {
// 		fmt.Println(k, info[k])
// 	}
// }

// func serveUp() {
// 	content := []byte("Hello world!")
// 	tmpfile, err := ioutil.TempFile("", "example")
// 	handle(err)

// 	defer os.Remove(tmpfile.Name()) // clean up
// 	fmt.Println(tmpfile.Name())
// 	if _, err := tmpfile.Write(content); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := tmpfile.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// }

func serveUp() {

	// Use npm package inliner to compile the whole project into a single html doc

	browser.OpenURL("http://localhost:8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	const quote = `
	<!DOCTYPE html>
	<html lang="en">

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<!-- Latest compiled and minified CSS -->
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">

	<!-- jQuery library -->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>

	<!-- Latest compiled JavaScript -->
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>

	<body>
		<h1>Hello world!</h1>
	</body>

	</html>
	
	
	`
	fmt.Fprintf(w, quote)
}

func main() {
	serveUp()
	// macSystemProfiler()
	// getCPUStats()
	// getMemStats()
	// getDiskStats()
}
