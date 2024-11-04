package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	sensorChannel := make(chan int64)
	meanChannel := make(chan int64)

	go randomCounter(sensorChannel)
	go arithmeticMean(sensorChannel, meanChannel)

	for val := range meanChannel {
		fmt.Printf("Arithmetic mean %d\n", val)
	}
}

func randomCounter(channel chan int64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	defer close(channel)

	timeout := time.After(time.Minute)

	for {
		select {
		case <-timeout:
			return
		case <-ticker.C:
			data, _ := rand.Int(rand.Reader, big.NewInt(100))
			channel <- data.Int64()
		}
	}
}

func arithmeticMean(sensor chan int64, mean chan int64) {
	var summ int64
	defer close(mean)
	for {
		summ = 0
		for i := 0; i < 10; i++ {
			val, ok := <-sensor
			summ += val
			if !ok {
				if i == 9 {
					mean <- (summ / 10)
				}
				return
			}
		}
		mean <- (summ / 10)
	}
}
