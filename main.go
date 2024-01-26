package main

func main() {
	router := InitializedApp()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
