package main

import (
	"log"

	"firestore-test/internal/cmd/inject" // Import the inject package
)

func main() {
	// Call the injector function (which will be in wire_gen.go after generation)
	app, err := inject.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the application
	// The Start method would come from your GinController or equivalent
	app.Start()
}
