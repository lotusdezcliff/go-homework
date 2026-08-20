package homework01

import (
	//"fmt"
	"strings"
)
// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement
	// example: [1, 1, 2, 2, 3] -> [01, 01, 10, 10, 11]
	// A ^ A = 0 => XOR([]) = not repeat element
	result := 0
	for _, val := range nums {
	result = result ^ val
	}
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	// example: 12321, 12321 % 10 = 1, (12321 % 100) // 10  = 21
	// base false case: negative and 0 at the end
	if x < 0 || (x % 10 == 0 && x != 0) {
	return false
	}
	// get from the last digi and append to the flip_x until flip_x >= x
	flip_x := 0
	for flip_x < x {
	flip_x = (flip_x * 10) + (x % 10)
	x = x / 10
	}
	// even digit case: x == flip_x, odd digit case x == flip_x / 10
	return x == flip_x || x == flip_x / 10

	return false
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement
	// valid: "{[]}" ; invalid: "([)]"
	// if open bracket then copy to new string, if close bracket, find the last element of the newstring, if not matching bracket return false, if matched, remove the last element of the new string
	// if final newstring is empty return true else return false
	var new_s []rune
    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            new_s = append(new_s, char)
        } else if char == ')' || char == '}' || char == ']' {
            if len(new_s) == 0 {
                return false
            }

            // Take the last element
            last_index := len(new_s) - 1
            last_element := new_s[last_index]

            // pop the last element
            new_s = new_s[:last_index]

            // check the bracket match
            if (char == ')' && last_element != '(') ||
                (char == '}' && last_element != '{') ||
                (char == ']' && last_element != '[') {
                return false
            }
        } else {
            return false
        }
    }

    return len(new_s) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	// var lcp []rune
	// base case
	if len(strs) == 0 {
		return ""
	}
	prefix_str := strs[0]
	for _, str := range strs {
		for strings.HasPrefix(str, prefix_str) == false {
			if len(prefix_str) == 0 {
                return ""
            }
			last_index := len(prefix_str) - 1
			prefix_str = prefix_str[:last_index]
		}
	}
	return prefix_str
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	// Only care about 9 that + 1 carry to the left
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
            digits[i] = digits[i] + 1 // normal case
            return digits
        }
        digits[i] = 0  // if digit is 9
    }
	return append([]int{1}, digits...) // every element is 9 case add the 1 to the top of digits [0,0,0....]
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	// fast & slow two pointers
	if len(nums) == 0 {
        return 0
    }
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow += 1
			nums[slow] = nums[fast]
		}
    }
    return slow + 1
}


// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	// [[1, 4], [3, 6]] -> [[1, 6]] because 4 >= 3; [[1, 2], [3, 6]] -> [[1, 2], [3, 6]] because 2 < 3
	// Edge case
	if len(intervals) <= 1 {
		return intervals
	}

	// sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        last := result[len(result)-1]

        if current[0] <= last[1] {
            // duplicate case
            if current[1] > last[1] {
                result[len(result)-1][1] = current[1]
            }
        } else {
            // no duplicate case
            result = append(result, current)
        }
    }
    return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// Hash map, store each number in the map for searching the matching case
    m := make(map[int]int)

    for i, num := range nums {
        needed := target - num
        // check needed number
        if prevIndex, ok := m[needed]; ok {
            // found case
            return []int{prevIndex, i}
        }
        // not found case
        m[num] = i
    }
    return nil
}
