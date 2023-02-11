package main

import (
	"fmt"
	"os"

	d2 "github.com/Mercwri/mercwri-armourer/internal/api"
)

func main() {
	sess := d2.NewDestinyAPISession(os.Getenv("D2_API_KEY"))
	user, err := sess.GetDestinyUser(os.Getenv("D2_USER"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
