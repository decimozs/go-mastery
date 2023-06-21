package main

import "fmt"

func main(){
	fmt.Println("Maps")
	/*
		Maps are also used to stored data with a key and value pairs
		maps do not allow duplicated data, and the default value of the
		map is nil
	*/

	// ways to create map 
	// explicit way
	var explicitMp = map[int]string{
		9 : "Nissan",
		1 : "Honda",
		3 : "Mazda",
		10 : "Lamborghini",
		1002 : "Ferrari",
	}
	// inferred way
	infferedMp := map[int]string{
		9 : "Nissan",
		1 : "Honda",
		3 : "Mazda",
		10 : "Lamborghini",
		1002 : "Ferrari",
	}
	// using make function
	makeMp := make(map[int] string)
	makeMp[9] = "Nissan"
	makeMp[1] = "Honda"
	makeMp[3] = "Mazda"
	makeMp[10] = "Lamborghini"
	makeMp[1002] = "Ferrari"

	// in terms of creating a empty map we should use make function to create a map
	emptyMp1 := make(map[int] string)
	var emptyMp2 map[int] string

	fmt.Println("Ways to create a map")
	fmt.Println(explicitMp)
	fmt.Println(infferedMp)
	fmt.Println(makeMp)
	fmt.Println("\nWays to create empty map")
	fmt.Println(emptyMp1 == nil) 			// this will return false
	fmt.Println(emptyMp2 == nil)			// this will return true

	// accesing elements in the map
	fmt.Println("\nAccesing elements in the map")
	fmt.Println(makeMp[1002])		// accesing the elements in the makeMp map with the key of 1002
	fmt.Println(makeMp[5])			// if the key is not in the map it will return to nil

	// updating elements in the map
	fmt.Println("\nUpdating elements in the map")
	makeMp[1002] = "Subaru"			// mapName[key] = value(change the value)
	fmt.Println(makeMp)

	// adding elements in the map
	fmt.Println("\nAdding elements in the map")
	makeMp[1003] = "Bugatti"			// mapName[key(unique key value)] = value(you want to add)
	fmt.Println(makeMp)

	// removing elements in the map
	fmt.Println("\nRemoving elements in the map")
	delete(makeMp, 1003)					// delete(mapName, key(the key you want to delete))
	fmt.Println(makeMp)

	// check for specific elements in the map
	fmt.Println("\nCheck for specific elements in the map")
	val1, ok1 := makeMp[9]			// checking for existing key and its value
	val2, ok2 := makeMp[6]			// checking for the non existing key and its value
	val3, ok3 := makeMp[1]			// checking for existing key and its value
	_, ok4 := makeMp[1002]			// checking only for existing key and not its value
	fmt.Println("After checking the specific elements in the map")
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	/*
		[note] that maps are references to hashtables
		if you change the content of element in a maps it will also affect the content of the other
	*/

	fmt.Println("\nMap References")
	// orignal map 
	map1 := map [int] string{
		1 : "Nissan",
		2 : "Honda",
		3 : "Mazda",
	}

	// reference map2 to map1
	map2 := map1

	// prin the properties of the map1 and map2
	fmt.Println(map1)
	fmt.Println(map2)

	// change the content value of the map2 index 2 to ferrari
	map2[2] = "Ferrari"

	fmt.Println("After changing the content in map2 that reference to map1")
	fmt.Println(map1)
	fmt.Println(map2)

}
