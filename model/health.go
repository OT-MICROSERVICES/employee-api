package model

// DetailedHealthCheck is a structure for detailed health check information
type DetailedHealthCheck struct {
	Message     string `json:"message"`
	ScyllaDB    string `json:"scylla_db"`
	EmployeeAPI string `json:"employee_api"`
	Redis       string `json:"redis"`
}
