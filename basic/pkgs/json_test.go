package pkgs

import (
	"encoding/json"
	"fmt"
	"goSecond/basic/model"
	"testing"
)

func TestJson2Map(t *testing.T) {

	str := `{"da":{"type":"hello","operate":"add","value":123},"abc":{"type":"world","operate":"minor","value":321,"show_value":323}}`

	bytes := []byte(str)

	_map := make(map[string]model.FieldItem, 0)

	err := json.Unmarshal(bytes, &_map)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(_map)
}
