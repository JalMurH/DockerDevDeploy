package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HaldleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/home", server.AddMiddlewares(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
