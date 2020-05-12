package dsstring

import(
	"fmt"
)


func FindNonRepeatingCha(s string)string {

	d := make(map[string]int)
	for _,eachChar := range s {
		c := fmt.Sprintf("%c",eachChar)
		d[c] = d[c]+1
	}
	for _,eachChar := range s {
		c := fmt.Sprintf("%c",eachChar)
		if val,ok := d[c];ok {
			if val == 1 {
				return c
			}
		}
	}
	return ""
}

func PrintAnagraTogether(arr []string)[][]string {

	output := make([][]string,0)
	hashTable := make(map[string][]string)
	for _,eachString := range arr {
		hashKey := sortString(eachString)
		hashTable[hashKey] = append(hashTable[hashKey],eachString)
	}
	for _,eachData := range hashTable {
		output = append(output,eachData)
	}
	return output
}

func LongestCommonPrefix(s []string)string {

	output := s[0]
	var noPrefix bool
	for _,eachString := range s {
		if noPrefix {
			break
		} 
		for {
			if len(eachString) < len(output) {
				output = output[:len(eachString)]
			} else {
				eachString = eachString[:len(output)]
			}
			if output == "" {
				noPrefix = true
			}
			if eachString == output {
				//best case for now we have prefix
				break
			} else {
				eachString = eachString[:len(eachString)-1]
				output = output[:len(output)-1]
			}
		}
	}
	return output
}

