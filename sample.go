package stats

import (
	"math/rand"
	"errors"
	"math"
)

// Sample returns sample from input with replacement or without
func Sample(input Float64Data, takenum int, replacement bool) ([]float64, error) {

	if input.Len() == 0 {
		return nil, EmptyInput
	}

	length := input.Len()
	if replacement {

		result := Float64Data{}
		rand.Seed(unixnano())

		// In every step, randomly take the num for
		for i := 0; i < takenum; i++ {
			idx := rand.Intn(length)
			result = append(result, input[idx])
		}

		return result, nil

	} else if !replacement && takenum <= length {

		rand.Seed(unixnano())

		// Get permutation of number of indexies
		perm := rand.Perm(length)
		result := Float64Data{}

		// Get element of input by permutated index
		for _, idx := range perm[0:takenum] {
			result = append(result, input[idx])
		}

		return result, nil

	}

	return nil, BoundsErr
}

// 扩充数据或者缩减数据
func Sample2(input Float64Data, takenum,digits int) ([]float64, error){
	var inputLen = input.Len()
	if inputLen == 0 {
		return nil, EmptyInput
	}

	if takenum < inputLen{
		return nil,errors.New("暂不考虑缩减数据")
	}

	values := make([]float64,0)
	result := make([]float64,0)
	round := takenum/inputLen

	for i:=0;i< inputLen;i++{
		for j:=0;j<round;j++ {
			if i==0{
				value := roundFloat(RandFloat64(input[i]/2,input[i]),digits)

				values = append(values,value)
			}else{
				value := roundFloat(RandFloat64(input[i-1],input[i]),digits)
				values = append(values,value)
			}
		}
		values = append(values,input[i])
	}

	resultLen := len(values)
	if removeLen := resultLen - takenum;removeLen > 0 {
 		period := resultLen/removeLen
		startIdx := period/2
		idxMap := make(map[int]interface{})

		for removeLen > 0{
			idxMap[startIdx] = true
			startIdx += period
			removeLen--
		}

		for i,value := range values{
			if _,exists := idxMap[i];exists{
				continue
			}
			result = append(result,value)
		}
	}

	return result,nil
}

// rand a float beetween min and max
func RandFloat64(min,max float64) float64{
	rand.Seed(unixnano())
	result := (rand.Float64()*(max-min))+ min

	return result
}

// fmt %.2f
func roundFloat(f float64, n int) float64 {
	if n <= 0 {
		return  f
	}
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}