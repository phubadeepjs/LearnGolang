package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon web services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon web services"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
