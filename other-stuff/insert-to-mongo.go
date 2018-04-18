// This was ripped from the web but successfully tested
// from my local machine with a 2-node mini replica set.
//
// This program assumes you have a mongo cluster, the names
// for which are declared in "func main()".
//
// This program also requires the installation of mgo.
//
//	You can get it by typing:
//	go get gopkg.in/mgo.v2
//
// Expected results:	Phone: +55 53 8116 9639
//
// Comments:	To be filled in.
//

package main
  
import (
        "fmt"
        "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Phone string
}

func main() {
        session, err := mgo.Dial("127.0.0.1:27017,127.0.0.1:37017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
                       &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
}
