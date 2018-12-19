package main
import (
    "fmt"
    "sort"
    "os"
    "io"
    "bufio"
    "strings"
    "strconv"
)

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
        }
    }
    return -1
}

func readFile() map[string] string {
    var per=make(map[string] string)
    var arr []string
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
        per[arr[0]]=arr[1]
    }
    return per
}

const threads=10
var succ =make(chan int)

func main() {
    
    defer close(succ)

    r :=readFile()
    tmp :=[10]Person{}
    per := []Person {}

    for k,v := range r {
        tmp[0].Name=k
        tmp[0].Age,_=strconv.Atoi(strings.Replace(v,"\n","",-1))
        per=append(per,tmp[0])
    }
    //p(per)
    //sort.Sort(Byage(per))
    //p(per)
    //p(rank(59361974761705217,Byage(per)))
    
    len_data := len(per)

    step := len_data/threads

    if step == 0 {
        step=len_data
    }

    p(step)
    p(len_data)

    for i:=0; i<=len_data-step;i=i+step {
        go func(data Byage) {
            sort.Sort(Byage(data))
            //p("start:",data)
            ret :=rank(59361974761705217,Byage(data))
            succ <- ret
        }(per[i:i+step])
    }

    c:=done() 
    p(c)
}
