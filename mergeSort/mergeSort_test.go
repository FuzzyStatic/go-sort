/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:08:00-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-18T12:24:42-05:00
 */

package mergeSort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // expected sort value

	s1 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test merge sort of %v", *s1)
	m1, _ := New(s1)
	m1.MergeSort(0, len(*s1)-1) // Start at 0 and go to the last index value

	if !reflect.DeepEqual(*s1, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *s1)
	}

	s2 := &[]int{6, 5, 4, 9, 8, 7, 3, 2, 1}
	t.Logf("Test merge sort of %v", *s2)
	m2, _ := New(s2)
	m2.MergeSort(0, len(*s2)-1) // Start at 0 and go to the last index value

	if !reflect.DeepEqual(*s2, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *s2)
	}

	s3 := &[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	t.Logf("Test merge sort of %v", *s3)
	m3, _ := New(s3)
	m3.MergeSort(0, len(*s3)-1) // Start at 0 and go to the last index value

	if !reflect.DeepEqual(*s3, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *s3)
	}
}
