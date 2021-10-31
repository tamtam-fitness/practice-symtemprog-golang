// package main

// import (
// 	"io"
// 	"net"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "ascii.jp:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
// 	io.Copy(os.Stdout, conn)
// }

package main

import (
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}

	defer oldFile.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}

	defer newFile.Close()
	io.Copy(newFile, oldFile)
}
