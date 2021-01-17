package main

import "fmt"

func main() {

	fmt.Println("\033[1;41m Text: White - Background: Red\033[0m ")
	fmt.Println("\033[1;42m Text: White - Background: Green\033[0m ")
	fmt.Println("\033[1;43m Text: White - Background: Yellow\033[0m ")
	fmt.Println("\033[1;44m Text: White - Background: Blue\033[0m ")
	fmt.Println("\033[1;45m Text: White - Background: Magenta\033[0m ")
	fmt.Println("\033[1;46m Text: White - Background: Cyan\033[0m ")
	fmt.Println("\033[1;47m Text: White - Background: Grey light\033[0m ")
	fmt.Println("\033[1;30m Text: Grey \033[0m ")
	fmt.Println("\033[2;30m Text: Grey normal \033[0m ")
	fmt.Println("\033[3;30m Text: Grey italic \033[0m ")
	fmt.Println("\033[4;30m Text: Grey underline \033[0m ")
	fmt.Println("\033[1;31m Text: Red \033[0m ")
	fmt.Println("\033[2;31m Text: Red normal\033[0m ")
	fmt.Println("\033[3;31m Text: Red italic\033[0m ")
	fmt.Println("\033[1;32m Text: Green \033[0m ")
	fmt.Println("\033[2;32m Text: Green normal \033[0m ")
	fmt.Println("\033[1;33m Text: Yellow \033[0m")
	fmt.Println("\033[2;33m Text: Yellow normal\033[0m")
	fmt.Println("\033[1;34m Text: Blue \033[0m ")
	fmt.Println("\033[2;34m Text: Blue normal\033[0m ")
	fmt.Println("\033[1;35m Text: Magenta \033[0m ")
	fmt.Println("\033[2;35m Text: Magenta normal\033[0m ")
	fmt.Println("\033[1;36m Text: Cyan \033[0m ")
	fmt.Println("\033[2;36m Text: Cyan normal\033[0m ")

}
