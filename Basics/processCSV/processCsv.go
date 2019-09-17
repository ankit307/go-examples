package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	eid         string
	name        string
	email       string
	phoneNo     int64
	designation string
	salary      string
}

func main() {
	path, _ := os.Getwd()
	filePath := path + "/Basics/processCSV/employee.csv"
	emp := scanFileToEmployeeArray(filePath)
	grouped := groupByDesignation(emp)
	for i, v := range grouped {
		fmt.Println("Employeees in Department :" + i)
		for _, k := range v {
			fmt.Printf("\t %v \n", k)
		}
	}
}

func scanFileToEmployeeArray(filePath string) []Employee {
	file, err := os.Open(filePath)
	employees := []Employee{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i == 0 {
			i++
			scanner.Text()
			continue
		}
		s := strings.Split(scanner.Text(), ",")
		phoneNo, _ := strconv.ParseInt(s[2], 10, 64)
		e := Employee{eid: s[0], name: s[1], email: s[2], phoneNo: phoneNo, designation: s[4], salary: s[5]}
		employees = append(employees, e)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return employees
}

func groupByDesignation(emp []Employee) map[string][]Employee {
	group := make(map[string][]Employee)
	for _, v := range emp {
		group[v.designation] = append(group[v.designation], v)
	}
	return group
}
