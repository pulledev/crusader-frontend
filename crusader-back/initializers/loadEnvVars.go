package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvVars() {
	// Open the file
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the file and read each line
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		data := strings.Split(line, "=")
		err := os.Setenv(data[0], data[1])
		if err != nil {
			log.Fatalf("failed to set env var: %s", err)
			return
		}
	}
}
