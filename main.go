package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Bicycle struct {
	Brand          string
	Model          string
	Price          int
	Size           int
	Rating         float64
	SafetyFeatures []string
}

func main() {
	// Initialize a list of bicycle options
	bicycles := []Bicycle{
		{"Hero", "RiderX", 7000, 16, 4.5, []string{"Strong Brakes", "Lightweight Frame"}},
		{"Btwin", "Speedster", 10000, 18, 4.7, []string{"Shock Absorbers", "Comfortable Seat"}},
		{"Firefox", "Turbo", 12000, 20, 4.6, []string{"Strong Brakes", "Anti-Skid Tires"}},
		{"Atlas", "Explorer", 8000, 16, 4.3, []string{"Training Wheels", "Durable Frame"}},
	}

	fmt.Println("Available Bicycles:")
	for _, bike := range bicycles {
		fmt.Printf("Brand: %s, Model: %s, Price: ₹%d, Size: %d-inch, Rating: %.1f\n",
			bike.Brand, bike.Model, bike.Price, bike.Size, bike.Rating)
	}

	// Simulate user needs
	budget := 10000
	desiredSize := 16
	fmt.Printf("\nFiltering options within budget ₹%d and size %d-inch...\n", budget, desiredSize)

	var shortlisted []Bicycle
	for _, bike := range bicycles {
		if bike.Price <= budget && bike.Size == desiredSize {
			shortlisted = append(shortlisted, bike)
		}
	}

	if len(shortlisted) == 0 {
		fmt.Println("No suitable bicycles found within budget and size requirements.")
		return
	}

	// Simulate selecting the best based on reviews
	fmt.Println("\nShortlisted Bicycles:")
	bestBike := shortlisted[0]
	for _, bike := range shortlisted {
		fmt.Printf("Brand: %s, Model: %s, Price: ₹%d, Rating: %.1f\n",
			bike.Brand, bike.Model, bike.Price, bike.Rating)
		if bike.Rating > bestBike.Rating {
			bestBike = bike
		}
	}

	fmt.Printf("\nFinal Selection: %s %s (₹%d) with %.1f rating\n", bestBike.Brand, bestBike.Model, bestBike.Price, bestBike.Rating)

	// Simulate customer feedback
	rand.Seed(time.Now().UnixNano())
	feedback := []string{
		"Great quality and safety features!",
		"Good value for money.",
		"Lightweight and easy to ride.",
		"Comfortable seat but could have better brakes.",
	}
	fmt.Printf("Customer Feedback: %s\n", feedback[rand.Intn(len(feedback))])
}
