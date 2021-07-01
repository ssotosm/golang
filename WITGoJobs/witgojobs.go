package main

import "fmt"

func main() {

	employerType := CreateEmployerType(1, 1, "name", "statement", "logo", "url", "description", "addressline1", "addressline2", "city", "state", "postalCode", "country")
	fmt.Println("Create employerType = ", employerType)

	employerType.UpdateEmployerType(1, 1, "nameUpdate", "statement", "logo", "url", "description", "addressline1", "addressline2", "city", "state", "postalCode", "country")
	fmt.Println("Update employerType = ", employerType)

	existEmployerType := employerType.DeleteEmployerType(1)
	fmt.Println("Delete employerType = ", employerType)
	fmt.Println("Delete employerType successfull = ", *existEmployerType)

	employerJob := CreateEmployerTypeJob(1, 1, "title", "description", "jobType", 1)
	fmt.Println("Create employerJob = ", employerJob)

	employerJob.UpdateEmployerTypeJob(1, 1, "titleUpdate", "description", "jobType", 1)
	fmt.Println("Update employerJob = ", employerJob)

}

/* EMPLOYER */

type EmployerType struct {
	Id           int
	Status       int
	Name         string
	Statement    string
	Logo         string
	Url          string
	Description  string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PostalCode   string
	Country      string
}

func CreateEmployerType(id int, status int, name string, statement string, logo string, url string, description string, addressLine1 string, addressLine2 string, city string, state string, postalCode string, country string) *EmployerType {

	pEmp := new(EmployerType)
	pEmp.Id = id
	pEmp.Status = status
	pEmp.Name = name
	pEmp.Statement = statement
	pEmp.Logo = logo
	pEmp.Url = url
	pEmp.Description = description
	pEmp.AddressLine1 = addressLine1
	pEmp.AddressLine2 = addressLine2
	pEmp.City = city
	pEmp.State = state
	pEmp.PostalCode = postalCode
	pEmp.Country = country

	return pEmp
}

func (pEmp *EmployerType) UpdateEmployerType(id int, status int, name string, statement string, logo string, url string, description string, addressLine1 string, addressLine2 string, city string, state string, postalCode string, country string) *EmployerType {

	pEmp.Id = id
	pEmp.Status = status
	pEmp.Name = name
	pEmp.Statement = statement
	pEmp.Logo = logo
	pEmp.Url = url
	pEmp.Description = description
	pEmp.AddressLine1 = addressLine1
	pEmp.AddressLine2 = addressLine2
	pEmp.City = city
	pEmp.State = state
	pEmp.PostalCode = postalCode
	pEmp.Country = country

	return pEmp

}

func (pEmp *EmployerType) DeleteEmployerType(id int) *bool {

	exist := false
	if pEmp.Id == id {
		pEmp = nil
		exist = true
		return &exist
	}
	return &exist
}

/* EMPLOYER JOB */

type EmployerTypeJob struct {
	Id             int
	EmployerTypeId int
	Title          string
	Description    string
	JobType        string
	IsFilled       int
}

func CreateEmployerTypeJob(id int, employerId int, title string, description string, jobType string, isFilled int) *EmployerTypeJob {

	pEmpJob := new(EmployerTypeJob)
	pEmpJob.Id = id
	pEmpJob.EmployerTypeId = employerId
	pEmpJob.Title = title
	pEmpJob.Description = description
	pEmpJob.JobType = jobType
	pEmpJob.IsFilled = isFilled

	return pEmpJob

}

func (pEmpJob *EmployerTypeJob) UpdateEmployerTypeJob(id int, employerId int, title string, description string, jobType string, isFilled int) *EmployerTypeJob {

	pEmpJob.Id = id
	pEmpJob.EmployerTypeId = employerId
	pEmpJob.Title = title
	pEmpJob.Description = description
	pEmpJob.JobType = jobType
	pEmpJob.IsFilled = isFilled

	return pEmpJob
}
