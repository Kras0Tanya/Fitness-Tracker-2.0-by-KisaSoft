package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("Parsing error:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Action info generation error:", err)
			continue
		}

		fmt.Println(info)
	}
}
