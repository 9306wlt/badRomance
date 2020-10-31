package arrays

import (
	"testing"
)

/**

给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var max = candies[0]
	for i:=1;i<len(candies);i++{
		if candies[i] > max{
			max = candies[i]
		}
	}
	out := make([]bool,0)
	for i:=0;i<len(candies);i++{
		if candies[i] + extraCandies >= max{
			out = append(out,true)
		}else{
			out = append(out,false)
		}
	}
	return out
}

func Test_kidsWithCandies(t *testing.T){
	candies := []int{4,2,1,1,2}
	kidsWithCandies(candies,1)
}
