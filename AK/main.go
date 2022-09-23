package main

func main() {
	if err := Expose(); err != nil {
		panic(err)
	}
}
