package main

func main() {
	router := InitializedUser()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
