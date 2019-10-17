# 100DaysOfCoding

Today is october 17 and I am going to start my 100DaysOfCoding Challenge.


## Day1
### 1108. Defanging an IP Address


```

func defangIPaddr(address string) string {
    return strings.Replace(address, ".", "[.]",len(address)/2)
}
```



