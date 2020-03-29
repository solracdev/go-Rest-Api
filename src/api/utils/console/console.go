package console

import (
	"encoding/json"
	"fmt"
)

// Pretty mesage
func Pretty(data interface{}) {

	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
