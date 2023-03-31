// creating a web application using handlers and mutexes
// /home/greetings/random
// Kevin Figueroa
// 2018117755
package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

// the HomeHanler handles the request to the /home endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" Executing the Home Handler")
	//server to home.html file
	http.ServeFile(w, r, "home.html")

}

// the GreetingsHandler handles the request to the /greetings endpoint
func GreetingsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" Executing the Greetings Handler")
	//gets the current time of the system
	currentTime := time.Now()
	//the default salute will be Good Morning!
	salute := "Good Morning!"
	//if the current time is in between 12 and 7pm say it is good afternoon
	if currentTime.Hour() >= 12 && currentTime.Hour() <= 19 {
		salute = "Good Afteroon"
		//then of the current itme is greater than 7 pm then it is set to good night
	} else if currentTime.Hour() >= 19 {
		salute = "Good Night"
	} //getting its current day
	day := currentTime.Weekday().String()
	weekend := ""
	//if the day is a weekend day outout that it is a weekend
	if day == "Friday" || day == "Saturday" || day == "Sunday" {
		weekend = "Yey! it's the Weekend."
	} else {
		//else its not a weekend
		weekend = ". The weekend will be here before you know it "
	}
	//this should no be the completed greeting message
	greetings := salute + " Today is " + day + " and it's: " + currentTime.Format("15:04") + weekend + ". Have a splendid " + day + "!"
	//write out the greeting message to the response
	w.Write([]byte(greetings))
}

// RandomQuoteHandler handles request to the random endpoint
func RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing the Random Quote Handler")
	//create maps to keep the list of quotes
	quotes := map[int]string{
		1: "There is no such thing as failure. It's a life lesson for future ~ Invajy",
		2: "When you suffer in life , be vigilant. Life is trying to teach a lesson. ~ Invajy",
		3: "Lessons in life will be repeated until they are learned. ~ Frank Sonnenberg",
		4: "The biggest lesson we have to give our children is truth. ~ Goldie Hawn",
		5: "Care about what other people think and you will always be their prisoner. ~ Lao Tzu",
		6: "We must be willing to let go of the life we've planned, so as to have the life that is waiting for us. ~ Joseph Campbell",
	}
	//it generates the random number between 1 and the lenght of the quote map
	rand.Seed(time.Now().UnixNano())
	randomResponse := rand.Intn(len(quotes)) + 1
	//gets the quote and formats it in a sentence
	randomQuote := "Random Quote of the day: " + quotes[randomResponse]
	//wrote the quots to the response
	w.Write([]byte(randomQuote))

}

func main() {
	//create the the new server mux for the http
	mux := http.NewServeMux()
	//hadler functions for the paths
	mux.HandleFunc("/home", HomeHandler)
	log.Println("Starting the Home server on /home")
	mux.HandleFunc("/greetings", GreetingsHandler)
	log.Println("Starting the Greetings server on /greetings")
	mux.HandleFunc("/random", RandomQuoteHandler)
	log.Println("Starting the Random server on /random")
	//creating the http server with the 8181 port number for the server
	server := &http.Server{
		Addr:    ":8181",
		Handler: mux, // setting the handler for the mux server
	}

	err := server.ListenAndServe()
	//if server fails to start then log an error message and terminate the program
	log.Fatal(err)

}
