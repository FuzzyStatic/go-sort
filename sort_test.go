/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:08:00-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-19T17:19:51-05:00
 */

package sort

import (
	"reflect"
	"testing"
)

var sorted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}  // expected sort value
var sorted2 = []int{1, 2, 4, 4, 5, 6, 7, 7, 8} // expected sort value

func TestBubbleSort(t *testing.T) {
	slice1 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test bubble sort of %v", *slice1)
	s1, _ := New(slice1)
	s1.BubbleSort()

	if !reflect.DeepEqual(*slice1, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice1)
	}

	slice2 := &[]int{6, 5, 4, 4, 8, 7, 7, 2, 1}
	t.Logf("Test bubble sort of %v", *slice2)
	s2, _ := New(slice2)
	s2.BubbleSort()

	if !reflect.DeepEqual(*slice2, sorted2) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted2, *slice2)
	}

	slice3 := &[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	t.Logf("Test bubble sort of %v", *slice3)
	s3, _ := New(slice3)
	s3.BubbleSort()

	if !reflect.DeepEqual(*slice3, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice3)
	}
}

func TestCountingSort(t *testing.T) {
	slice1 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test count sort of %v", *slice1)
	s1, _ := New(slice1)
	s1.CountingSort()

	if !reflect.DeepEqual(*slice1, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice1)
	}

	slice2 := &[]int{6, 5, 4, 4, 8, 7, 7, 2, 1}
	t.Logf("Test count sort of %v", *slice2)
	s2, _ := New(slice2)
	s2.CountingSort()

	if !reflect.DeepEqual(*slice2, sorted2) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted2, *slice2)
	}

	slice3 := &[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	t.Logf("Test count sort of %v", *slice3)
	s3, _ := New(slice3)
	s3.CountingSort()

	if !reflect.DeepEqual(*slice3, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice3)
	}
}

func TestInsertionSort(t *testing.T) {
	slice1 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test insertion sort of %v", *slice1)
	s1, _ := New(slice1)
	s1.InsertionSort()

	if !reflect.DeepEqual(*slice1, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice1)
	}

	slice2 := &[]int{6, 5, 4, 4, 8, 7, 7, 2, 1}
	t.Logf("Test insertion sort of %v", *slice2)
	s2, _ := New(slice2)
	s2.InsertionSort()

	if !reflect.DeepEqual(*slice2, sorted2) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted2, *slice2)
	}

	slice3 := &[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	t.Logf("Test insertion sort of %v", *slice3)
	s3, _ := New(slice3)
	s3.InsertionSort()

	if !reflect.DeepEqual(*slice3, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice3)
	}
}

func TestMergeSort(t *testing.T) {
	slice1 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test merge sort of %v", *slice1)
	s1, _ := New(slice1)
	s1.MergeSort(0, len(*slice1)-1) // Start at 0 and go to the last index value

	if !reflect.DeepEqual(*slice1, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice1)
	}

	slice2 := &[]int{6, 5, 4, 4, 8, 7, 7, 2, 1}
	t.Logf("Test merge sort of %v", *slice2)
	s2, _ := New(slice2)
	s2.MergeSort(0, len(*slice2)-1) // Start at 0 and go to the last index value

	if !reflect.DeepEqual(*slice2, sorted2) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted2, *slice2)
	}

	slice3 := &[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	t.Logf("Test merge sort of %v", *slice3)
	s3, _ := New(slice3)
	s3.MergeSort(0, len(*slice3)-1) // Start at 0 and go to the last index value

	if !reflect.DeepEqual(*slice3, sorted) {
		t.Errorf("Expected slice of %v, but it was %v instead.", sorted, *slice3)
	}
}
