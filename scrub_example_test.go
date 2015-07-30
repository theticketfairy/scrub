package scrub

import "fmt"

type Employee struct {
	Name   string
	Title  string
	Salary float64
	Owns   float64
}

func (e *Employee) Form() Form {
	name := NewStringField("name", e.Name)
	name.Required()

	title := NewStringField("title", e.Title)
	title.Required()

	salary := NewFloat64Field("salary", e.Salary)
	salary.Min(38000)

	owns := NewFloat64Field("owns", e.Owns)
	owns.Between(0.25, 0.75)

	return Form{name, title, salary, owns}
}

type Employees []*Employee

type Company struct {
	Name        string
	Established int64
	President   *Employee
	Employees   Employees
}

func (c *Company) Form() Form {
	name := NewStringField("name", c.Name)
	name.Required()

	est := NewInt64Field("established", c.Established)
	est.Between(1900, 2015)

	president := NewNestedField("president", c.President)

	employees := NewNestedListField("employees", func() []Validated {
		cast := make([]Validated, len(c.Employees))
		for i := range c.Employees {
			cast[i] = c.Employees[i]
		}
		return cast
	})
	return Form{name, est, president, employees}
}

func Example() {
	c := &Company{
		"Pear",
		2016,
		&Employee{"Jeve Stobs", "CEO", 428532.35, 0.2},
		[]*Employee{
			{"Gill Bates", "COO", 20251.71, 0.5},
			{"Mon Elusk", "CTO", 299823.75, 0.6},
			{"", "Engineer", 100000, 0.75},
			{},
		},
	}

	e := Validate(c)

	fmt.Println(e.Describe())
	//Output:
	//* [max] established - The value of this field must be between 1900 and 2015
	//* [multi] president
	//*   [min] owns - The value of this field must be between 0.25 and 0.75
	//* [multi] employees
	//*   [multi] employees.0
	//*     [min] salary - The value of this field must be at least 38000
	//*   [multi] employees.2
	//*     [required] name - This field is required
	//*   [multi] employees.3
	//*     [required] name - This field is required
	//*     [required] title - This field is required
	//*     [min] salary - The value of this field must be at least 38000
	//*     [min] owns - The value of this field must be between 0.25 and 0.75
}
