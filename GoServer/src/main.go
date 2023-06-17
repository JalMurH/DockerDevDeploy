package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HaldleRoot)
	server.Handle("POST", "/home", server.AddMiddlewares(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
