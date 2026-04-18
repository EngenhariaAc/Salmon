package main

import "fmt"

func main() {
    distance := 0

    fmt.Println("🐟 Simulating salmon migration...")

    for i := 0; i < 5; i++ {
        distance += 50
        fmt.Printf("Distance traveled: %d km\n", distance)
    }

    fmt.Println("🏁 Salmon reached spawning grounds!")
}
