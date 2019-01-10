package main
import (
    "fmt"
    "math"
)

/*
revers strings
*/
func m(s string) string{

    N := len(s)
    if N <= 1 {
        return s
    }

    a := s[0:N/2]
    b := s[N/2:N]
    return m(b)+m(a)
}

type mi []int

/*
    目标数字由那两个元素加和求得，同一元素仅可使用一次
*/
func twoSumV2(nums []int,target int) [] int {

    var m=make(map[int] int)
    var diff []int

    for i:=0; i<len(nums); i++ {
        m[nums[i]]=i
        diff=append(diff,target-nums[i])
    }

    for i:=0;i<len(diff);i++ {
        if _,ok :=m[diff[i]];ok && i != m[diff[i]]{
            return []int{i,m[diff[i]]}
        }
    }
    return []int{-1,-1}
}

/*同上*/
func twoSum(nums []int,target int) []int {

    i :=0
    j :=0
    l:=len(nums)
loopfor:
    for i =0;i<l; i++ {
        firstone := target-nums[i]
        for j=i;j<l;j++ {
            if firstone == nums[j] {
                break loopfor
            }
        }
    }

    if i >=l || j >= l {
        i = -1
        j = -1
    }
    //ret := []int{i,j}
    return []int{i,j}
}

func mod (x int) int {
    return x%10
}
/*
    数字反转
*/
func sub ( x int) int {

    var array []int
    var sum int
    for i:=0;;i++{
        if int(math.Abs(float64(x))) <= 9 {
            array=append(array,x)
            break
        }else {
            array=append(array,mod(x))
        }
       x=x/10
    }

    l:=len(array)
    for i:=0;i<len(array);i++ {
        sum += array[i]*int(math.Pow10(l-1))
        l--
    }
    if sum >int(math.Pow(2,31)-1) || sum < -int(math.Pow(2,31)) {
        sum = 0
    }
    return sum
}

/*
    判断回型数字
*/
func isPalindrome ( x int) bool {
    
    if x <0 {
        return false
    } 

    if x==sub(x) {
        return true
    }
    return false
}

/*
    数字反转 V2
*/
func sub2(x int ) int {
    
    if x >=0 && x <=9 {
        return x
    }

    rev := 0

    for x != 0 {
        rev = rev*10 + x%10
        x=x/10
    }
    if rev >int(math.Pow(2,31)-1) || rev < -int(math.Pow(2,31)) {
        rev = 0
    }
    return rev
}
func main() {
    //nums :=[]int{-1,-2,-3,-4,-5}
    //fmt.Println(twoSumV2(nums,-8))
    fmt.Println(sub2(-123))
}
