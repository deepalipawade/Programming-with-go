// You can edit this code!
// Click here and start typing.
// What is a pasacal's triangle ?
// In mathematics, Pascal's triangle is a triangular array of the binomial coefficients that arises in probability theory, combinatorics, and algebra.
package main

import "fmt"

// pascalTri ... is map of row count to actual row containing coefficients
var pascalTri map[int32][]int32

func main() {
	fmt.Println("Hello, 世界", pascalTri)
	pascalTri = map[int32][]int32{}
	fmt.Println("Output : ", buildPascalTri(3))
}

func buildPascalTri(n int32) map[int32][]int32 {
	if n == int32(1) {          // base confition  .... this will initiate acutal buildling of pascals triangle
		pascalTri[n] = append(pascalTri[n], 1)
		return pascalTri
	}
  pascalTri = buildPascalTri(n - 1) // recursive call to n - 1  (previou row as we need to reach a base condition in order to start builing pascal triangle)
	for i := int32(0); i < n; i++ {
		if i == 0 || i == n-1 { // condition of corner blocks or end elements of map which will always be 1 
			pascalTri[n] = append(pascalTri[n], 1)
			continue
		}
		prev_row := pascalTri[n-1]
		pascalTri[n] = append(pascalTri[n], prev_row[i-1]+prev_row[i])   
	}
	return pascalTri
}
