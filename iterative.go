package main

import (
    "fmt"
    "time"
)

// Struktur untuk menyimpan data daerah
type RegionDemand struct {
    Region string
    Demand int
}

// Quick Sort Iteratif
func quickSortIterative(data []RegionDemand) []RegionDemand {
    stack := []struct {
        start, end int
    }{{0, len(data) - 1}}

    for len(stack) > 0 {
        n := len(stack) - 1
        start, end := stack[n].start, stack[n].end
        stack = stack[:n]

        if start >= end {
            continue
        }

        pivot := data[end].Demand
        low := start

        for i := start; i < end; i++ {
            if data[i].Demand < pivot {
                data[i], data[low] = data[low], data[i]
                low++
            }
        }

        data[low], data[end] = data[end], data[low]

        stack = append(stack, struct{ start, end int }{start, low - 1})
        stack = append(stack, struct{ start, end int }{low + 1, end})
    }

    return data
}

func main() {
    // Data simulasi
    data := []RegionDemand{
        {"Jakarta", 1500},
        {"Bandung", 1200},
        {"Surabaya", 1800},
        {"Yogyakarta", 800},
        {"Medan", 1400},
    }

    // Mengukur waktu eksekusi
    startTime := time.Now()

    // Pengurutan
    sortedData := quickSortIterative(data)

    duration := time.Since(startTime)

    // Output hasil
    for _, region := range sortedData {
        fmt.Printf("Region: %s, Demand: %d\n", region.Region, region.Demand)
    }

    fmt.Printf("Execution Time: %s\n", duration.Nanoseconds())
}
