package main

// import (
// 	"io"
// 	"os"
// )

// import (
// 	// "os"
// 	// "bytes"
// 	// "fmt"
// 	"io"
// 	"net"
// 	"os"
// )

// func main() {
// 	// file, err := os.Create("test.txt")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// file.Write([]byte("os.File example\n"))
// 	// file.Close()

// 	// os.Stdout.Write([]byte("os.Stdout example\n"))

// 	// var buffer bytes.Buffer
// 	// buffer.Write([]byte("bytes.Buffer example\n"))
// 	// buffer.Write([]byte("bytes.Buffer example2\n"))
// 	// fmt.Println(buffer.String())
// 	conn, err := net.Dial("tcp", "ascii.jp:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
// 	io.Copy(os.Stdout, conn)
// }
// import (
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("http.ResponseWriter sample"))
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	file, err := os.Create("multiwriter.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := io.MultiWriter(file, os.Stdout)
// 	io.WriteString(writer, "io.MultiWriter example\n")
// }

import (
	"encoding/json"
	"os"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
