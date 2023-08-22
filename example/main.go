package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/jnka9755/go-05SDKUSER/user"
)

func main() {

	userTrans := userSdk.NewHttpClient("http://localhost:5001", "")

	user, err := userTrans.Get("446120f0-e6da-4d6a-9ec1-746f69db3sd")

	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal server error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
