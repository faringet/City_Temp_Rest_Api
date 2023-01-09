# City Temp Rest Api
С помощью этого сервиса можно узнать точную температуру в любом городе нашей планеты.
___
### На данный момент:
- сервис принимает от юзера по RestAPI наименование города и записывает температуру в БД
- обработка запросов с помощю [**Gin Web Framework**](https://gin-gonic.com/docs/)
- работа с БД и миграциями с помощью [**GORM**](https://gorm.io/docs/)
- запись осуществляется в PostgreSQL
- gin также пишет логи в консоль
- логирование [**Logrus'ом**](https://github.com/sirupsen/logrus) 
- тесты [**Testify**](https://github.com/stretchr/testify)
- swagger документация
___
### Взаимодействие через запросы:
![](https://github.com/faringet/City_Temp_Rest_Api/blob/master/swagger%20screenshots/all.png)

Чтобы запросить температуру, интересующего нас города, направляем POST запрос с json'ом, тело которого имеет вид: 
```javascript 
{
    "City" : "запрашиваемый город на латинице"
}
 ``` 
 Получим ответ:
 ```javascript 
 {
    "sub": {
        "ID": 11,
        "CreatedAt": "2022-12-25T18:34:20.675281+03:00",
        "UpdatedAt": "2022-12-25T18:34:20.675281+03:00",
        "DeletedAt": null,
        "City": "запрашиваемый город на латинице",
        "Temperature": текущая температура
    }
}
 ```
 
Для знакомства со всеми запросами [**линк на скрины из swagger'а**](https://github.com/faringet/City_Temp_Rest_Api/tree/master/swagger%20screenshots)
 
