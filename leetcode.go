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
/*
    index in a mountain array
*/

func peakIndexInMountainArray(A []int) int {

    for k,_ := range A {
        if A[k+1] < A[k] {
            return k
        }
    }
    return 0
}
/*
    找出能被自己包含的所有数字取余的数字，如：128 由 1 2 8组成
*/

func findNumber( i int) []int {

    var array []int
    for i !=0 {
        array = append(array,i%10)
        i = i/10
    }
    return array
}

func selfDividingNumbers(left int, right int) []int {

    var array []int
    for i:=left;i<=right;i++ {
        k := i
        isVaild := true
        for k !=0 {
            if num := k%10; num == 0 || i % num != 0{
                isVaild = false
                break
            }
            k = k/10
        }
        if isVaild {
            array = append(array,i)
        }
    }
    return array
}
type byNum []int

func (a byNum) Len() int {
    return len(a)
}

func (p byNum) Less(i,j int) bool {
    return p[i]<p[j]
}

func (p byNum) Swap (i,j int) {
    p[i],p[j] = p[j],p[i]
}
func arrayPairSum(A []int) int {
    sort.Sort(byNum(A))   
    
    sm := 0
    for i:=0; i < len(A); i+=2 {
        sm +=A[i]
    }
    return sm
}
/*
    记录每个人调用ping的次数
*/
type RecentCounter struct {
    curName string
    beforeName []string
    beforeTm []int   
    cn int
    point int
}


func Constructor() RecentCounter {
    return RecentCounter{} 
}

const K int = 3000
func findKey( this *RecentCounter,t int) int {

    fmt.Println("len:",len(this.beforeTm))
    for i:=this.point;i<len(this.beforeTm);i++ {
        if this.beforeTm[i]+K >= t {
            this.point = i
            fmt.Println("i:",i)
            return len(this.beforeTm)-i+1
        }
    }
    return 1
}
func (this *RecentCounter) Ping(t int) int {

    this.cn=1
    for _,v := range this.beforeName {
        if this.curName == v {
            this.cn = findKey(this,t)
            break
        }
    }
    this.beforeName=append(this.beforeName,this.curName)
    this.beforeTm=append(this.beforeTm,t)
    return this.cn
}
/*
     Fibonacci numbers
*/
func fib(N int) int {

    if N == 1 || N == 0{
        return N
    }
    return fib(N-1)+fib(N-2)
}
func fib2(N int) int {
    if N == 0 {return 0}
    if N == 1 {return 1}
    
    a, b := 0, 1
    for i:=2; i<=N; i++ {
        a, b = b, a+b
    }
    
    return b
}
func sortArrayByParityII(A []int) []int {
    
    odd := 1
    even := 0
    var sortArray [20000]int

    for i:=0;i<len(A);i++ {
        if A[i] % 2 ==0 {
            sortArray[even]=A[i]
            even +=2
        }else {
            sortArray[odd]=A[i]
            odd +=2
        }
    }
    return sortArray[0:len(A)]
}

var mp=make(map[string] int)
func splitString( s string) {

    nmk := strings.Index(s," ")
    before_num := s[0:nmk]
    num,_:=strconv.Atoi(before_num)
    firstPoint := strings.Index(s,".")
    lastPoint := strings.LastIndex(s,".")

    mp[s[nmk+1:]]=num + mp[s[nmk+1:]]
    mp[s[firstPoint+1:]]=num + mp[s[firstPoint+1:]]

    if firstPoint != lastPoint {
        mp[s[lastPoint+1:]]=num + mp[s[lastPoint+1:]]
    }
}
func subdomainVisits(cpdomains []string) []string {
    var res []string
    for _,v := range cpdomains {
        splitString(v)
    }    

    for k,v := range mp {
        res = append(res,strconv.Itoa(v)+" "+k)
    }
    return res
}

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
/*
    reverse words in a string
*/
func reverseString2(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}
func reverseWords(s string) string {
    words := strings.Split(s," ") 
    var ans string
    for _,v := range words {
        ans = ans  + " " + reverseString2(v) 
    }
    return ans[1:]
}

func reverseString(s []byte){

    end := len(s) -1
    for start := 0;start <len(s)/2;start++ {
        s[start],s[end] = s[end],s[start]
        end--
    }
}
/*
    Number of Lines To Write String
*/
func numberOfLines(widths []int, S string) []int {
    
    var bts,lines  int
    for _,v := range S {
        bts = bts + widths[v-97]
        if bts == 100 {
            lines++
            bts = 0
        }
        if bts > 100 {
            lines++
            bts = widths[v-97]
        }
    }
    if bts != 0 {
        lines++
    }
    return []int {lines,bts}
}

func findComplement(num int) int {

    bts := []rune(fmt.Sprintf("%b",num))
    cn := 0
    num_len := len(bts)
    for k:=0;k< num_len;k++ {
        cn = cn + int(bts[k] ^ 1 - '0')*int(math.Pow(2,float64(num_len - k -1)))
    }
    return cn
}
func findComplement2(num int) int {
  h := high(num)
  mask := 1 << uint(h) -1
  return mask & (^num)
}

func high(num int) int {
  i:=0 
  for num > 0 {
    i++
    num = num >> 1
  }
  return i
}

/*

*/
func findAlphabet(s string) bool {

    mp := map[rune] int {
        'q':1,'w':1,'e':1,'r':1,'t':1,'y':1,'u':1,'i':1,'o':1,'p':1,
        'a':2,'s':2,'d':2,'f':2,'g':2,'h':2,'j':2,'k':2,'l':2,
        'z':3,'x':3,'c':3,'v':3,'b':3,'n':3,'m':3,
    }

    row := 0
    for k,v := range s {

        if v >=65 && v <=90 {
            v = v+32
        }

        if k ==0 {
            row = mp[v]
        }
        if v,ok := mp[v]; ok && row != v {
            return false
        }
    }
    return true
}
func findWords(words []string) []string {
    
    var ans []string
    for _,v := range words {
        if findAlphabet(v) {
            ans = append(ans,v)
        }
    }
    return ans
}

func uncommonFromSentences(A string, B string) []string {
    allString := append(strings.Split(A," "),strings.Split(B," ")...)   
    ans := allString[:0]
    var mp=make(map[string] int)

    for _,v := range allString {
         
        if _,ok := mp[v] ;ok {
            mp[v]++
        }else{
            mp[v]=1
        }
    }
    for k,v := range mp {
        if v == 1 {
            ans = append(ans,k)
        }
    }
    return ans
}

func islandPerimeter(grid [][]int) int {

    cn := 0
    for line:=0;line<len(grid);line++ {
        for row:=0;row<len(grid[0]);row++ {
            if grid[line][row] == 0 {
                continue
            }

            UP    := !(line == 0 ) //false
            DOWN  := !(line == len(grid) - 1) //false

            LEFT  := !(row == 0 ) //false
            RIGHT := !(row == len(grid[0]) - 1) //false

            tmp := 4

            if UP {
                tmp = tmp - grid[line -1][row]
            }
            if DOWN {
                tmp = tmp - grid[line +1][row]
            }
            if LEFT {
                tmp = tmp - grid[line][row -1]
            }
            if RIGHT {
                tmp = tmp - grid[line][row +1]
            }
            cn = cn + tmp
        }
    }    
    return cn
}

func calPoints(ops []string) int {
    var ans []int
    cn := 0
    for _,v := range ops {
        switch {
            case v == "+" :
                ans = append(ans, ans[len(ans)-1] + ans[len(ans)-2])
            case v == "C" :
                ans = ans[:len(ans)-1]
            case v == "D" :
                ans = append(ans,ans[len(ans)-1]*2)
            default :
                point,_:=strconv.Atoi(v) 
                ans = append(ans,point)
        }
    }
    for _,v := range ans {
        cn = cn +v
    }
    return cn
}

func min(x,y int) int {
    if x > y {
        return y
    }
    return x
}
func distributeCandies(candies []int) int {
    candCount := len(candies)   

    var kinds = make(map[int] int)

    for _,kind := range candies {
        kinds[kind]++
    }

    return min(len(kinds),candCount/2)
}

func fizzBuzz(n int) []string {
    
    var ans []string
    for i:=1;i <= n;i++ {
        mod3 := (i % 3 ==0)
        mod5 := (i % 5 ==0)
        if mod3&&mod5 {
            ans=append(ans,"FizzBuzz")
            continue
        }
        if mod3 {
            ans=append(ans,"Fizz")
            continue
        }   
        if mod5 {
            ans=append(ans,"Buzz")
            continue
        }
        ans = append(ans,strconv.Itoa(i))
    }
    return ans
}

func singleNumber(nums []int) int {
    var singleMp = make(map[int] int)   

    for i:=0;i<len(nums);i++ {
        if _,ok := singleMp[nums[i]];ok {
            delete(singleMp,nums[i])
        }else{
            singleMp[nums[i]]=0
        }
    }

    for k,_ := range singleMp {
        return k
    }
    return 0
}
func matrixReshape(nums [][]int, r int, c int) [][]int {
    
    original_nums := len(nums)*len(nums[0])
    var ans [][]int
    var ans0 []int
    ans_i :=0
    ans_k :=0
    if r*c > original_nums {
        return nums
    }

    for i:=0;i<len(nums);i++{
        for k:=0;k<len(nums[0]);k++{
            if ans_i > r -1{
                break
            }
            ans0 = append(ans0,nums[i][k])
            if ans_k >= c -1 {
                ans = append(ans,ans0)
                ans_i++
                ans_k=0
                ans0 = []int{}
            }else{
                ans_k++
            }
        }
    }
    return ans
}
/*
    自定义排序
*/
func isAlienSorted(words []string, order string) bool {

    var orderMp = make(map[byte] int)

    for k,v := range order {
        orderMp[byte(v)]=k
    }

    wordsLength := len(words)

    for i:=0;i<wordsLength -1;i++ {

        minWord := min(len(words[i]),len(words[i+1]))

        if len(words[i+1]) < len(words[i]) && words[i][:minWord] == words[i+1] {
            return false
        }
        for k:=0;k<minWord ;k++ {
            if orderMp[words[i][k]] == orderMp[words[i+1][k]] {
                continue
            }else if orderMp[words[i][k]] < orderMp[words[i+1][k]] {
                break
            }else {
                return false
            }
        }
    }

    return true
}

/*
    TODO
*/
func isTriangle(c,k,g int) (bool,int) {
    return c<k+g&&k<c+g&&g<c+k,c+k+g
}
func largestPerimeter(A []int) int {
    max := 0
    sort.Sort(byNum(A))
    for i:=0;i<=len(A)-3;i++{
        for j:=i+1;j<=len(A)-2;j++ {
            for k:=len(A)-1;k>=j+1;k-- {
                ok,perimeter := isTriangle(A[i],A[j],A[k])
                if !ok {
                    break
                }
                if perimeter > max{
                    max = perimeter
                }
            }
        }
    }
    return max
}
func main() {
    fmt.Println(largestPerimeter(em))
}

