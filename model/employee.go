package model

import (
	"time"
)

// Employee struct will be the data mapping interface of all employee REST API data
type Employee struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Designation    string    `json:"designation"`
	Department     string    `json:"department"`
	JoiningDate    time.Time `json:"joining_date"`
	Address        string    `json:"address"`
	OfficeLocation string    `json:"office_location"`
	Status         string    `json:"status"`
	EmailID        string    `json:"email"`
	AnnualPackage  float64   `json:"annual_package"`
	PhoneNumber    string    `json:"phone_number"`
}

// HealthCheck is a structure for health check information
type HealthCheck struct {
	Message string `json:"message"`
}

// DetailedHealthCheck is a structure for detailed health check information
type DetailedHealthCheck struct {
	Message string `json:"message"`
	ScyllaDB string `json:"scylla_db"`
	EmployeeAPI string `json:"employee_api"`
}
