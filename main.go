package main

func main() {
	router := InitializedApp()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

	//PZN Structure Tech Approach
	// controller -> job_controller, user_controller, etc
	// service -> job_service, user_service, etc
	// repository -> job_repository, user_repository, etc
}
