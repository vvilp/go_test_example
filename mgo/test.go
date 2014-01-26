package main 

import "fmt"
import "labix.org/v2/mgo"

func main() {
	fmt.Println("test mgo")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)		
	}
	defer session.Close()

	c := session.DB("mydb").C("testData")

	fmt.Println(c.Find({}))

}