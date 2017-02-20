/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-19T18:51:10-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-19T19:05:33-05:00
 */

package sort

// SelectionSort is based on the idea of finding the minimum or maximum element in
// an unsorted array and then putting it in its correct position in a sorted array.
func (s *Sort) SelectionSort() {
	// Temporary variable to store the position of minimum element
	var minimum int

	// Reduces the effective size of the array by one in each iteration.
	for i := 0; i < len(*s.slice)-1; i++ {
		// Assuming the first element to be the minimum of the unsorted array
		minimum = i

		// Gives the effective size of the unsorted array
		for j := i + 1; j < len(*s.slice); j++ {
			if (*s.slice)[j] < (*s.slice)[minimum] { // Finds the minimum element
				minimum = j
			}
		}

		// Putting minimum element on its proper position.
		(*s.slice)[minimum], (*s.slice)[i] = (*s.slice)[i], (*s.slice)[minimum]
	}
}
