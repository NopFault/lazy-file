package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const CHUNKSIZE = 2048

func fillString(retunString string, toLength int) string {
	for {
		strlen := len(retunString)
		if strlen < toLength {
			retunString = retunString + "."
			continue
		}
		break
	}
	return retunString
}

func download(host string, port int) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("[*] Connected")

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	conn.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), "."), 10, 64)

	conn.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ".")

	newFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer newFile.Close()
	var rcBytes int64

	fmt.Println("[*] Starting Download")
	fmt.Print("Downloading: |")
	for {
		fmt.Print(".")

		if (fileSize - rcBytes) < CHUNKSIZE {

			fileData := make([]byte, (rcBytes+CHUNKSIZE)-fileSize)
			decBytes := make([]byte, (rcBytes+CHUNKSIZE)-fileSize)

			conn.Read(fileData)
			for i := 0; i < len(fileData); i++ {
				decBytes[i] = fileData[i] ^ byte(fileName[0])
			}

			io.CopyN(newFile, bytes.NewReader(decBytes), (fileSize - rcBytes))
			break
		}

		fileData := make([]byte, CHUNKSIZE)
		decBytes := make([]byte, CHUNKSIZE)

		conn.Read(fileData)
		for i := 0; i < len(fileData); i++ {
			decBytes[i] = fileData[i] ^ byte(fileName[0])
		}

		io.CopyN(newFile, bytes.NewReader(decBytes), CHUNKSIZE)
		rcBytes += CHUNKSIZE
	}
	fmt.Print("|\n\n")
	fmt.Println("[*] File received completely")
}

func giveFileToClient(filename string, conn net.Conn) {

	defer conn.Close()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[Error] Cant open file:")
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("[Error] Cant get file info:")
		fmt.Println(err)
		return
	}

	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)

	conn.Write([]byte(fileSize))
	conn.Write([]byte(fileName))

	fmt.Println("[*] Start sending file")

	for {

		chunkedBytes := make([]byte, CHUNKSIZE)
		encBytes := make([]byte, CHUNKSIZE)

		_, err := file.Read(chunkedBytes)
		if err == io.EOF {
			break
		}

		for i := 0; i < len(chunkedBytes); i++ {
			encBytes[i] = chunkedBytes[i] ^ byte(fileName[0])
		}

		conn.Write(encBytes)
	}
	fmt.Println("[*] File has been sent")
	fmt.Println("[*] Close connection")
	return
}

func server(host string, port int, file string) {

	server, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("[Error] Starting server: ", err)
		os.Exit(1)
	}

	defer server.Close()

	fmt.Println("[*] Waiting for connections")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("[Error] Accepting connection: ", err)
			os.Exit(1)
		}

		fmt.Println("[*] Client connected")
		go giveFileToClient(file, conn)
	}
}
func main() {
	host := flag.String("h", "0.0.0.0", "Server IP")
	port := flag.Int("p", 0, "Server Port")
	file := flag.String("f", "", "File to share")

	flag.Parse()

	if *port > 0 {
		if len(*file) > 0 {
			server(*host, *port, *file)
		} else {
			download(*host, *port)
		}
	} else {
		fmt.Println("-p [port] is required")
	}

}
