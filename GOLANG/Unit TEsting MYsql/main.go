package main

func main() {
	a := App{}
	a.Initialize("root", "zymr@123", "test")

	a.Run(":8080")
}
