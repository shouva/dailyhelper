package dailyhelper

import "strings"

// SnakeCaseToCamelCase :
func SnakeCaseToCamelCase(inputUnderScoreStr string, firstUpperCase bool) (camelCase string) {
	//snake_case to camelCase

	isToUpper := false

	for k, v := range inputUnderScoreStr {
		if k == 0 {
			if firstUpperCase {
				camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
			} else {
				camelCase = strings.ToLower(string(inputUnderScoreStr[0]))
			}
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return

}

// BeautySingularity : converting NameId or Nameid to nameID etc
func BeautySingularity(id string) string {
	id = strings.ReplaceAll(id, "Id", "ID")
	id = strings.ReplaceAll(id, "Api", "API")
	id = strings.ReplaceAll(id, "Ip", "IP")
	id = strings.ReplaceAll(id, "Uuid", "UUID")
	return id
}
