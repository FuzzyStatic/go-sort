/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-20T20:11:10-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-20T22:28:55-05:00
 */

package sort

import "math/rand"

func (s *Sort) partition(start, end int) int {
	i := start + 1
	piv := (*s.slice)[start] // Make the first element as pivot element.

	for j := start + 1; j <= end; j++ {
		/* Rearrange the array by putting elements which are less than pivot on one side
		   and which are greater that on other. */
		if (*s.slice)[j] < piv {
			(*s.slice)[i], (*s.slice)[j] = (*s.slice)[j], (*s.slice)[i]
			i++
		}
	}
	// Put the pivot element in its proper place.
	(*s.slice)[start], (*s.slice)[i-1] = (*s.slice)[i-1], (*s.slice)[start]
	return i - 1 // Return the position of the pivot
}

func (s *Sort) randomPartition(start, end int) int {
	// Chooses position of pivot randomly by using rand.Intn() function .
	random := start + rand.Intn(end-start+1)
	(*s.slice)[random], (*s.slice)[start] = (*s.slice)[start], (*s.slice)[random]
	return s.partition(start, end) // Call the above partition function
}

// QuickSort is based on the divide-and-conquer approach based on the idea of choosing
// one element as a pivot element and partitioning the array around it such that:
// Left side of pivot contains all the elements that are less than the pivot element
// Right side contains all elements greater than the pivot
func (s *Sort) QuickSort(start, end int) {
	if start < end {
		// Stores the position of pivot element
		pivPos := s.randomPartition(start, end)
		s.QuickSort(start, pivPos-1) //sorts the left side of pivot.
		s.QuickSort(pivPos+1, end)   //sorts the right side of pivot.
	}
}
