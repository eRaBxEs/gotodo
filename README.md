# gotodo

A simple todo app written in Go using go modules and VueJS

# Technical details

- Used go-pg as the Go Object Relational Mapping
- Used Echo as the minimalist framework to easily create and establish RESTful APIs' call
- Used Postgres as the database
- Used go-playground validator for validation purpose
- Used Uber zap for logging purpose
- Used sane-go library for managing the configuration file

# Usage

1. Clone this project
2. After cloning this project
3. Create the required database schema using the 'app/migrations' folder
4. Navigate into the 'app/bin' folder and run the executable

# App api features

- Get all tasks
- Create a task
- Delete a task by id
