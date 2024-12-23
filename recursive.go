package main

import "fmt"

// Struktur untuk menyimpan data daerah
type RegionDemand struct {
    Region string
    Demand int
}

// Quick Sort Rekursif
func quickSortRecursive(data []RegionDemand) []RegionDemand {
    if len(data) <= 1 {
        return data
    }

    pivot := data[len(data)/2].Demand
    left := []RegionDemand{}
    right := []RegionDemand{}
    middle := []RegionDemand{}

    for _, val := range data {
        if val.Demand < pivot {
            left = append(left, val)
        } else if val.Demand > pivot {
            right = append(right, val)
        } else {
            middle = append(middle, val)
        }
    }

    left = quickSortRecursive(left)
    right = quickSortRecursive(right)

    return append(append(left, middle...), right...)
}

func main() {
    // Data simulasi
    data := []RegionDemand{
        {"Jakarta", 1500},g
        {"Bandung", 1200},
        {"Surabaya", 1800},
        {"Yogyakarta", 800},
        {"Medan", 1400},
    }

    // Pengurutan
    sortedData := quickSortRecursive(data)

    // Output hasil
    for _, region := range sortedData {
        fmt.Printf("Region: %s, Demand: %d\n", region.Region, region.Demand)
    }
}