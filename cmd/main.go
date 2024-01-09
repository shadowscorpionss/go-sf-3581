package main

import (
	"go-sf-35-8-1/pkg/proverbs"
	"log"
	"math/rand"
	"net"
	"time"
)

// Сетевой адрес.
//
// Служба будет слушать запросы на всех IP-адресах
// компьютера на порту 12345.
// Например, 127.0.0.1:12345
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

func main() {
	// Запуск сетевой службы по протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Подключения обрабатываются в бесконечном цикле.
	// Иначе после обслуживания первого подключения сервер
	//завершит работу.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn) {
	// Закрытие соединения.
	defer conn.Close()

	for {
		r := rand.Intn(len(proverbs.Proverbs))
		pr := proverbs.Proverbs[r]
		_, err := conn.Write([]byte(pr + "\r\n"))
		if err != nil {
			log.Println("error", err)
			break
		}
		time.Sleep(3 * time.Second)
	}
}
