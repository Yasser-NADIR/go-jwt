package main

import (
	"goJWT/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {
}
