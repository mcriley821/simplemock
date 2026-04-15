package main

//go:generate simplemock -iface Notifier -out os.Stdout

type Notifier interface {
	Notify(msg string)
	NotifyAll(msgs ...string)
	NotifyWithCode(code int, msg string)
}

func main() {
	//
}
