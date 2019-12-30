This project is reservation system utilizing CQRS and Event Sourcing.

Backend is written in Go.

MongoDB is currently used for event store and Redis for message bus implementation.

Frontend is SPA, written using React and TypeScript.

Status
==

Currently focus is on figuring out architecture, while also implementing basic features, in order to ensure best development workflow and decoupling components as much as possible.
Only basic validation and security features are implemented.

Annotated directory/package structure
==
```API Blueprint
├── app
│   ├── aggregates - Aggregates and logic for restoring them from events. Any service can use any aggregate
│   ├── server - Server initialization logic, middlewares
│   ├── services - Contains individual services. Service boundries can only be crossed by dispatching commands
│   │   ├── reservations - Service package example. Each service can contain following files by convention
│   │   │   ├── command_handlers.go
│   │   │   ├── command_validators.go
│   │   │   ├── http - Each service defines it's own routes
│   │   │   │   └── endpoints.go
│   │   │   ├── query.go
│   │   │   └── service.go
├── cmd *Entry points for 
│   ├── all - All services in one binary
│   │   └── main.go
│   ├── guests - Service in separate binary
├── frontend - React app
├── go.mod
├── go.sum
└── pkg - Reusable parts of code
    ├── bus - Message bus definition
    ├── event - Event repository definition
    ├── mongo
    │   └── event - MongoDB event repository implementation
    ├── redis
    │   └── bus - Redis message bus implementation
    └── validator - Common validators
```
