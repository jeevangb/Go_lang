package shape

import (
	"GoProgrammes/Desktop/go_assignments/Day2_Day3/assignment_8/question_d/model"
	"fmt"
)

func AreaOfCircle(r model.Radius) {

	res := 3.14 * r.Rad * r.Rad
	fmt.Println(res)

}
