# 100DaysOfCoding

Today is october 17 and I am going to start my 100DaysOfCoding Challenge.


## Day1
### 1108. Defanging an IP Address


```
func defangIPaddr(address string) string {
    return strings.Replace(address, ".", "[.]",len(address)/2)
}
```
* Runtime: 0 ms, faster than 100.00% 
* Memory Usage: 1.9 MB, less than 100.00%

### 771. Jewels and Stones


```
func numJewelsInStones(J string, S string) int {
        var inJewel [128]bool
	for i:=0; i<len(J); i++{
	     inJewel [ J[i]-65 ] = true
	}
	ans:=0
	for i:=0; i<len(S);i++{
		if inJewel [S[i]-65] == true{
			ans++ 
	        }
	}
	return ans
}
```
* Runtime: 0 ms, faster than 100.00%
* Memory Usage:  2 MB, less than 100.00%
### 14. Longest Common Prefix


```
func min(x int,y int)int{
	if x<y {
		return x
	}
	return y
}
func longestCommonPrefix(strs []string) string {
	if len(strs)==0{
	  return ""
	}
	if len(strs) ==1{
		return strs[0]
	}
	ans := strs[0]        
	for j:=1; j<len(strs);j++{
		if len(strs[j]) < len(ans){       //choose a string of minimum length as ans
			ans = strs[j]
		}
	}        
	for i:=0; i<len(ans);i++{                 //	assume all chars of ans match
		for j:=0; j<len(strs);j++{
			if ans[i] !=  strs[j][i]{
				return ans[0:i]   //if anyone does not match then we no need to explore more
			}
		}
	}
	return ans                                //if nothing is returned yet then all char match
}
```
* Runtime: 0 ms, faster than 100.00% 
* Memory Usage:  2.4 MB, less than 100.00%
## Day2
### 136. single number
```
func singleNumber(nums []int) int { 
    for i:=1; i<len(nums); i++{
		nums[i] = nums[i]^nums[i-1]
	}
    return nums[len(nums)-1]
}
```
* Runtime: 8 ms, faster than 95.40% O(N)
* Memory Usage:  4.7 MB, less than 100.00% O(1)

