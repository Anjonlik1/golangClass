package main

import "fmt"

type ControlPanel struct {
	devices []Device
	
}
type Device interface {
	TurnOn() string
	TurnOff() string
	Status() string
}
type SmartBulb struct {
	isOn       bool
	brightness int
}

func (b SmartBulb) TurnOn() string {
	fmt.Println("Turn on")
}
func (t SmartBulb) TurnOff() string {
	fmt.Println("turn on")
}
func (t SmartBulb) Status() string {
	fmt.Println("turn on")
}

type SmartThermostat struct {
	isOn        bool
	temperature float64
}

func (t SmartThermostat) TurnOn() string {
	fmt.Println("turn on")
}
func (t SmartThermostat) TurnOff() string {
	fmt.Println("turn on")
}
func (t SmartThermostat) Status() string {
	fmt.Println("turn on")
}

type SmartLock struct {
	isLocked bool
}

func (l SmartLock) TurnOn() string {
	fmt.Println("status")
}
func (l SmartLock) TurnOff() string {
	fmt.Println("status")
}
func (l SmartLock) Status() string {
	fmt.Println("status")
}

func(cp ControlPanel) TurnOff(){
	
}

func main() {

}
