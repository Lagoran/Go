package main

import (
	"bufio"
	"flag"
	"fmt"
	"message"
	"os"
	"strings"
)

func main() {
	//Define flags
	var name string
	var greeting string
	var prompt bool
	var preview bool

	//Parse flags
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "phrase to use for the greeting")
	flag.BoolVar(&prompt, "prompt", false, "use prompt for imput name and greeting")
	flag.BoolVar(&preview, "preview", false, "use preview to output message without writing for output file")

	flag.Parse()

	//Show usage if flags are invalid
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(11)
	}

	//Optionally print flag and exit based on env variables
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Your Name", name)
		fmt.Println("Your Greeting", greeting)
		fmt.Println("Prompt", prompt)
		fmt.Println("Preview", preview)

		os.Exit(0)
	}

	//Conditionally read from stdin
	if prompt {
		greeting, name = renderPrompt()
	}

	//Generate the message
	message := message.Greeting(greeting, name)

	//Either preview the message or write it to file
	if preview {
		fmt.Println(message)
	} else {

	}
	//write content
	f, err := os.OpenFile("/home/ec2-user/Training/Go/ACloudGuru-SystemToolingWithGo/Go/function_basics_package/fileout.out", os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error while opening the file /home/ec2-user/Training/Go/ACloudGuru-SystemToolingWithGo/Go/function_basics_package/fileout.out")
		os.Exit(1)
	}

	defer f.Close()

	//Up until here we were printing the message to the Stdout
	//Now we would like to write it to the a file

	_, err = f.Write([]byte(message))

	// _ := os.WriteFile("/home/ec2-user/Training/Go/ACloudGuru-SystemToolingWithGo/Go/function_basics_package/fileout.out", []byte(message), 0644)

	if err != nil {
		fmt.Println("Unable to write to specified file!")
		os.Exit(1)
	}
}

func renderPrompt() (greeting, name string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}
