package main

import (
	"fmt"
	"sort"
)	

func main() {

	// Slices + there are like array but little different and mostly used in go

	var anime_list = []string{"bleach", "attack on titians", "naruto"}
	fmt.Printf("type of anime list is %T\n", anime_list)

	anime_list = append(anime_list, "Death Note", "Chainsaw man")
	fmt.Println(anime_list)
	// Here the appened  will add the list , it takes 2 arguments 1st one the array list you want to append,
	// the athings you want to add
	anime_list = append(anime_list[:3])
	fmt.Println(anime_list)
	// the [:] help us to slice the  position you have passed and it is non inclusive

	animeScore := make([]int, 4)

	animeScore[0] = 121
	animeScore[1] = 232
	animeScore[2] = 232

	fmt.Println(animeScore)

	animeScore = append(animeScore, 33, 434, 333)
	fmt.Println(animeScore)

	sort.Ints(animeScore)
	fmt.Println(animeScore)

	// how to remove a value from a slices based on index

	var courses = []string{"wew", "hey", "wow", "humph"}
	fmt.Println(courses)

}
