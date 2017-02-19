/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-19T14:16:29-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-19T14:27:15-05:00
 */

package sort

// BubbleSort is based on the idea of repeatedly comparing pairs of adjacent elements
// and then swapping their positions if they exist in the wrong order.
func (s *Sort) BubbleSort() {
	var temp int

	for k := 0; k < len(*s.slice)-1; k++ {
		// (len(*s.slice)-k-1) is for ignoring comparisons of elements which have already been compared in earlier iterations
		for i := 0; i < len(*s.slice)-k-1; i++ {
			if (*s.slice)[i] > (*s.slice)[i+1] {
				// Here swapping of positions is being done.
				temp = (*s.slice)[i]
				(*s.slice)[i] = (*s.slice)[i+1]
				(*s.slice)[i+1] = temp
			}
		}
	}
}
