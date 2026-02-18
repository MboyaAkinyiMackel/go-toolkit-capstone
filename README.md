# go-toolkit-capstone
# Go Toolkit Capstone Project

## Prompt-Powered Kickstart: Beginner Toolkit for Go

## Project Overview

This toolkit introduces the Go programming language from a beginner’s perspective.
It demonstrates how to install Go, create a simple runnable project, and extend it into a small REST API with a themed Hello World response and simple UI animation.

The project also documents how generative AI was used to accelerate learning, troubleshooting, and implementation.

## Key Features
* Beginner-friendly Go setup guide
* Simple runnable Hello World project
* Small REST API using Go’s `net/http` package
* Themed Hello World web output
* Basic UI animation served from Go backend
* Documentation of AI-assisted learning process

## Project Structure
go-toolkit-capstone/

├── main.go        → Main Go application (API + Hello World)

├── go.mod         → Go module configuration

└── README.md      → Project documentation

## Technology Used
* Go programming language
* Standard Go `net/http` library
* Command Prompt / Terminal
* Git and GitHub for version control

## Getting Started

### 1. Verify Go Installation
Open terminal:

```
go version
```

Expected output example:

```
go version go1.xx windows/amd64
```
### 2. Clone the Repository

```
git clone https://github.com/MboyaAkinyiMackel/go-toolkit-capstone.git
cd go-toolkit-capstone
```
### 3. Run the Project

```
go run main.go
```

You should see:

```
Server running at http://localhost:8080/hello
```
### 4. View the Application

Open browser:

```
http://localhost:8080/hello
```
You will see:

* A themed Hello World greeting
* Animated text
* Output served from a Go REST endpoint

## API Endpoint

### GET `/hello`
Returns:

* A themed greeting page
* Basic animation
* Demonstrates Go web server capability

This shows how Go can serve both API responses and simple UI output.

## Bonus Features Implemented
To extend beyond a basic Hello World:

* Created a simple REST API using Go
* Served an HTML themed greeting page
* Added lightweight CSS animation
* Demonstrated backend + frontend integration

This highlights real-world backend potential using Go.

## AI Prompt Usage
Generative AI assisted in:

* Installing and configuring Go
* Troubleshooting PATH issues
* Creating the initial Hello World code
* Debugging Git and file extension issues
* Designing the REST API bonus feature

This significantly reduced research time and improved clarity during development.

## Challenges Encountered & Fixes

**Go command not recognized**
* Cause: PATH not configured
* Fix: Added Go installation directory to system PATH and restarted terminal

**File saved as `.txt` instead of `.go`**
* Enabled file extensions in Windows
* Renamed file correctly

**Git push rejected**
```
git pull origin main --allow-unrelated-histories
git push
```
Resolved remote/local repository mismatch.

## References

* Go Official Documentation: https://go.dev/doc/
* Go net/http Package Guide: https://pkg.go.dev/net/http
* GitHub Documentation: https://docs.github.com/
* Online tutorials and AI-assisted learning sessions during development

## Purpose of This Repository

This repository serves as:
* A beginner-friendly Go starter toolkit
* Documentation of AI-assisted learning
* Capstone project submission
* Reference material for new Go learners
 
## Conclusion
This project demonstrates how generative AI can accelerate learning a new programming language.
By combining guided prompts, hands-on experimentation, and documentation, a functional beginner toolkit for Go was successfully created.
