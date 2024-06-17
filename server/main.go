package main

func main() {

	server := NewServer(
		ServerConfig{
			Port:      8888,
			RedisAddr: "localhost:6379",
		},
	)

	server.Run()

}
