package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name string
	BaseSalary, WorkingMonth int
}

type Senior struct {
	Name string
	BaseSalary, WorkingMonth int
	PerformanceRate float64

}

type Manager struct {
	Name string
	BaseSalary, WorkingMonth int
	PerformanceRate float64
	BonusManagerRate float64
}


//Input: Junior { Name: "Junior 1", BaseSalary: 100000, WorkingMonth: 12 }
//Output: 100000

//Input: Senior { Name: "Senior 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5 }
//Output: 250000

//Input: Manager { Name: "Manager 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5 }
//Output: 300000

	func (j *Junior) GetBonus() float64 {
		r := float64(j.WorkingMonth) / 12
		if r > 1 {
			r = 1
		}
		return float64(j.BaseSalary) * r
	}

	func (s *Senior) GetBonus() float64 {
		r := float64(s.WorkingMonth) / 12
		if r > 1 {
			r = 1
		}
		return 2*float64(s.BaseSalary)*r + s.PerformanceRate*float64(s.BaseSalary)
	}


	func (m *Manager) GetBonus() float64 {
		r := float64(m.WorkingMonth) / 12
		if r > 1 {
			r = 1
		}
		return 2*float64(m.BaseSalary)*r + m.PerformanceRate*float64(m.BaseSalary) + m.BonusManagerRate*float64(m.BaseSalary)
	}

func EmployeeBonus(employee Employee) float64 {

return employee.GetBonus()
	// return 0.0 // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	total := 0.0
	for _, employee := range employees {
		total += employee.GetBonus()
	}
	return total
	// return nil // TODO: replace this
}


func main(){
	adi := Junior{
		Name: "Adi",
		BaseSalary: 1000000,
		WorkingMonth: 12,
	}
	budi := Senior{
		Name: "Budi",
		BaseSalary: 2000000,
		WorkingMonth: 12,
		PerformanceRate: 1.2,
	}

	ciko := Manager{
		Name: "Ciko",
		BaseSalary: 3000000,
		WorkingMonth: 12,
		PerformanceRate: 1.5,
		BonusManagerRate: 0.1,
	}

	employees := []Employee{&adi, &budi, &ciko}
	fmt.Println(TotalEmployeeBonus(employees))
}