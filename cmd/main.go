package main

func main() {
	app := NewApp()

	defer func() {
		app.entClient.Close()
		app.rabbitmq.Conn.Close()
	}()

	app.Start()
}
