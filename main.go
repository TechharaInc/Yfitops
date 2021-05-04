// クッキンアイドル アイ!マイ!まいん! - Wikipedia
// https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%83%E3%82%AD%E3%83%B3%E3%82%A2%E3%82%A4%E3%83%89%E3%83%AB_%E3%82%A2%E3%82%A4!%E3%83%9E%E3%82%A4!%E3%81%BE%E3%81%84%E3%82%93!

package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("AAGH .env COULD NOT LOADED!")
	}

}
