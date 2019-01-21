package main
import (
    "fmt"
    "math"
    "sort"
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

func romanToInt(s string) int {

    roman := map[string] int {
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
        "IV":4,
        "IX":9,
        "XL":40,
        "XC":90,
        "CD":400,
        "CM":900,
    }

    var max,begin,sum int

    for i :=1 ; i <= len(s) ;{
        //fmt.Println("begin:",begin)
        //fmt.Println("i:",i)
        if k,ok :=roman[s[begin:i]];ok && k > max {
            max = roman[string(s[begin:i])]
            //fmt.Println(s[begin:i])
            i++
        } else {
            begin=i-1
            sum += max
            max =0
        }
    }
    return sum+max
}

/*
    找出最长的共同前缀
*/

func min(x,y int) int {
    if x > y {
        return y
    }
    return x
}
func startWithKey( s string,key string) bool {
    return key == s[0:len(key)] 
}
func longestCommonPrefix(strs []string) string {
      
    var m int
    var mp = make(map[int] string)
    if len(strs) > 0 {
        m = len(strs[0])
        mp[len(strs[0])]=strs[0]
    }else {
        return ""
    }
    for i:=1;i < len(strs); i++ {
        m = min(len(strs[i]),m)
        mp[len(strs[i])]=strs[i]
    }

    cn := m
    for i:=0;i<len(strs);i++ {
        if cn>0 && !startWithKey(strs[i],mp[m][0:cn]) {
            cn--
            i=-1 //确保i从0开始
        }
    }

    return mp[m][0:cn]
}

/*
    
*/
func isValid(s string) bool {

    mp := map[string] string{
        "(":")",
        "[":"]",
        "{":"}",
    }


    s_len :=len(s)
    var s_cli []string
    var v string

    if s_len % 2 != 0 {
        return false
    }
    for i:=0; i<s_len;i++ {
        if _,ok := mp[string(s[i])]; ok {
            s_cli=append(s_cli,string(s[i]))
        }else {

            if len(s_cli) <=0 {
                return false
            }

            v = s_cli[len(s_cli)-1]
            if  mp[v] != string(s[i]){
                return false
            }
            s_cli=s_cli[:len(s_cli)-1]
        }
    }
    return len(s_cli)==0
}
/*
     合并两个有序单链表，合并后的新链表依然有序
*/
type ListNode struct {
    Val int
    Next *ListNode
}

//func mergeTwoLists(l1,l2 *ListNode) *ListNode {
//
//}

/*
    移出重复的元素
*/
func removeDuplicates(nums []int) int {

    nl :=len(nums)
    if nl == 0 {
        return nl
    }
    
    cur := 0
    for i:=0;i<nl;i++ {
        
        if i ==nl-1 {
            nums[cur]=nums[i]
            cur++
            break
        }
        if nums[i] != nums[i+1] {
            nums[cur]=nums[i]
            cur++
        }
    }
    return cur
}
/*
    移出指定元素值
*/
func removeElement(nums []int, val int) int {

    k :=0
    for i:=0;i<len(nums);i++ {
        if nums[i] == val {
            continue
        }
        nums[k]=nums[i]
        k++
    }
    return k
}
/*
    查找指定字符串的位置，并返回数组下标
*/
func strStr(haystack string, needle string) int {

    l_needle := len(needle)
    l_haystack := len(haystack)

    if l_needle == 0 {
        return 0
    }

    if l_haystack < l_needle {
        return -1
    }


    startKey :=needle[0]

    for i :=0;i<l_haystack;i++ {
        if haystack[i] == startKey {
            if i+l_needle > l_haystack {
                return -1
            }
            if haystack[i:i+l_needle] == needle {
                return i
            }
        }
    }
    return -1
}
/*
    返回指定元素的位置，如果不存在，则返回其插入数组后的下标（排序）
*/
func searchInsert(nums []int, target int) int {

    end := len(nums)-1

    switch {
        case end < 0 :
            return 0
        case nums[0] >= target :
            return 0
        case nums[end] == target :
            return end
        case nums[end] < target :
            return end+1
    }

    for i:=0;i<end;i++ {
        if nums[i] == target {
            return i
        }
        if nums[i] < target && nums[i+1] > target {
            return i+1
        }
    }

    return 0
}

func v2 (nums []int, target int) int { 
    var i int
    for i=0 ; i<len(nums); i++ {
        if nums[i] >= target {
            return i
        }
    }
    return i+1
}

func toLowerCase ( str string) string {

    b :=[]byte(str)
    for i:=0;i<len(str);i++ {
        if str[i] >=65 && str[i] <=90 {
            b[i]=b[i]+32
        }
    }
    return string(b)
}
/*
    在石头中找宝石
*/

func numJewelsInStones(J string, S string) int {
    mp :=make(map[byte] int)
    cn :=0

    for i:=0;i<len(J);i++ {
        mp[J[i]]=i
    }

    for i:=0;i<len(S);i++ {
        if _,ok :=mp[S[i]];ok {
            cn++
        }
    }
    fmt.Println(mp)
    return cn
}

/*
    查找字符位置
*/
func findStr( s string,key byte) int {
    for i:=0;i<len(s);i++ {
        if s[i] == key {
            return i 
        }
    }
    return -1
}
/*
    找出正确且去重后的邮箱地址
*/

func numUniqueEmails(emails []string) int {

    mp := make (map[string] int)

    for i:=0;i<len(emails);i++ {
        em := []byte(emails[i])
        k :=0
        findDot := 0
        findPlus := false
        for _,val := range emails[i] {
            if val == '@' {
                findPlus =false
                findDot++
            }
            if val == '.'||findPlus {
                if findDot == 0 {
                    continue
                }
            }
            if val == '+' {
                findPlus =true
                continue
            }

            em[k]=byte(val)
            k++
        }
        if k > 0 && findDot == 1{
            mp[string(em)[0:k]]++
        }
    }   

    return len(mp)
}
/*
    描述很复杂...
*/
func repeatedNTimes(A []int) int {
    
    //mp := make(map[int] int)
    l := len(A)
    var mp [10000]int
    for i:=0;i < l; i++ {
        mp[A[i]]++
        if mp[A[i]] > 1 {
            return A[i]
        }
    }

    return 0
}
/*
    摩斯密码映射
*/
func uniqueMorseRepresentations(words []string) int {
    M := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}   

    morse_mp := make(map[string] int)

    for i:=0; i < len(words); i++ {
        morse :=""
        for _,v := range words[i] {
            morse = morse+M[v-'a']       
        }
        morse_mp[morse]++
    }
    return len(morse_mp)
}
/*
    偶数排在奇数前
*/

func sortArrayByParity(A []int) []int {
    
    var B [5000]int

    l := len(A)
    start := 0
    end := l-1
    for i:=0; i<l; i++ {
        if A[i] % 2 ==0 {
            B[start]=A[i]
            start++
        }else {
            B[end]=A[i]
            end--
        }
    }
    return B[0:l]
}
/*
    反转二维数组
*/
func invertArray( a []int) []int {
    for i:=0;i<len(a);i++ {
        a[i]=a[i]^1
    }
    return a
}

func reverseArray( a []int) []int {

    l := len(a)
    for i:=0;i<l/2;i++{
        tmp := a[i]
        a[i]=a[l-i-1]
        a[l-i-1]=tmp
    }
    return a
}
func flipAndInvertImage(A [][]int) [][]int {
    
    for i:=0;i<len(A);i++ {
        reverseArray(A[i])
        invertArray(A[i])
    }
    return A
}
/*
    是否回到原点
*/
func judgeCircle(moves string) bool {
    
    var k,h int
    if len(moves)%2 != 0 {
        return false
    }
    for i:=0;i<len(moves);i++ {
        switch {
            case moves[i] == 'U':
                k++
            case moves[i] == 'D':
                k--
            case moves[i] == 'R':
                h++
            case moves[i] == 'L':
                h--
        }
    }
    return k==0&&h==0
}
/*
    返回非排序列个数
*/
func minDeletionSize(A []string) int {

    stringLen := len(A)
    cn := 0
    for i:=0;i<len(A[0]);i++ {
        for k:=0;k<stringLen-1;k++ {
            if A[k][i] > A[k+1][i] {
                cn++
                break
            }
        }
    }
    return cn
}

/*
    在信息论中，两个等长字符串之间的汉明距离（英语：Hamming distance）是两个字符串对应位置的不同字符的个数。换句话说，它就是将一个字符串变换成另外一个字符串所需要替换的字符个数
*/
func max(x,y int) int {
    if x < y {
        return y
    }
    return x
}

func hammingDistance(x int, y int) int {
    
    mi := min(x,y)
    mx := max(x,y)
    cn :=0
    MAX := int(math.Pow(2,31))
    k := mx -mi

    for i:=0;i<MAX;i-- {
        k = k-2
        if k <=0 {
            break
        }
    }
    return cn
}

/*
    元素平方后排序
*/

type bySquare []int

func (a bySquare) Len() int {
    return len(a)
}

func (p bySquare) Less(i,j int) bool {
    return p[i]*p[i] <p[j]*p[j]
}

func (p bySquare) Swap (i,j int) {
    p[i],p[j] = p[j],p[i]
}
func sortedSquares(A []int) []int {
    for k,_ := range A {
        A[k]=A[k]*A[k]
    }
    sort.Sort(bySquare(A))   
    return A
}
func main() {
    em :=[]int {}

    fmt.Println(sortedSquares(em))
}
