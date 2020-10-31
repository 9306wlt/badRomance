package arrays

import (
	"fmt"
	"testing"
)

/**

给你一个整数数组 nums 。

如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。

返回好数对的数目。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-good-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func numIdenticalPairs(nums []int) int {
	var count int
	for i:=0;i<len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] == nums[j]{
				count++
			}
		}
	}
	return count
}

func Test_numIdenticalPairs(t *testing.T){
	a := make([]int,0)
	a = append(a,1,1,1,1)
	fmt.Println(numIdenticalPairs(a))
}
