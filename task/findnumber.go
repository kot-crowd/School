/* Папки для работы в папке "Загрузки"* или в репозитории https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_sep_1 */
package main

import (
	//"bytes"
	"encoding/csv"
	"fmt"

	//"io/ioutil"
	"os"
	"path/filepath"
	//"strings"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	//buf := bytes.NewBuffer(nil)
	if info.IsDir() {
		return nil
	}
	if err != nil {
		return err
	}
	buf, err := os.Open(path)
	if err != nil {
		return err
	}
	//Отложенное закрытие файла
	defer buf.Close()
	//Создаём объект читальщик данных csv формата
	r := csv.NewReader(buf)
	//Читаем всё сразу
	s, err := r.ReadAll()
	if err != nil {
		return err
	}
	//if path == "dirs/dir1/dir2/dir3/dir5/dir5/file1.txt" {
	fmt.Printf("%s\n ", path) //Выводим путь к читаемомму файлу
	//}
	// Если найден файл csv выводим элемент 5 строки 3-го столбца
	if len(s) > 1 {
		for _, elem := range s {
			fmt.Println(elem)
			//fmt.Println(s[indx])
		}
		fmt.Println(s[4][2])
		/*}
		// Читаем данные из того же файла
		fmt.Printf("%s\n ", path)
		//csv.NewReader(buf)
		buf, err1 := ioutil.ReadFile(path)
		//dataFromFile := csv.NewReader(buf)
		fmt.Printf("%s\n ", string(buf))
		if err1 != nil {
			return err1
		}
		if strings.Contains(string(buf), ",") {
			//fmt.Printf("%d ", len(info.Name())+1)
			fmt.Printf("%s ", info.Name())*/
		return nil
	} else {
		fmt.Printf("%s\n", s) //Выводим содержимое файла
		return nil
	}
}

func main() {
	//const root = "./dirs/dir1/dir2/dir3/dir5/dir5/"
	const root = "/home/kt/Загрузки/task/dir1/dir2/dir3/dir5/dir5/"
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
