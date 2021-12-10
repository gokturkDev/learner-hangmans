package main

func inferType(x string) {
	switch x {
	case "PARENT":
		println(x)
	case "CHILD":
		println(x)
	default:
		println("NOT RECOGNIZED")
	}
}

func main() {
	inferType("PARENT")
}
