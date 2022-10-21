package main

import (
	"fmt"
	"log"
	"net"
	"timefromntp/package/ntp_"
)

func main() {
	var (
		ntp    *ntp_.Ntp
		buffer []byte
		err    error
		ret    int
	)
	// Ссылка на NTP-сервер Alibaba Cloud, NTP имеет много бесплатных серверов, которые могут использовать time.windows.com
	conn, err := net.Dial("udp", "ntp1.aliyun.com:123")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		conn.Close()
	}()
	ntp = ntp_.NewNtp()
	conn.Write(ntp.GetBytes())
	buffer = make([]byte, 2048)
	ret, err = conn.Read(buffer)
	if err == nil {
		if ret > 0 {
			ntp.Parse(buffer, true)
			fmt.Println(fmt.Sprintf(
				"LI:% d\r\n Версия:% d\r\n Мод:% d\r\n Точность:% d\r\n Опрос:% d\r\n Точность системы:% d\r\n Задержка:% ds\r\n Максимальная ошибка:% d\r\n Представление часов:% d\r\n Отметка времени:% d% d% d% d\r\n ",
				ntp.Li,
				ntp.Vn,
				ntp.Mode,
				ntp.Stratum,
				ntp.Poll,
				ntp.Precision,
				ntp.RootDelay,
				ntp.RootDispersion,
				ntp.ReferenceIdentifier,
				ntp.ReferenceTimestamp,
				ntp.OriginateTimestamp,
				ntp.ReceiveTimestamp,
				ntp.TransmitTimestamp,
			))
		}
	}
}
