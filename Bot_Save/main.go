package main

import (
		"log"
		"flag"
	)

func main(){
	t := mustToken()

}
func mustToken()(string){
	token := flag.String(
		name: "token-bot-token", 
		value: "", 
		usage: "token for access to telegram bot",
)

flag.Parse()

	if *&token == ""{
	log.Fatal(v...:"token is not specified ")
}	
	
return  token

}