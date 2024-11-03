package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	sensorChannel := make(chan int64)
	stopChannel := make(chan bool)
	meanChannel := make(chan int64)
	doneChannel := make(chan bool)

	go randomCounter(sensorChannel, stopChannel)
	go arithmeticMean(sensorChannel, meanChannel, doneChannel)

LOOP:
	for {
		select {
		case <-stopChannel:
			close(sensorChannel)
		case val := <-meanChannel:
			fmt.Printf("Arithmetic mean %d\n", val)
		case <-doneChannel:
			break LOOP
		}
	}
}

func randomCounter(channel chan int64, stop chan bool) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	timeout := time.After(time.Minute)

	for {
		select {
		case <-timeout:
			stop <- true
			return
		case <-ticker.C:
			data, _ := rand.Int(rand.Reader, big.NewInt(100))
			channel <- data.Int64()
		}
	}
}

func arithmeticMean(sensor chan int64, mean chan int64, done chan bool) {
	var summ int64
	for {
		summ = 0
		for i := 0; i < 10; i++ {
			val, ok := <-sensor
			summ += val
			if !ok {
				if i == 9 {
					mean <- (summ / 10)
				}
				done <- true
			}
		}
		mean <- (summ / 10)
	}
}
