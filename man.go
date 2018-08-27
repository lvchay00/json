package main

import (
	"encoding/json"
	"fmt"
)


type User struct {
    Name    string
    Age     int
    Roles   []string
    Skill   map[string]float64
}

func main() {

    skill := make(map[string]float64)

    skill["python"] = 99.5
    skill["elixir"] = 90
    skill["ruby"] = 80.0

    user := User{
        Name:"rsj217",
        Age: 27,
        Roles: []string{"Owner", "Master"},
        Skill: skill,
    }
//生成json
    rs, err := json.Marshal(user)
    if err != nil{
       fmt.Println(err)
    }
    fmt.Println(string(rs))
//解析json
      jsonToStruct := new(User)
       err  = json.Unmarshal([]byte(`{"Name":"rsj217","Age":27,"Roles":["Owner","Master"],"Skill":{"elixir":90,"python":99.5,"ruby":80}}`),jsonToStruct)
     fmt.Printf("%v",jsonToStruct)

}
/* package main

import (
    "encoding/json"
    "fmt"
)
type Animal struct {
    Name  string
    Order string
}
func main() {

    var animals Animal
    err := json.Unmarshal([]byte(`{"Name": "Platypus", "Order": "Monotremata"}`), &animals)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", animals)
} */