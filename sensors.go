package main

import (
    "fmt"

    "github.com/yryz/ds18b20"
)

func main() {
    sensors, err := ds18b20.Sensors()
    if err != nil {
        panic(err)
    }

    fmt.Printf("w1_temp:\n")

    for _, sensor := range sensors {
	i := 1
        t, err := ds18b20.Temperature(sensor)
        if err == nil {
		fmt.Printf("temp%v \n temp%v_output:%.2f \n", i, i, t)
        }
	i =i+1
    }
}
