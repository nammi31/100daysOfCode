# 100DaysOfCoding

Today is october 17 and I am going to start my 100DaysOfCoding Challenge.
## Day3 31 oct, 2019
### 1221. Split a String in Balanced Strings
```
func balancedStringSplit(s string) int {
    cnt := 0
    ans := 0
    for i:=0; i<len(s);i++{
        if s[i] == 'R'{
            cnt++
        }else{
            cnt--
        }
        if cnt == 0{
            ans++
        }
    }
    return ans
}
```
* Runtime:  0 ms, faster than 100.00% O(N)
* Memory Usage: 2 MB, less than 100.00% O(1)
### 657. Robot Return to Origin
```
func judgeCircle(moves string) bool {   
    x :=0
    y :=0
    for i:=0; i<len(moves);i++{
        if moves[i] == 'U'{
            y++
        }else if moves[i] == 'D'{
            y--
        }else if moves[i] == 'R'{
            x++
        }else if moves[i] == 'L'{
            x--
        }
    }
    return x == 0 && y == 0
    
}
```
* Runtime:  4 ms, faster than 85.90% O(N)
* Memory Usage:3.5 MB, less than 100.00% O(1)


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
		if len(strs[j]) < len(ans){       // choose a string of minimum length as ans
			ans = strs[j]
		}
	}        
	for i:=0; i<len(ans);i++{                 // assume all chars of ans match
		for j:=0; j<len(strs);j++{
			if ans[i] !=  strs[j][i]{
				return ans[0:i]   // if anyone does not match then we no need to explore more
			}
		}
	}
	return ans                                // if nothing is returned yet then all chars has matched
}
```
* Runtime: 0 ms, faster than 100.00% 
* Memory Usage:  2.4 MB, less than 100.00%
## Day2
### 136. single number
```ha
func singleNumber(nums []int) int { 
    for i:=1; i<len(nums); i++{
		nums[i] = nums[i]^nums[i-1]     // current num update with xor of all previous num in array
	}
    return nums[len(nums)-1]                    // so last one is our final result
}
```
* Runtime: 8 ms, faster than 95.40% O(N)
* Memory Usage:  4.7 MB, less than 100.00% O(1)


