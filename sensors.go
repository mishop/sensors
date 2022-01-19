package main

import (
    "fmt"
    "errors"
    "io/ioutil"
    "strconv"
    "strings"
)

var ErrReadSensor = errors.New("failed to read sensor temperature")

// Sensors get all connected sensor IDs as array
func Sensors() ([]string, error) {
    data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
    if err != nil {
        return nil, err
    }

    sensors := strings.Split(string(data), "\n")
    if len(sensors) > 0 {
        sensors = sensors[:len(sensors)-1]
    }

    return sensors, nil
}

// Temperature get the temperature of a given sensor
func Temperature(sensor string) (float64, error) {
    data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + sensor + "/w1_slave")
    if err != nil {
        return 0.0, ErrReadSensor
    }

    raw := string(data)

    if !strings.Contains(raw, " YES") {
        return 0.0, ErrReadSensor
    }

    i := strings.LastIndex(raw, "t=")
    if i == -1 {
        return 0.0, ErrReadSensor
    }

    c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
    if err != nil {
        return 0.0, ErrReadSensor
    }

    return c / 1000.0, nil
}


func main() {
    sensors, err := Sensors()
    if err != nil {
        panic(err)
    }

    fmt.Printf("w1_temp:\n")

    for i, sensor := range sensors {
        t, err := Temperature(sensor)
        if err == nil {
		fmt.Printf("temp%v: \n  temp%v_input: %.2f \n", i, i, t)
        }
    }
}
