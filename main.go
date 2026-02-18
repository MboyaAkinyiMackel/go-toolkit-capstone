package main

import (
	"fmt"
	"net/http"
	"time"
)

// API handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
	<head>
		<title>Hello Go</title>
		<style>
			body {
				font-family: Arial;
				text-align: center;
				margin-top: 50px;
			}
			h1 {
				animation: fade 2s infinite alternate;
			}
			@keyframes fade {
				from {opacity: 0.3;}
				to {opacity: 1;}
			}
		</style>
	</head>
	<body>
		<h1>✨ Hello from Go Toolkit API ✨</h1>
		<p>It has been a blast doing this project 🎉</p>
        <p>Danke</p>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running at http://localhost:8080/hello")

	time.Sleep(1 * time.Second)

	http.ListenAndServe(":8080", nil)
}

