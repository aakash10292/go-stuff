package main
import "fmt"

func main(){
    // Declaring a variable e1 based on a literal struct type
    // We're using var so obviously initialisation it with zero-value
    var e1 struct {
        flag bool
        counter int16
        pi float32
    }
    
    fmt.Printf("%+v\n",e1)
}
