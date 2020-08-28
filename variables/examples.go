package main

/*
Type ,  decleration , implicit type
Scope Visibility -> lower scope gets precedence

No unused variable is not allowed
*/

import (
	"fmt"
)

/*
here full decleration compulsory (declared and limited to the scope of package main)
*/
var o int = 30

var Itest int = 100

/*
 to export globally first lettter upper case must
*/

/*
 Block declarations  (for keeping things cleaner)
*/
var (
	actorName string = "SSR"
	movie     string = "kai po che"
)

func main() {
	var i int      // decleration and type
	i = 42         // assignement
	var j int = 35 // short way
	k := 24        // implict type declaration
	fmt.Println(" i = ", i, " j= ", j, " k= ", k, " o= ", o, " Itest= ", Itest)
	fmt.Println(" actor = ", actorName, " movie =", movie)

	l := 99
	n := 91. // by default considers it float64 type

	fmt.Printf(" Value %v , Type %T ", l, l)
	fmt.Printf(" Value %v , Type %T \n", n, n)

}
