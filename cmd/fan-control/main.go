package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/mdlayher/lmsensors"
	"github.com/shirou/gopsutil/load"
	"go.bug.st/serial"
)

// readSensors run a scan an groups sensors values as average for each device
func readSensors(scanner *lmsensors.Scanner) map[string]float64 {
	counts := make(map[string]int)
	sums := make(map[string]float64)

	devices, err := scanner.Scan()
	if err == nil {
		for _, device := range devices {
			w := strings.Split(device.Name, "-")
			if len(w) > 1 {
				if _, err := strconv.Atoi(w[len(w)-1]); err == nil {
					device.Name = strings.Join(w[:len(w)-1], "-")
				}
			}
			for _, sensor := range device.Sensors {
				switch s := sensor.(type) {
				case *lmsensors.FanSensor:
					counts[device.Name]++
					sums[device.Name] += float64(s.Input)

				case *lmsensors.TemperatureSensor:
					counts[device.Name]++
					sums[device.Name] += s.Input

				case *lmsensors.PowerSensor:
					counts[device.Name]++
					sums[device.Name] += float64(s.Average)
				}

			}
		}
	}
	readings := make(map[string]float64)
	for device, count := range counts {
		readings[device] = sums[device] / float64(count)

	}
	return readings
}

func readCPU() float64 {
	avgStat, err := load.Avg()
	if err == nil {
		return avgStat.Load1
	}

	return 0
}

func runReadingSensors(ctx context.Context, interval time.Duration, outChan chan map[string]float64) {
	scanner := lmsensors.New()

	ticker := time.Tick(interval)
	defer close(outChan)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker:
			// Read
			readings := readSensors(scanner)
			readings["cpu"] = readCPU()
			outChan <- readings

		}
	}

}

func main() {
	log.Println("fan-control starting")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	// Print the list of detected ports
	for _, port := range ports {
		log.Printf("Found port: %v\n", port)
	}
	sensorsChan := make(chan map[string]float64)
	go runReadingSensors(ctx, time.Second*5, sensorsChan)

	for {
		select {
		case <-ctx.Done():
			log.Print("fan-control exiting")
			return
		case readings := <-sensorsChan:
			log.Println(readings)
		}
	}

}
