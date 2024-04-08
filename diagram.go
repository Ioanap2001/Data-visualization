package main

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"
)

func GeneratePieChart(gamesPlayed, gamesWon int) error {
	// Calculate games lost
	gamesLost := gamesPlayed - gamesWon

	// Calculate percentages
	wonPercentage := float64(gamesWon) / float64(gamesPlayed) * 100
	lostPercentage := float64(gamesLost) / float64(gamesPlayed) * 100

	// Create a pie chart
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: float64(gamesWon), Label: fmt.Sprintf("Games Won (%.2f%%)", wonPercentage)},
			{Value: float64(gamesLost), Label: fmt.Sprintf("Games Lost (%.2f%%)", lostPercentage)},
		},
	}

	// Save the chart as a PNG file
	file, err := os.Create("games_chart.png")
	if err != nil {
		return err
	}
	defer file.Close()

	err = pie.Render(chart.PNG, file)
	if err != nil {
		return err
	}

	fmt.Println("Pie chart saved as 'games_chart.png'")
	return nil
}
