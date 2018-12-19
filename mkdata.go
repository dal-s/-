package main
import (
    "fmt"
    "math/rand"
    "time"
    "os"
)
/*
生成系列随机数据，用于程序测试
*/
var letters= []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
var src = rand.NewSource(time.Now().UnixNano())


func randSeq(n int) string {
    b :=make([]rune,n)
    for i:=range b {
        b[i]=letters[src.Int63() % int64(len(letters))]
    }
    r := rand.New(src)
    s := fmt.Sprintf("%s %d\n",string(b),r.Intn(99999999999999999))
    return s
}

func main() {
    f,err := os.Create("/home/1.dat")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    for i:=0;i<=9999999;i++{
        f.WriteString(randSeq(10))
    }
    f.Sync()
}
