package main

import (
	"bufio"
	"fmt"
	"message"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("/home/ec2-user/Training/Go/ACloudGuru-SystemToolingWithGo/Go/function_basics/fileout.out", os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error while opening the file /home/ec2-user/Training/Go/ACloudGuru-SystemToolingWithGo/Go/function_basics/fileout.out")
		os.Exit(1)
	}

	defer f.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(phrase, name)
	fmt.Println(message, "!!")

	// //Up until here we were printing the message to the Stdout
	// //Now we would like to write it to the a file

	_, err = f.Write([]byte(message))

	// _ := os.WriteFile("/home/ec2-user/Training/Go/ACloudGuru-SystemToolingWithGo/Go/function_basics/fileout.out", []byte(message), 0644)

	if err != nil {
		fmt.Println("Unable to write to specified file!")
		os.Exit(1)
	}

}
