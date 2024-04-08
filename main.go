package main

import (
	"SysSpecter/network"
	"flag"
	"fmt"
	"strconv"
)

var version = "1.2.1"

func banner() {
	fmt.Print("\033[94m"+`

	██████▓██   ██▓  ██████   ██████  ██▓███  ▓█████  ▄████▄  ▄▄▄█████▓▓█████  ██▀███  
	▒██    ▒ ▒██  ██▒▒██    ▒ ▒██    ▒ ▓██░  ██▒▓█   ▀ ▒██▀ ▀█  ▓  ██▒ ▓▒▓█   ▀ ▓██ ▒ ██▒
	░ ▓██▄    ▒██ ██░░ ▓██▄   ░ ▓██▄   ▓██░ ██▓▒▒███   ▒▓█    ▄ ▒ ▓██░ ▒░▒███   ▓██ ░▄█ ▒
	  ▒   ██▒ ░ ▐██▓░  ▒   ██▒  ▒   ██▒▒██▄█▓▒ ▒▒▓█  ▄ ▒▓▓▄ ▄██▒░ ▓██▓ ░ ▒▓█  ▄ ▒██▀▀█▄  
	▒██████▒▒ ░ ██▒▓░▒██████▒▒▒██████▒▒▒██▒ ░  ░░▒████▒▒ ▓███▀ ░  ▒██▒ ░ ░▒████▒░██▓ ▒██▒
	▒ ▒▓▒ ▒ ░  ██▒▒▒ ▒ ▒▓▒ ▒ ░▒ ▒▓▒ ▒ ░▒▓▒░ ░  ░░░ ▒░ ░░ ░▒ ▒  ░  ▒ ░░   ░░ ▒░ ░░ ▒▓ ░▒▓░
	░ ░▒  ░ ░▓██ ░▒░ ░ ░▒  ░ ░░ ░▒  ░ ░░▒ ░      ░ ░  ░  ░  ▒       ░     ░ ░  ░  ░▒ ░ ▒░
	░  ░  ░  ▒ ▒ ░░  ░  ░  ░  ░  ░  ░  ░░          ░   ░          ░         ░     ░░   ░ 
		  ░  ░ ░           ░        ░              ░  ░░ ░                  ░  ░   ░     
			 ░ ░                                       ░                                 																												 
			SysSpecter - System Analyzer API - v`, version, "\033[0m")

	fmt.Print("\n\n\n")
}

func run_test() {
	network.GetEstablishedConnections()
	network.ListInterfaces()
}

func main() {
	host := flag.String("host", "127.0.0.1", "Specify the host to run the server on (default: 127.0.0.1)")
	port := flag.Int("port", 8080, "Specify the port to run the server on (default: 8008)")
	interface_ := flag.String("interface", "", "Specify interface to start program")
	limit_ip := flag.String("limit_ip", "", "Restrict API access to a specified IP address to prevent public usage.")
	header_password := flag.String("password", "", "Restrict API access to a specified Header Password \"X-Pass\" to prevent public usage.")

	showHelp := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}

	banner()
	run_test()

	if *interface_ == "" {
		fmt.Println("[!] Please choose \033[93mat last one interface\033[0m: ", network.ListInterfaces())
	} else if !Contains(network.ListInterfaces(), *interface_) {
		fmt.Println("[x] \033[91mincorrect interface\033[0m please choose correct.")
		return
	} else {
		network.DefaultInterface = *interface_
		fmt.Printf("  →  \033[96mSysSpecter\033[0m API running on \033[92mhttp://%s:%d\033[0m ←\n\n", *host, *port)

		if *host != "localhost" && *host != "127.0.0.1" {
			fmt.Println("[!] \033[93mWarning!!!\033[0m: Your host is not set to localhost or 127.0.0.1. Please be aware that your API may be accessible from the public network.")
		}

		// run CalculatePerSecond with go thread
		go network.CalculatePerSecond()
		DefaultHeaderPassword = *header_password
		DefaultIpRestrict = *limit_ip
		Run(*host, *port)
		println("[x] \033[91mPermission Denied\033[0m. Run it as sudo. if you think your port is busy do \"lsof -P -i -n | grep :" + strconv.Itoa(*port) + "\" ")
	}

}
