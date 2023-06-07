package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//a := "1 149,6088607594936;1 179,0666666666666"
	// Ожидаем и считываем введённую строку в rs (строка с пробела, для строки без пробелов хатило бы fmt.Scan())
	var a string
	//_, err := fmt.Scan(&a)
	a, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	b := strings.Split(a, ";")
	c := strings.Split(b[0], " ")
	d := c[0] + c[1]
	d = strings.Replace(d, ",", ".", -1)
	e := strings.Split(b[1], " ")
	f := e[0] + strings.Trim(e[1], "\\")
	f = strings.Replace(f, ",", ".", -1)
	fmt.Println(d, f)
	g, err1 := strconv.ParseFloat(d, 64)
	if err1 != nil {
		panic(err1)
	}
	h, err2 := strconv.ParseFloat(f, 64)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%.4f \n", g/h)
}
