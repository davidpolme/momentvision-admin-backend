package config

import (
	"fmt"
	"os"	

)

var (	
	PORT = os.Getenv("PORT")
)

func init() {
	if PORT == "" {
		PORT = "8000"
	}
	fmt.Println("PORT = ", PORT)
}
