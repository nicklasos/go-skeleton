package main

func main() {
	exitIfError(configLoad())
	exitIfError(loggerCreate())
	exitIfError(dbConnect())
	exitIfError(routersServe())
}
