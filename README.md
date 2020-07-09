# norsi-trans-project

go run main.go reader.go writeTest.go queue.go testDataGenereate.go data.go proto.go [номер теста - от 1 до 5] [генерация данных или запуск теста - true или false соответственно]


| Номер теста   |  Вид таблицы       | Круто |
| ------------- |:------------------:| -----:|
| 1             | dt timestamp,target text,contact text,direction boolean,subject text,size int,PRIMARY KEY (target, dt, direction)   | $1600 |
| 2             | dt timestamp,field0 text,field1 text,field2 text,field3 text,field4 text,field5 text,field6 text,field7 text,field8 text,field9 text,size0 int,size1 int,size2 int,size3 int,size4 int,size5 int,size6 int,size7 int,size8 int,size9 int,PRIMARY KEY (dt,field0 ,field1 ,field2 ,field3 ,field4 ,field5 ,field6 ,field7 ,field8 ,field9 ,size0 ,size1 ,size2 ,size3 ,size4 ,size5 ,size6 ,size7 ,size8 ,size9)|   $12 |
| 3             | dt timestamp,data blob ,PRIMARY KEY (dt)        |    $1 |
| 4             | dt timestamp,data blob ,PRIMARY KEY (dt)         |    $1 
| 5             | dt timestamp,target text,contact text,direction boolean,subject text,size int,PRIMARY KEY (target, dt, direction,subject,size,contact)         |  $1 