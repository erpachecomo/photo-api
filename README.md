├── cmd
│   └── [service-name]
│       └── main.go
├── internal
│   ├── api
│   │   └── handler.go
│   ├── model
│   │   └── model.go
│   ├── repository
│   │   └── repository.go
│   └── service
│       └── service.go
├── pkg
│   └── [reusable-package]
│       └── [package-files]
├── configs
│   └── config.yaml
├── go.mod
├── go.sum

Project structure

* cmd: Contains the main application entry point.
* internal: Holds the core application logic, not intended for external use.
    * api: Defines API handlers.
    * model: Defines data structures.
    * repository: Handles database interactions.
    * service: Implements business logic.
* pkg: Contains reusable packages.
* configs: Stores configuration files.
* go.mod: and go.sum: Manage dependencies.

This structure promotes modularity, maintainability, and testability in Go REST API projects.
