package main

import "fmt"

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  infinity-dog help")
	fmt.Println("  infinity-dog sample     [hours]")
	fmt.Println("  infinity-dog services   [hits|exceptions|data] [info|warn|debug]")
	fmt.Println("  infinity-dog exceptions [service]")
	fmt.Println("  infinity-dog messages   [service]")
	fmt.Println("  infinity-dog import")
	fmt.Println("  infinity-dog screen")
	fmt.Println("")
}
