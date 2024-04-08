# Data-visualization
data visualization based on a random number guessing game 

#Cloning the repo
#Building

$ go mod init mygraph
$ go mod tidy

which downloads all external packages that are required.
If both of them are already existing:
$ go build

should be enought to build the software

#Running

go run main.go diagram.go

#About the Project

    Package Import:
        Imports necessary packages for working with files (os) and generating charts (github.com/wcharczuk/go-chart).

    GeneratePieChart Function:
        Calculates the number of games won and lost based on the total number of games played.
        Calculates the percentage of games won and lost.
        Creates a pie chart using the go-chart library, specifying its width, height, and values (games won and lost).
        Saves the pie chart as a PNG file named games_chart.png.
        Prints a message indicating that the pie chart has been saved.

    Main Function (not shown):
        Likely triggers the execution of the GeneratePieChart function with appropriate game statistics as parameters.
        The GeneratePieChart function may be called with the game statistics obtained from another part of the program, such as after a game session.

    Execution (not shown):
        The main function or another part of the program calls the GeneratePieChart function with the appropriate game statistics (games played and games won) as arguments.
        Upon execution, the GeneratePieChart function calculates the number of games lost and their percentages, creates a pie chart, and saves it as a PNG file. Overall, this code is a reusable function for generating a pie chart representation of game statistics, providing a visual summary of games won and lost.


