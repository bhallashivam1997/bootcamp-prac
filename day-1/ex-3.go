package main

import "fmt"

type employeeSalary interface {
	salaryFullTime() float64
	salaryContractor() float64
	salaryFreelancer() float64
}

func salaryFullTime(){
	fmt.Println("How many day has the employee worked ?")
	var totalDays int
	fmt.Scanln(&totalDays)
	ans := (totalDays*500.0) + 1000.0
	fmt.Printf("Total Salary is: %d\n", ans)
	return
}

func salaryContractor(){
	fmt.Println("How many day has the employee worked ?")
	var totalDays int
	fmt.Scanln(&totalDays)
	ans := (totalDays*100.0) + 200.0
	fmt.Printf("Total Salary is: %d\n", ans)
	return
}

func salaryFreelancer(){
	fmt.Println("How many hours have the employee worked ?")
	var totalDays int
	fmt.Scanln(&totalDays)
	ans := totalDays*10.0
	fmt.Printf("Total Salary is: %d\n", ans)
	return
}

func main(){
	fmt.Println("Enter Your Employee type: Fulltime or Contractor or Freelancer")
	var employeeType string
	fmt.Scanln(&employeeType)
	if employeeType == "Fulltime"{
		salaryFullTime()
	}else if employeeType == "Contractor"{
		salaryContractor()
	}else if employeeType == "Freelancer"{
		salaryFreelancer()
	}else{
		fmt.Println("Wrong Employee Type")
	}

}
