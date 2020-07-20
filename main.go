package main

func main() {
	// Setup
	router := newRouter()
	router.Logger.Fatal(router.Start(":8000"))
}
