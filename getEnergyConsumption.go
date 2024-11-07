package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

func calculatePowerConsumption(processID int32) (float64, error) {
	// Get the process
	proc, err := process.NewProcess(processID)
	if err != nil {
		return 0, err
	}

	// Get the CPU usage of the process over a short interval (1 second)
	cpuPercent, err := proc.Percent(time.Second)
	if err != nil {
		return 0, err
	}

	// Get the CPU frequency
	cpuInfos, err := cpu.Info()
	if err != nil || len(cpuInfos) == 0 {
		return 0, fmt.Errorf("unable to get CPU information")
	}
	cpuFreq := cpuInfos[0].Mhz

	numCores := 6

	// Calculate the instantaneous power consumption in Watts
	powerConsumption := (cpuPercent / 100) * cpuFreq * float64(numCores) / 1000
	return powerConsumption, nil
}

func main() {

	fmt.Print("Enter the process ID: ")
	var processID int32
	_, err := fmt.Scanf("%d", &processID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Handle OS signals for graceful exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var cumulativeEnergy float64

	go func() {
		for {
			powerConsumption, err := calculatePowerConsumption(processID)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			energyConsumption := powerConsumption * (1.0 / 3600.0) // in Wh for 1 second

			cumulativeEnergy += energyConsumption

			fmt.Printf("Power consumption of process %d: %.5f W\n", processID, powerConsumption)
			fmt.Printf("Cumulative energy consumption of process %d until now: %.8f Wh\n", processID, cumulativeEnergy)

			time.Sleep(1 * time.Second)
		}
	}()

	<-sigChan
	fmt.Println("Exiting...")
}
