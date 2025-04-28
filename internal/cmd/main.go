package main

import (
	"firestore-test/internal/infra/config/instance"
	"firestore-test/internal/infra/config/property"
	"gitlab.falabella.tech/rtl/logistics-corp/rso-portfolio/libraries/golang/lightms"
)

func main() {
	registerProperties()
	registerPrimaries()
	lightms.Run()
}

func registerProperties() {
	lightms.SetPropFilePath("internal/resources/properties.yml")
	lightms.AddProperty(property.GetServerProperty())
	lightms.AddProperty(property.GetApplicationProperty())
	lightms.AddProperty(property.GetFirestoreProperty())
}

func registerPrimaries() {
	lightms.AddPrimary(instance.GetControllerInstance)
}
