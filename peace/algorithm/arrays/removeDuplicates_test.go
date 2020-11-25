package arrays

import (
	"testing"
)

/**
26. 删除排序数组中的重复项

给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func removeDuplicates(nums []int) int {

	if len(nums) <= 1{
		return len(nums)
	}
	//快慢指针
	slow,fast := 0,0

	for fast < len(nums){
		//当快慢指针所指向元素不等，说明两个元素不相同，则先将慢指针++，然后将快指针指向的值赋给慢指针
		if nums[slow] != nums[fast]{
			//这里添加元素有点像栈的味道，当有元素要进栈，先腾出一个空间，然后再放进去。
			slow++
			nums[slow] = nums[fast]
		}
		//当快慢指针相等，则快指针继续前行。
		fast++
	}
	return slow+1

}

func Test_removeDuplicates(t *testing.T){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	removeDuplicates(nums)
}
