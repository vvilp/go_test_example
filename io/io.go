package main

import "fmt"
import "os"
import "bufio"
import "os/exec"

func test_basic_read() {
	file, _ := os.Open("testfile")
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println((buf))
	}
}

func test_buffer_read() {
	file, _ := os.Open("testfile")
	defer file.Close()

	buf_reader := bufio.NewReader(file)
	buf_writer := bufio.NewWriter(os.Stdout)
	defer buf_writer.Flush()

	buf := make([]byte, 1024)
	for {
		n, _ := buf_reader.Read(buf)
		if n == 0 {
			break
		}
		buf_writer.Write(buf[0:n])
	}
}

func test_buffer_read_line() {
	file, _ := os.Open("testfile")
	defer file.Close()

	buf_reader := bufio.NewReader(file)

	for {
		buf, is_prefix, _ := buf_reader.ReadLine()
		if is_prefix == false {
			break
		}
		fmt.Println(string(buf))
	}
}

func test_command() {
	cmd := exec.Command("ls")
	buf, _ := cmd.Output()
	fmt.Println(string(buf))
}

func main() {
	// test_basic_read()
	// test_buffer_read()
	test_command()
}