# Simple Web Server using Go and Gin

## Overview
This is a simple web server built with the Gin web framework. It provides three routes that return static responses.

### Routes
1. **/name**  
   - **Method**: GET  
   - **Response**:Hello, I am Afomia!.  

2. **/hobby**  
   - **Method**: GET  
   - **Response**: I love exploring new places!.  

3. **/dream**  
   - **Method**: GET  
   - **Response**: I dream of becoming a Backend Developer!. 

## Getting Started

### Prerequisites
- Go 1.16+ installed
- A terminal or command-line interface

### Installation

1. Clone the repository:
  
   git clone https://github.com/Afomiat/simple-web-server.git
   cd simple-web-server

2. Install dependencies:
    go mod tidy

3. Run the server:
    go run main.go

4. Access the routes locally at:
    http://localhost:8080/name
    http://localhost:8080/hobby
    http://localhost:8080/dream





