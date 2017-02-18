/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:56:00-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-18T13:00:46-05:00
 */

package sort

// Sort slice
type Sort struct {
	slice *[]int
}

// New create new Sort object with slice
func New(slice *[]int) (*Sort, error) {
	var s = Sort{slice: slice}
	return &s, nil
}
