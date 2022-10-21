package ntp_

import (
	"bytes"
	"encoding/binary"
)

const (
	UNIX_STA_TIMESTAMP = 2208988800
)

/*
*
Протокол NTP http://www.ntp.org/documentation.html
@author mengdj@outlook.com
*/
type Ntp struct {
	//1:32bits
	Li        uint8 //2 bits
	Vn        uint8 //3 bits
	Mode      uint8 //3 bits
	Stratum   uint8
	Poll      uint8
	Precision uint8
	//2:
	RootDelay           int32
	RootDispersion      int32
	ReferenceIdentifier int32
	// 64-битная временная метка
	ReferenceTimestamp uint64 // Укажите время последней калибровки системных часов
	OriginateTimestamp uint64 // Укажите время, когда клиент инициировал запрос к серверу
	ReceiveTimestamp   uint64 // относится ко времени, когда сервер получает запрос клиента
	TransmitTimestamp  uint64 // Укажите время, когда сервер отправляет метку времени клиенту
}

func NewNtp() (p *Ntp) {
	// Другие параметры обычно возвращаются сервером
	p = &Ntp{Li: 0, Vn: 3, Mode: 3, Stratum: 0}
	return p
}

/*
*

	Построение информации протокола NTP
*/
func (this *Ntp) GetBytes() []byte {
	// Обратите внимание, что в сети используется порядок байтов с прямым порядком байтов
	buf := &bytes.Buffer{}
	head := (this.Li << 6) | (this.Vn << 3) | ((this.Mode << 5) >> 5)
	binary.Write(buf, binary.BigEndian, uint8(head))
	binary.Write(buf, binary.BigEndian, this.Stratum)
	binary.Write(buf, binary.BigEndian, this.Poll)
	binary.Write(buf, binary.BigEndian, this.Precision)
	// Записать другие байтовые данные
	binary.Write(buf, binary.BigEndian, this.RootDelay)
	binary.Write(buf, binary.BigEndian, this.RootDispersion)
	binary.Write(buf, binary.BigEndian, this.ReferenceIdentifier)
	binary.Write(buf, binary.BigEndian, this.ReferenceTimestamp)
	binary.Write(buf, binary.BigEndian, this.OriginateTimestamp)
	binary.Write(buf, binary.BigEndian, this.ReceiveTimestamp)
	binary.Write(buf, binary.BigEndian, this.TransmitTimestamp)
	//[27 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	return buf.Bytes()
}

func (this *Ntp) Parse(bf []byte, useUnixSec bool) {
	var (
		bit8  uint8
		bit32 int32
		bit64 uint64
		rb    *bytes.Reader
	)
	// Похоже, что этот двоичный файл. Чтение может быть прочитано только по порядку, и не может быть прочитано за раз. Если вы хотите читать за один раз, вы можете использовать только слайс bf.
	rb = bytes.NewReader(bf)
	binary.Read(rb, binary.BigEndian, &bit8)
	// Смещение 6 бит вправо для получения первых двух LI
	this.Li = bit8 >> 6
	// смещение 2 цифр вправо и 5 цифр вправо, чтобы получить первые 3 цифры
	this.Vn = (bit8 << 2) >> 5
	// Смещение 5 бит влево, затем смещение 5 бит вправо, чтобы получить последние 3 бита
	this.Mode = (bit8 << 5) >> 5
	binary.Read(rb, binary.BigEndian, &bit8)
	this.Stratum = bit8
	binary.Read(rb, binary.BigEndian, &bit8)
	this.Poll = bit8
	binary.Read(rb, binary.BigEndian, &bit8)
	this.Precision = bit8

	//32bits
	binary.Read(rb, binary.BigEndian, &bit32)
	this.RootDelay = bit32
	binary.Read(rb, binary.BigEndian, &bit32)
	this.RootDispersion = bit32
	binary.Read(rb, binary.BigEndian, &bit32)
	this.ReferenceIdentifier = bit32

	// Следующие поля представляют собой 64-битные метки времени (NTP - это 64-битные метки времени)
	binary.Read(rb, binary.BigEndian, &bit64)
	this.ReferenceTimestamp = bit64
	binary.Read(rb, binary.BigEndian, &bit64)
	this.OriginateTimestamp = bit64
	binary.Read(rb, binary.BigEndian, &bit64)
	this.ReceiveTimestamp = bit64
	binary.Read(rb, binary.BigEndian, &bit64)
	this.TransmitTimestamp = bit64
	// Преобразование в метку времени Unix, сначала смещение 32 бита, чтобы получить целую часть 64-битной метки времени, а затем запуск метки времени ntp 1 января 1900 года 0h0m0s
	if useUnixSec {
		this.ReferenceTimestamp = (this.ReceiveTimestamp >> 32) - UNIX_STA_TIMESTAMP
		if this.OriginateTimestamp > 0 {
			this.OriginateTimestamp = (this.OriginateTimestamp >> 32) - UNIX_STA_TIMESTAMP
		}
		this.ReceiveTimestamp = (this.ReceiveTimestamp >> 32) - UNIX_STA_TIMESTAMP
		this.TransmitTimestamp = (this.TransmitTimestamp >> 32) - UNIX_STA_TIMESTAMP
	}
}
