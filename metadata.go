package main

type PropertyName string

type PropertyArgument any

type PropertyIdentifier struct {
	Scope string
	PropertyName
	PropertyArgument
}

type PropertyValue any

type PropertyReader interface {
	Read(PropertyIdentifier, File) PropertyValue
}
