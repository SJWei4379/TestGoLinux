package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello, world\n")
	fmt.Fprintf(w, "恭喜！项目运行成功！！！")
	//io.WriteString(w, "Hello, world!\n")
	fmt.Println("成功通过ip访问！！！")
}

func main() {

	fmt.Println("成功部署到服务器！")
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":8888", nil)

	if err != nil {

		panic(err)

	}

}
