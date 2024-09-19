package main

func main() {

	// // Get the list of pools
	// pools := getPoolList("https://montreal.ca/lieux?mtl_content.lieux.installation.code=PISI&mtl_content.lieux.available_activities.code=ACT0")

	// // Print the list of pools
	// for _, pool := range pools {
	// 	fmt.Println(pool)
	// }

	url := "https://montreal.ca/lieux/piscine-schubert"

	getPoolSchedule(url)

}
