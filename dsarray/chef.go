package dsarray
import (
	"strings"
	"math"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func CookDish(numberOfDays int,numberOfIngredients int,ingredients []string) string {
	
    counterMap := make(map[string]int)
    var result string
	dayWiseIntg := make([]string,0)
	//break point will found out the quantity of dominating type weight
	breakingPoint := int(math.Ceil(float64(numberOfIngredients)*float64(0.6)))
	var ingredientCont int
    for i := 0; i < len(ingredients); i++ {
		ingredientCont = ingredientCont+1
        ingredientId := ingredients[i]
        dayWiseIntg = append(dayWiseIntg,ingredientId)
        var intg string
		intg =  getIngType(ingredientId)
		counterMap[intg] =  counterMap[intg] +1
		//can we make dish check
		majorIngrediat := canWeMakeDish(counterMap,breakingPoint)

		// if we have enought ingredient and also chef also have dominating type 
		//then we can think of cook
		if majorIngrediat != "" && ingredientCont >= numberOfIngredients {
			extraIntg := numberOfIngredients-breakingPoint
			majorIntgCount := breakingPoint
			for i,eachIng := range dayWiseIntg {
				if eachIng != "" {
					intg :=  getIngType(eachIng)	
					if strings.Contains(eachIng,majorIngrediat) && majorIntgCount >0 {
						//it means it's major ing 
						majorIntgCount = majorIntgCount-1	
						//updating ingt count
						counterMap[intg]  = counterMap[intg] -1
						result += eachIng+":"
						dayWiseIntg[i] = ""	
					} else if extraIntg >0  {
						//extra ing
						extraIntg = extraIntg-1
						//updating ingt count
						counterMap[intg]  = counterMap[intg] -1
						result += eachIng+":"
						dayWiseIntg[i] = ""	
					} 
				}
			}

			result = strings.TrimRight(result, ":")
			ingredientCont = ingredientCont-numberOfIngredients
		} else {
			// today is not the day for cook try next day
			result += "-" 
		}
	}
	return result
}

//
// getIngType will return the type of ingrediant form ingID
//
func getIngType(id string)string {
	var intg string
	if strings.Contains(id, "FAT") {
		intg = "FAT"
	}else if strings.Contains(id, "FIBER") {
		intg = "FIBER"
	} else if strings.Contains(id, "CARB"){
		intg = "CARB"
	}
	return intg
}

//
// canWeMakeDish will check are we have proper ingrediant to make dish
// with dominating ingrediant
//
func canWeMakeDish(counterMap map[string]int,breakingPoint int)string {
	for key,eachVal := range  counterMap {
		if eachVal >= breakingPoint {
			return key
		}
	}
	return ""
}