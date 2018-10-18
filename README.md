# TodoApp 
Todo app 

## Requirements
Build a Go server side application exposing an API for managing a to do list

## Usage
* GET "/" displays welcome page of deployed server
* GET "/todos" displays list of todos
* POST "/todos" adds a new todo item to the list
* DELETE "/todos/:id" deletes the todo item by the id
* PUT "/todos/:id" Updates any edits on a todo item by id

## Deployment

### Heroku
Deployed server [here]("https://to-do-app-fion.herokuapp.com/")

### Docker

#### Clone and get repo
`go get -d github.com/fionwan/todoApp`
`cd $GOPATH/src/github.com/fionwan/todoApp`

#### Build Docker image
`docker build -t todoapp .`

#### Run Docker image
`docker run -it todoapp`

#### Build application
`go build`

#### Run application 
`./todoApp`
