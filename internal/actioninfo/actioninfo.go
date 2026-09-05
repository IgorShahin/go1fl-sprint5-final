package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, step := range dataset {
		err := dp.Parse(step)
		if err != nil {
			fmt.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
