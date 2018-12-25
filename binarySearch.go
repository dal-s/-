package main
import (
    "fmt"
    "sort"
    "os"
    "io"
    "bufio"
    "strings"
    "strconv"
    "sync/atomic"
    "time"
    "reflect"
    "runtime"
)
/*
Function:实现多线程二分法查找算法

TODO-LIST:
1.多线程读取数据
2.done函数在全文查找失败时会继续阻塞 [Done]
3.当其中一个线程返回查找结果时，通知其他线程结束运行
4.排序算法优化
*/

var p = fmt.Println
type Person struct {
    Name string
    Age int
}

func (p Person) string() string{
    return fmt.Sprintf("%s: %d",p.Name,p.Age)
}
type Byage []Person

func (a Byage) Len() int { 
    return len(a) 
}
func (a Byage) Swap(i, j int){ 
    a[i], a[j] = a[j], a[i]
}
func (a Byage) Less(i, j int) bool { 
    return a[i].Age < a[j].Age
}

type ReadFunc func() [] Person

func getFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func runTime(f ReadFunc ) ReadFunc {
    return func() []Person {
        defer func(t time.Time) {
            fmt.Printf("--- Time Elapsed (%s):%v ---\n",getFunctionName(f),time.Since(t))
        }(time.Now())

        return f()
    }
}

//func rank(key int,i []int) int {
//func rank(key string,i []string) int {
func rank(key int,i Byage) int {
    var s=0
    var e=len(i)-1
    for s <= e {
        var mid = s + (e-s)/2
        if   key < i[mid].Age {
            e=mid-1
        }else if key > i[mid].Age {
            s=mid+1
        }else {
            return mid
        }
    }
    return -1
}

func done() int {
    
    for {
        select {
            case c:= <-succ:
                if c != -1 {
                    return c
                }
            default :
                if atomic.LoadUint64(&finish_num) == threads{
                    return -1
                }
        }
    }
}

func readFile() []Person {
    var arr []string
    tmp :=[10]Person{}
    data := []Person {}

    f,err := os.Open("/home/mart_coo/lijianping1/1.dat")
    if err != nil {
        p("read file error! ",err)
    }

    defer f.Close()
    br := bufio.NewReader(f)
    for {
        a,err := br.ReadString('\n')
        if err == io.EOF {
            break
        }else if err != nil {
            panic(err)
        }

        arr =strings.Split(a," ")

        tmp[0].Name=arr[0]
        tmp[0].Age,_=strconv.Atoi(strings.Replace(arr[1],"\n","",-1))
        data=append(data,tmp[0])
    }
    return data
}

const threads=3
var succ =make(chan int)
var finish_num uint64

func main() {
    
    defer close(succ)

    per:=runTime(readFile)() //7s

    //r :=readFile()

    t:=time.Now()
    sort.Sort(Byage(per)) //5s
    p("Sort use:",time.Since(t))

    //p(per)
    //p(rank(32939,Byage(per)))
    
    len_data := len(per)

    step := len_data/threads

    if step == 0 {
        step=len_data
    }

    for i:=0; i<=len_data-step;i=i+step {
        go func(data Byage) {
            sort.Sort(Byage(data))
            ret :=rank(1,Byage(data))
            succ <- ret
            atomic.AddUint64(&finish_num,1)
        }(per[i:i+step])
    }

    c:=done() 
    p(c)
}
