package main

import (
	"fmt"
	"log"

	//"github.com/go-telegram-bot-api/telegram-bot-api"
	//те что для сервера
	//"fmt"           // вывод в консоль
	// библиотека использвания html-шаблонов
	//"log"           // библиотека по логированию ошибок приложения, журналирования log.Fatal("invalid whole number")
	"net/http" // библиотека по доступу к сетям
)

func main() {
	log.Println("Zapusk")
	http.HandleFunc("/", homePage) // при адресе / вызов функции HomePage
	log.Println("Zapusk2")
	http.HandleFunc("/about", about) // при адресе / вызов функции HomePage
	log.Println("Zapusk3")
	if err := http.ListenAndServe(":80", nil); err != nil {
		// запускаем сервер, который слушает 80 порт. Если сервер не запустился, например занят порт другим предложением, то  выводится ошибка в лог
		log.Fatal("server ne start. Neobxodim reboot. nil определяет тип сервера", err)
		// пример того что выводит при включенном  перед запуском сервера IIS на 80 порту :
		// 2018/09/09 21:19:48 server ne start. Neobxodim reboot. nil определяет тип сервераlisten tcp :80: bind: An attempt was made to access a socket in a way forbidden by its access permissions.
	}
	// сервер слушает входной 80 порт, если он закрыт то ошибка в лог

}

// Aghfh
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Index Page")
}

/** Public is a public function.
 *  usage: ...
 */
func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "about.html")

}
