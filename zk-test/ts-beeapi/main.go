package main

func main() {
	bee := NewTsBeeReqs()
	bee.SetBaseURL("http://localhost:8000")
	bee.RunAuth()
}
