package main

import (
	"fmt"
)

func main() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php", "application/x-www-form-urlencoded", 
							strings.NewReader("name=cjb"))
	if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(body))
}
