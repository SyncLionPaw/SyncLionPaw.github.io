package main


// partition func 
func dualPartiiton(nums []int, target int, l, r int) (int, int) {
	less, great := l-1, r+1
	for i := 0; i < great; i++ {
		if d := target - nums[i]; d == 0 {
			continue
		} else if d < 0 {
			great--
			nums[great], nums[i] = nums[i], nums[great]
			i--
		} else {
			less++
			nums[less], nums[i] = nums[i], nums[less]
		}
	}
	return less, great
}
