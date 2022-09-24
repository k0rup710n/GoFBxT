package prints

import "github.com/fatih/color"

func PrintBanner() {
	color.New(color.FgGreen, color.Bold).Println("  ________      _________________________  ______________")
	color.New(color.FgGreen, color.Bold).Println(" /  _____/  ____\\_   _____/\\______   \\   \\/  /\\__    ___/")
	color.New(color.FgGreen, color.Bold).Println("/   \\  ___ /  _ \\|    __)   |    |  _/\\     /   |    |   ")
	color.New(color.FgGreen, color.Bold).Println("\\    \\_\\  (  <_> )     \\    |    |   \\/     \\   |    |   ")
	color.New(color.FgGreen, color.Bold).Println(" \\______  /\\____/\\___  /    |______  /___/\\  \\  |____|   ")
	color.New(color.FgGreen, color.Bold).Println("        \\/           \\/            \\/      \\_/           ")
	color.New(color.FgGreen, color.Bold).Println("\t\t~ Go Facebook Extractor v 1.0")
	color.New(color.FgGreen, color.Bold).Print("\t\t\tBy")
	color.New(color.FgYellow, color.Bold).Print(" Gremlin\n\n")
}

func PunctPrinter(sym) {
	color.New(color.FgWhite, color.Bold).Printf("[")
	color.New(color.FgGreen, color.Bold).Printf(sym)
	color.New(color.FgWhite, color.Bold).Printf("]")
}

func PrintPlus() {
	color.New(color.FgWhite, color.Bold).Printf("[")
	color.New(color.FgGreen, color.Bold).Printf("+")
	color.New(color.FgWhite, color.Bold).Printf("] ")
}

func PrintExclamation() {
	color.New(color.FgWhite, color.Bold).Printf("[")
	color.New(color.FgRed, color.Bold).Printf("!")
	color.New(color.FgWhite, color.Bold).Printf("] ")
}

func PrintInterrogation() {
	color.New(color.FgWhite, color.Bold).Printf("[")
	color.New(color.FgYellow, color.Bold).Printf("?")
	color.New(color.FgWhite, color.Bold).Printf("] ")
}

func PrintToken() {
	color.New(color.FgWhite, color.Bold).Printf("[")
	color.New(color.FgYellow, color.Bold).Printf("TOKEN")
	color.New(color.FgWhite, color.Bold).Printf("] ")
}
