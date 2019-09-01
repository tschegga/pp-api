package data

// Status represents the status of the microservice.
// It currently consists of it's name and it's version.
type Status struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
