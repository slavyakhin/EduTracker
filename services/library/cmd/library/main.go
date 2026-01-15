package main

const (
	envLocal = "local"
	envDev   = "dev"
	enbProd  = "prod"
)

func main() {
	// TODO: load config

	// TODO: setupLogger

	// TODO: create application instance
	// TODO: run application in goroutine

	// TODO: graceful shutdown (chan)
	/*
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		recievedSignal := <-stop
	*/

}
