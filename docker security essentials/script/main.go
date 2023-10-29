package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func createLowLeverUser() {

	var username string
	var group string

	fmt.Print("Enter a low-prev username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		log.Fatal("Error", err)
		os.Exit(1)
	}

	fmt.Print("Enter a group name: ")
	_, err = fmt.Scanln(&group)
	if err != nil {
		log.Fatal("Error", err)
		os.Exit(1)
	}

	fmt.Println("%s %s", username, group)

	dockerfileContent := fmt.Sprintf(`
FROM ubuntu:latest

LABEL maintainer = "Sam Sepiol"

RUN groupadd -r %s && useradd -r -g %s %s
RUN chsh -s /usr/sbin/nologin root
`, group, group, username)

	errr := os.WriteFile("output/Dockerfile", []byte(dockerfileContent), 0644)
	if errr != nil {
		fmt.Println("Error writing Dockerfile: ", err)
		os.Exit(1)
	}
	fmt.Println("Created a docker file with low level user, check the file in output folder!")
}

// TODO: Build the file
func buildContainer() {
	cmd := exec.Command("docker", "build", "output/", "-t", "test")
	fmt.Println("Building docker image... ")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error building the docker file %v \n", err)
		return
	}
	fmt.Println("Image built successfully")
	startIndex := strings.LastIndex(string(output), "Successfully built")

	if startIndex == -1 {
		fmt.Println("Output does not contain 'Successfully built'.")
		return
	}

	// Extract the part of the string following "Successfully built"
	extractedValue := strings.TrimSpace(string(output[startIndex+len("Successfully built"):]))

	fmt.Println("Extracted Value:", extractedValue)
}

func main() {
	fmt.Println(`
Docker-Security
1. Create a dockerfile with low lever user
2. Run the DockerFile
3. Run in readonly mode
4. Disable inter-container communication
5. Build the docker file
`)

	var option string

	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Fatal(err)
	}

	if option == "1" {
		createLowLeverUser()
	}
	if option == "5" {
		buildContainer()
	}
}
