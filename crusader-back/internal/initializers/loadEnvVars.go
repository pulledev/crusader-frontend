package initializers

import (
	"bufio"
	"fmt"
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

		// If there is a comment that is marked with #
		if strings.HasPrefix(line, "#") {
			continue
		}

		//Removes every " in the line, because its just for blank spaces
		line = strings.Replace(line, "\"", "", -1)
		//Split the Line into 2 parts between the first "="
		fmt.Println("Nach der Tri: ", line)
		data := strings.SplitN(line, "=", 2)

		fmt.Println(data[1])

		if len(data) < 2 {
			log.Printf("Skip data %s. Reason: '=' count under 2\n", data)
			continue
		}

		err := os.Setenv(data[0], data[1])

		if err != nil {
			log.Fatalf("failed to set env var: %s", err)
			return
		}

	}
}
