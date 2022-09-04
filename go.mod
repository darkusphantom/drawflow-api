module drawflow-api

go 1.19

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	models v1.0.0
)

replace (
	models v1.0.0 => ./models
)