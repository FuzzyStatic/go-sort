/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:08:00-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-18T12:27:49-05:00
 */

package mergeSort

// MergeSort slice
type MergeSort struct {
	slice *[]int
}

// New create new MergSort object with slice
func New(s *[]int) (*MergeSort, error) {
	var m = MergeSort{slice: s}
	return &m, nil
}

// Merge merges two slices by comparing their elements
func (m *MergeSort) Merge(start, mid, end int) {
	p := start
	q := mid + 1

	sAux := make([]int, end-start+1)
	k := 0

	for i := start; i <= end; i++ {
		if p > mid { //checks if first part comes to an end or not
			sAux[k] = (*m.slice)[q]
			k++
			q++
		} else if q > end { //checks if second part comes to an end or not
			sAux[k] = (*m.slice)[p]
			k++
			p++
		} else if (*m.slice)[p] < (*m.slice)[q] { //checks which part has smaller element
			sAux[k] = (*m.slice)[p]
			k++
			p++
		} else {
			sAux[k] = (*m.slice)[q]
			k++
			q++
		}
	}

	for p := 0; p < k; p++ {
		/* Now the real array has elements in sorted manner including both parts*/
		(*m.slice)[start] = sAux[p]
		start++
	}
}

// MergeSort splits slice into two, sorts (or splits again) and merges the sorted slices
func (m *MergeSort) MergeSort(start, end int) {
	if start < end {
		mid := (start + end) / 2 // defines the current array in 2 parts
		m.MergeSort(start, mid)  // sort the 1st part of array
		m.MergeSort(mid+1, end)  // sort the 2nd part of array
		m.Merge(start, mid, end)
	}
}
