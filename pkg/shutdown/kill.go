package shutdown

import "syscall"

// killTheApp sends SIGTERM to the parent application to quit
func KillTheApp() {
	pid := syscall.Getpid()
	syscall.Kill(pid, syscall.SIGTERM)
}
