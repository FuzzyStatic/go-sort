/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-19T14:42:30-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-19T17:25:04-05:00
 */

package sort

// InsertionSort is based on the idea that one element from the input elements is
// consumed in each iteration to find its correct position i.e, the position to which
// it belongs in a sorted array.
func (s *Sort) InsertionSort() {
	for i := 0; i < len(*s.slice); i++ {
		// Storing current element whose left side is checked for its correct position
		temp := (*s.slice)[i]
		j := i

		/* Check whether the adjacent element in left side is greater or less than the
		   current element. */
		for j > 0 && temp < (*s.slice)[j-1] {

			// Moving the left side element to one position forward.
			(*s.slice)[j] = (*s.slice)[j-1]
			j = j - 1
		}
		// Moving current element to its  correct position.
		(*s.slice)[j] = temp
	}
}
