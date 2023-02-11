package main

import "os"

func main() {
	NewDestinyAPISession(os.Getenv("D2_API_KEY"))
}
