package main
import (
    "fmt"
    "sort"
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
    p(e)
    for s <= e {
        var mid = s + (e-s)/2
        p(key,i[mid].Age)
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

func main() {
    
    per := []Person {
        {"A",31},
        {"B",20},
        {"C",48},
        {"D",18},
        {"E",19},
        {"F",122},
    }
    sort.Sort(Byage(per))
    p(per)
    p(rank(31,Byage(per)))
}
