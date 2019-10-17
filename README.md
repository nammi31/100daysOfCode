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

