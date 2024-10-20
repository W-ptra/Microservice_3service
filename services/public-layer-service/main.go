package main

import public_layer_api "github.com/W-ptra/microservice_3service/public-layer-service/api"

// import (
// 	"net/http"
// 	"fmt"
// 	"io"
// )

func main() {
	public_layer_api.Run()
	//url := fmt.Sprintf("%v:%v/users?pageNum=%v&pageSize=%v",os.Getenv("USER_SERVICE_HOST"),os.Getenv("USER_SERVICE_PORT"),pageNum,pageSize)
	//fmt.Println(url)
	// resp,_ := http.Get("http://localhost:2000/users")
	// defer resp.Body.Close()

	// body,_ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
}