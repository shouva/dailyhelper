package dailyhelper

import "encoding/json"

// Struct2Map :
func Struct2Map(in interface{}) (out map[string]interface{}, err error) {
	inrec, err := json.Marshal(in)
	if err != nil {
		return
	}
	json.Unmarshal(inrec, &out)
	return
}
