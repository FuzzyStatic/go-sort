/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:28:32-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-19T19:04:09-05:00
 */

package sort

// CountingSort the frequencies of distinct elements of the array to be sorted is
// counted and stored in an auxiliary array, by mapping its value as an index of the
// auxiliary array
func (s *Sort) CountingSort() {
	var k int

	// First, find the maximum value in s.slice
	for _, v := range *s.slice {
		if v > k {
			k = v
		}
	}

	sAux := make([]int, k+1) // Create sAux of size k + 1

	// Initialize the elements of sAux with 0
	for i := 0; i <= k; i++ {
		sAux[i] = 0
	}

	// Store the frequencies of each distinct element of s.slice,
	// by mapping its value as the index of sAux slice
	for _, v := range *s.slice {
		sAux[v] = sAux[v] + 1
	}

	sorted := make([]int, len(*s.slice))
	j := 0

	for i := 0; i <= k; i++ {
		// Aux stores which element occurs how many times,
		// Add i in sorted according to the number of times i occured in s.slice
		for tmp := sAux[i]; tmp > 0; tmp-- {
			sorted[j] = i
			j++
		}
	}

	*s.slice = sorted
}
