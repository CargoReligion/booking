# Booking App

This application is a booking system that allows coaches and students to manage and schedule sessions. It consists of a client application written in Svelte and a separate server application written in Go with a Postgres DB.

## Features

- Coach dashboard for managing available slots
- Student dashboard for booking sessions
- User impersonation for testing different roles
- Session feedback system

## Prerequisites

- Docker
- Docker Compose

## Running the Application

### Client (Svelte App)

To run the client application:

1. Navigate to the client directory: cd path/to/client
   
2. Build and start the Docker container: docker-compose up --build
   
3. Access the application in your web browser at: http://localhost:5173/
   
### Server

To run the server application:

1. Navigate to the server directory: cd path/to/server
   
2. Start the Docker container: docker-compose up api
   
The server will start and be available for the client application to interact with.

## Development

For development purposes, you can run both the client and server simultaneously using Docker Compose. This setup ensures that both applications are running in isolated environments with all necessary dependencies.
