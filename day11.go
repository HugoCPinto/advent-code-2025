package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Device struct {
	Id	string
	Outputs	[]string
}

func readInput(path string) []Device {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	devices := []Device{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		deviceId := strings.Split(line, ":")[0]
		outputsS := strings.Split((strings.Split(line, ":")[1]), " ")
		
		outputs := []string{}
		for _, output := range outputsS{
			outputs = append(outputs, output)
		}

		device := Device{
			Id: deviceId,
			Outputs: outputs,
		}

		devices = append(devices, device)

	}

	return devices
}

func recursiveFindPath(devices []Device, deviceId string, deviceOutput string) int {
	
	if deviceOutput == "out" {
		return 1
	}

	total := 0
	for _,d := range devices{
		if d.Id == deviceOutput {
			for _,o := range d.Outputs {
				total += recursiveFindPath(devices, d.Id, o)
			}			
		}
	}
	return total
}

func processInput(devices []Device) int {

	total := 0
	for _,device := range devices {
		if device.Id == "you" {
			for _,output := range device.Outputs {
				total += recursiveFindPath(devices, device.Id, output)
			}
		}
	}

	return total
}

func main(){
	fmt.Println("Day 11")

	devices := readInput("inputs/day11-1.txt")
	//fmt.Println(devices)
	total := processInput(devices)
	fmt.Println("Total paths: ", total)
}