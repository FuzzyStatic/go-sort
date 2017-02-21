/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:08:00-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-20T22:00:16-05:00
 */

package sort

// Merge two slices by comparing their elements
func (s *Sort) merge(start, mid, end int) {
	p := start
	q := mid + 1

	sAux := make([]int, end-start+1)
	k := 0

	for i := start; i <= end; i++ {
		if p > mid { // Checks if first part comes to an end or not
			sAux[k] = (*s.slice)[q]
			k++
			q++
		} else if q > end { // Checks if second part comes to an end or not
			sAux[k] = (*s.slice)[p]
			k++
			p++
		} else if (*s.slice)[p] < (*s.slice)[q] { // Checks which part has smaller element
			sAux[k] = (*s.slice)[p]
			k++
			p++
		} else {
			sAux[k] = (*s.slice)[q]
			k++
			q++
		}
	}

	for p := 0; p < k; p++ {
		/* Now the real array has elements in sorted manner including both parts */
		(*s.slice)[start] = sAux[p]
		start++
	}
}

// MergeSort is a divide-and-conquer algorithm based on the idea of breaking down
// a list into several sub-lists until each sublist consists of a single element and
// merging those sublists in a manner that results into a sorted list.
func (s *Sort) MergeSort(start, end int) {
	if start < end {
		mid := (start + end) / 2 // Defines the current array in 2 parts
		s.MergeSort(start, mid)  // Sort the 1st part of array
		s.MergeSort(mid+1, end)  // Sort the 2nd part of array
		s.merge(start, mid, end)
	}
}
