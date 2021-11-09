package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("postgres", "tatva972000", "postgres")

	a.Run(":8080")
}
