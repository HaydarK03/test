package main

import "syscall"

func poweroff() {
	syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
}

// kjfakljflk
func main() {
	poweroff()
}
