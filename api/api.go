package api

import (
	"employee-api/client"
	"employee-api/config"
	"employee-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	redisEnabled = config.ReadConfigAndProperty().Redis.Enabled
)

// @Summary ReadEmployeesDesignation is a method to read all employee designation
// @Schemes http
// @Description Read all employee location data from database
// @Tags employee
// @Accept json
// @Produce json
// @Success 200 {object} model.Designation
// @Router /search/designation [get]
// ReadEmployeesDesignation is a method to read all employee location
func ReadEmployeesDesignation(c *gin.Context) {
	var employee model.Employee
	var designationResponse model.Designation
	var redisError error
	var redisData string
	if redisEnabled {
		redisData, redisError = client.CreateRedisClient().HGet(ctx, "employee", "designation").Result()
		if redisError != nil {
			logrus.Warnf("Unable to read data from Redis %v", redisError)
		}
		_ = json.Unmarshal([]byte(redisData), &designationResponse)
		if redisError == nil {
			logrus.Infof("Successfully fetched the data for designation from the Redis")
			c.JSON(http.StatusOK, designationResponse)
			return
		}
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	designation := make(map[string]int)
	data := scyllaClient.Query("SELECT designation FROM employee_info").Iter()
	for data.Scan(&employee.Designation) {
		_, exist := designation[employee.Designation]
		if exist {
			designation[employee.Designation] += 1
		} else {
			designation[employee.Designation] = 1
		}
	}
	jsonData, err := json.Marshal(designation)
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	if redisEnabled && redisError == redis.Nil {
		writeinRedis("designation", string(jsonData))
	}
	logrus.Infof("Successfully fetched the data for all designation from the ScyllaDB")
	json.Unmarshal(jsonData, &designationResponse)
	c.JSON(http.StatusOK, designationResponse)
	return
}

// @Summary ReadEmployeesLocation is a method to read all employee location
// @Schemes http
// @Description Read all employee location data from database
// @Tags employee
// @Accept json
// @Produce json
// @Success 200 {object} model.Location
// @Router /search/location [get]
// ReadEmployeesLocation is a method to read all employee location
func ReadEmployeesLocation(c *gin.Context) {
	var employee model.Employee
	var locationResponse model.Location
	var redisError error
	var redisData string
	if redisEnabled {
		redisData, redisError = client.CreateRedisClient().HGet(ctx, "employee", "location").Result()
		if redisError != nil {
			logrus.Warnf("Unable to read data from Redis %v", redisError)
		}
		json.Unmarshal([]byte(redisData), &locationResponse)
		if redisError == nil {
			logrus.Infof("Successfully fetched the data for location from the Redis")
			c.JSON(http.StatusOK, locationResponse)
			return
		}
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	location := make(map[string]int)
	data := scyllaClient.Query("SELECT office_location FROM employee_info").Iter()
	for data.Scan(&employee.OfficeLocation) {
		_, exist := location[employee.OfficeLocation]
		if exist {
			location[employee.OfficeLocation] += 1
		} else {
			location[employee.OfficeLocation] = 1
		}
	}
	jsonData, err := json.Marshal(location)
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	if redisEnabled && redisError == redis.Nil {
		writeinRedis("location", string(jsonData))
	}
	logrus.Infof("Successfully fetched the data for all location from the ScyllaDB")
	json.Unmarshal(jsonData, &locationResponse)
	c.JSON(http.StatusOK, locationResponse)
	return
}

// @Summary ReadCompleteEmployeesData is a method to read all employee's information
// @Schemes http
// @Description Read all employee data from database
// @Tags employee
// @Accept json
// @Produce json
// @Success 200 {array} model.Employee
// @Router /search/all [get]
// ReadCompleteEmployeesData is a method to read all employee information
func ReadCompleteEmployeesData(c *gin.Context) {
	var employee model.Employee
	var response []model.Employee
	var redisError error
	var redisData string
	if redisEnabled {
		redisData, redisError = client.CreateRedisClient().HGet(ctx, "employee", "all_data").Result()
		if redisError != nil {
			logrus.Warnf("Unable to read data from Redis %v", redisError)
		}
		json.Unmarshal([]byte(redisData), &response)
		if redisError == nil {
			logrus.Infof("Successfully fetched the data for all employee from the Redis")
			c.JSON(http.StatusOK, response)
			return
		}
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	data := scyllaClient.Query("SELECT id, name, designation, department, joining_date, address, office_location, status, email, phone_number FROM employee_info").Iter()
	for data.Scan(&employee.ID, &employee.Name, &employee.Designation, &employee.Department, &employee.JoiningDate, &employee.Address, &employee.OfficeLocation, &employee.Status, &employee.EmailID, &employee.PhoneNumber) {
		response = append(response, employee)
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	if redisEnabled && redisError == redis.Nil {
		writeinRedis("all_data", string(jsonData))
	}
	logrus.Infof("Successfully fetched the data for all employee from the ScyllaDB")
	c.JSON(http.StatusOK, response)
}

// @Summary ReadEmployeeData is a method to read employee information
// @Schemes http
// @Description Read data from database
// @Tags employee
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} model.Employee
// @Router /search [get]
// ReadEmployeeData is a method to read employee information
func ReadEmployeeData(c *gin.Context) {
	var response model.Employee
	mapData := map[string]interface{}{}
	var redisError error
	var redisData string
	id, keyExists := c.GetQuery("id")
	if !keyExists {
		logrus.Errorf("Query request of data without params")
		errorResponse(c, "Unable to perform search operation, query params not defined")
		return
	}
	if redisEnabled {
		redisData, redisError = client.CreateRedisClient().HGet(ctx, "employee", id).Result()
		if redisError != nil {
			logrus.Warnf("Unable to read data from Redis %v", redisError)
		}
		json.Unmarshal([]byte(redisData), &response)
		if redisError == nil {
			logrus.Infof("Successfully fetched the data for %s from the Redis", id)
			c.JSON(http.StatusOK, response)
			return
		}
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	data := scyllaClient.Query("SELECT * FROM employee_info where id = ?", id).Iter()
	for data.MapScan(mapData) {
		jsonData, err := json.Marshal(mapData)
		if err != nil {
			logrus.Errorf("Error in reading data from scylladb: %v", err)
			errorResponse(c, "Cannot read data from the system, request failure")
			return
		}
		if redisEnabled && redisError == redis.Nil {
			writeinRedis(id, string(jsonData))
		}
		json.Unmarshal(jsonData, &response)
		logrus.Infof("Successfully fetched the data for %s from the ScyllaDB", id)
		c.JSON(http.StatusOK, response)
		return
	}
}

// @Summary CreateEmployeeData is a method to write employee information in database
// @Schemes http
// @Description Write data in database
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Employee Data"
// @Success 200 {object} model.Employee
// @Router /create [post]
// CreateEmployeeData is a method to write employee information in database
func CreateEmployeeData(c *gin.Context) {
	var request model.Employee
	if err := c.BindJSON(&request); err != nil {
		logrus.Errorf("Error parsing the request body in JSON: %v", err)
		errorResponse(c, "Unable to Bind JSON in defined format, seems malformed")
		return
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in writing data to scylladb: %v", err)
		errorResponse(c, "Cannot write data to the system, request failure")
		return
	}
	defer scyllaClient.Close()
	date, err := time.Parse("2006-01-02", request.JoiningDate)
	if err != nil {
		logrus.Errorf("Error in writing data to scylladb: %v", err)
		errorResponse(c, "Cannot write data to the system, request failure")
		return
	}
	queryString := "INSERT INTO employee_info(id, name, designation, department, joining_date, address, office_location, status, email, phone_number) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	if err := scyllaClient.Query(queryString,
		request.ID, request.Name, request.Designation, request.Department, date, request.Address, request.OfficeLocation, request.Status, request.EmailID, request.PhoneNumber).Exec(); err != nil {
		logrus.Errorf("Error in writing data to scylladb: %v", err)
		errorResponse(c, "Cannot write data to the system, request failure")
		return
	}
	data := model.CustomMessage{
		Message: "Successfully created the data for the user",
	}
	logrus.Infof("Successfully created the employee record")
	c.JSON(http.StatusOK, data)
}

// writeinRedis is a method to write the data in Redis cache
func writeinRedis(cacheKey, cacheValue string) {
	_, err := client.CreateRedisClient().HSet(ctx, "employee", cacheKey, cacheValue).Result()
	if err != nil {
		logrus.Errorf("Error in reading writing data to Redis: %v", err)
	}
}

// errorResponse is a method to return bad http code in Gin
func errorResponse(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, model.CustomMessage{
		Message: err,
	})
}
