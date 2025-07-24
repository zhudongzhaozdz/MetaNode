package main

import (
	"fmt"
	"slices"
	"strings"
)

// 只出现一次的数字
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for index, value := range nums {
		fmt.Println("index = %d,value =%d", index, value)
		num := value
		mv, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			m[num] = mv + 1
		}
	}
	var ret int
	for k, v := range m {
		fmt.Println("k = %d,v =%d", k, v)
		if v == 1 {
			ret = k
		}
	}
	return ret
}

// 回文数
func isPalindrome(x int) bool {
	var y int = 0
	// 输入数字的位数
	var len int = 0
	var temp int = x
	// 负数一定不是回文
	if x < 0 {
		return false
	}
	// 计算输入数字的位数
	for temp > 0 {
		temp /= 10
		len++
	}
	// 0到9只有一位数的一定是回文
	if len == 1 {
		return true
	}
	if len%2 == 0 {
		// 偶数刚好对称
		for i := len; i > len/2; i-- {
			y = y*10 + x%10
			x /= 10
		}
	} else {
		// 奇数截断取整
		for i := len; i > (len/2)+1; i-- {
			y = y*10 + x%10
			x /= 10
		}
		x /= 10
	}
	fmt.Println(x, y)
	if x == y {
		return true
	}
	return false
}

// 有效的括号
func isValid(s string) bool {
	m := make(map[uint8]uint8)
	m[']'] = '['
	m['}'] = '{'
	m[')'] = '('
	stackLen := 0
	var stack [10010]uint8
	if len(s)%2 == 1 {
		return false
	}
	for index, value := range s {
		fmt.Println("index = %d,value =%d", index, value)
		if stackLen > 0 && stack[stackLen] == m[uint8(value)] {
			stackLen--
		} else {
			stackLen++
			stack[stackLen] = uint8(value)
		}
	}
	fmt.Println("m value =%d ", m)
	// 入栈数，如果相等，返回true,否则，返回false
	if stackLen == 0 {
		return true
	}
	return false
}

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = longestCommonPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func longestCommonPrefix(str0 string, str string) string {
	prefix := strings.Builder{}
	for i := 0; i < len(str0) && i < len(str); i++ {
		if str0[i] != str[i] {
			break
		}
		prefix.WriteByte(str0[i])
	}
	return prefix.String()
}

// 最长公共前缀
func LongestCommonPrefix1(strs []string) string {
	prefix0 := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix := strings.Builder{}
		for j := 0; j < len(prefix0) && j < len(strs[i]); j++ {
			if prefix0[j] != strs[i][j] {
				break
			}
			prefix.WriteByte(prefix0[j])
		}
		prefix0 = prefix.String()
		if len(prefix0) == 0 {
			break
		}
	}
	return prefix0
}

// 数组最后一位加一
func plusOne(digits []int) []int {
	over := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + over
		digits[i] = sum % 10
		over = sum / 10
	}
	fmt.Println("digits = %d", digits)
	if over > 0 {
		ret := make([]int, 1)
		ret[0] = over
		return append(ret, digits...)
	}
	return digits
}

// 两数之和
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, x := range nums {
		if p, ok := hashMap[target-x]; ok {
			return []int{p, i}
		}
		hashMap[x] = i
	}
	return nil
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	k := 1
	//保存不重复的新数据
	hashMap := map[int]int{}
	//第一项不会是重复项
	hashMap[k-1] = nums[0]
	for i := 1; i < len(nums); i++ {
		// nums[i] 不是重复项
		if nums[i] != nums[i-1] {
			// 保留 nums[i]
			nums[k] = nums[i]
			hashMap[k] = nums[k]
			k++
		}
	}
	fmt.Println("nnums = %d", hashMap)
	return k
}

// 合并区间
func merge(intervals [][]int) (ans [][]int) {
	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	fmt.Println("intervals = %d", intervals)
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else {
			ans = append(ans, p) // 新的合并区间
		}
	}
	return
}
func main() {
	// nums := []int{4, 1, 2, 1, 2}
	// var ret int = singleNumber(nums)
	// fmt.Println("ret = ", ret)

	// x := 10011
	// var ret bool = isPalindrome(x)
	// fmt.Println("ret = ", ret)

	// var s string = "{(}){{}}[]"
	// var ret bool = isValid(s)
	// fmt.Println("ret = ", ret)

	// var strs []string = []string{"dog", "racecar", "car"}
	// // var strs []string = []string{"flower", "flow", "flight"}
	// var ret string = LongestCommonPrefix1(strs)
	// if len(ret) != 0 {
	// 	fmt.Println("ret = ", ret)
	// } else {
	// 	fmt.Println("输入不存在公共前缀。", ret)
	// }

	// var digits []int = []int{1, 2, 3, 4, 9}
	// var ret []int = plusOne(digits)
	// fmt.Println("ret = %d", ret)

	// var nums []int = []int{2, 5, 8, 90}
	// var target = 130
	// var ret []int = twoSum(nums, target)
	// fmt.Println("ret = %d", ret)

	// var nums []int = []int{1, 2, 2, 3, 3, 32}
	// var ret int = removeDuplicates(nums)
	// fmt.Println("ret = %d", ret)

	var intervals [][]int = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 21}, {12, 14}}
	var ans [][]int = merge(intervals)
	fmt.Println("ans = %d", ans)
}
