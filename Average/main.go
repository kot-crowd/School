data, err := ioutil.ReadAll(os.Stdin)

if err := json.Unmarshal(data, &s); err != nil {
          fmt.Println(err)
        return
    }

type student struct {
    LastName   string
    FirstName  string
    MiddleName string
    Birthday   string
    Address    string
    Phone      string
    Rating     []int
}

type List struct {
    ID       int
    Number   string
    Year     int
    Students []student
}
type response struct {
    Average float32
}
/************************************************/
type myStruct struct {
	Students []struct {
		Rating []int64 `json:"Rating"`
	} `json:"Students"`
}

type myResult struct {
	Average float32 `json:"Average"`
}
/************************************************/
for _, studentObject := range groupObject.Students {

    countStudent ++

    for _, rating := range studentObject.Rating {
      countRating ++
    }
    }

    type studentStruct struct {
        LastName string
        FirstName string
        MiddleName string
        Birthday string
        Address string
        Phone string
        Rating []int
      }
      
      type groupStruct struct {
        ID string
        Number string
        Year int
        Students []studentStruct
      }
 /**************************************************/ 
 os.Stdout.WriteString(string(data2))
 /**************************************************/
 fmt.Scan(&str)

fmt.Println(string(unmarshaled))
/*************************************************/
   w := os.Stdout
   w.WriteString(string(data))        
/*************************************************/
data :=[]byte(`{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
    },
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[4,5,6]
    },
    ]}`)
 /*****************************************************/
 func main() {
    data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
       // ...
    }
    panic(string(data)) //с этим и работаем
 
 }
/************************************************************/
json.NewDecoder(os.Stdin).Decode(&переменнаяГдеЛежитГлавнаяСтруктура)
/***************************************************************/
/*

Сердечно благодарю коллег, кто не поленился и создал набросок json-файла.

Обратите внимание на то, что в состав общего json, назовем структуру с ним, например Group, входит элемент массив, который называется Students, 
элементами которого является отдельные структуры, которые так же надо описывать, описание этой структуры приводить не буду, 
пусть она называется Student, тогда при описании Group, в конце надо будет добавить описание Students оно будет выглядеть как-то так: */

Students []Student

//Нам необходимо создать переменную, в которой мы бы могли хранить вводимую информацию. В описании к задаче предлагается использовать:

byteValue, err := ioutil.ReadAll(источник информации)  // источником может быть файл, либо терминал (os.Stdin)

//После этого, следует создать экземпляр нашей структуры Group.

var result Group

//и "распаковать" в него json, используем метод из пакета json - Unmarshal, передать это значение в созданный ранее экземпляр.

json.Unmarshal(byteValue, &result)

/*Теперь мы можем обращаться по названию поля, к соответствующей информации нашем json.

Например: */

result.Students /*выведет нам массив с информацией по студентам. Мы можем перебрать все значения в цикле for и вытащить значения Rating

Имея эту информацию можно посчитать количество студентов, и количество оценок.*/
/*******************************************************************/
_, err = os.Stdout.Write(dataOut)
/*******************************************************************/
io.WriteString(os.Stdout, fmt.Sprintf("%s", exp))
/******************************************************************/
       