package main

const a = "hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return a + name
}
