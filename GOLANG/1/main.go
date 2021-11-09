package main

func main() {
	a := App{}
	// a.Initialize("root", "zymr@123", "test")
	a.Initialize("postgres", "tatva972000", "postgres")
	a.DB.AutoMigrate(&user{})
	a.Run(":8080")
}
