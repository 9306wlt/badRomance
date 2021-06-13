package method

import (
	"fmt"
	"testing"
)

type Employee struct {
	salary float64
}

func(e *Employee)giveRaise(p float64)float64{
	return e.salary*(1+p)
}

func Test_employee_salary(t *testing.T)  {
	empl := new(Employee)
	empl.salary = 12000.00
	fmt.Printf("原来薪水 %.2f\n",empl.salary)
	fmt.Printf("增加的薪水后 %.2f\n",empl.giveRaise(0.13))
	fmt.Println(empl.salary)
}
