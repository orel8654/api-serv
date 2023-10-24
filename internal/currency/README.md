## Абстракция:
 - Delivery - слой доставки. Способ обратиться к бизнес логике (или способ обращения к сервису).
 - Service - слой бизнес-логики. Способ получения информация или ее обратки в соответствии с заданным алгоритмом.
 - Repo - слой репозитория. Способ взаимодействия с данными конкретного хранилища. 
 - Contract - способ взаимодействия слоев.

## Аксиома:
 - Общение между слоями происходит по правилу сверху-вниз:
 - ```Delivery->Service->Repo```.
 - Каждый слой не знает о последующем:
   -  ```Delivery->Contranct->Service->Contranct->Repo```.
 - Обращение к какой-то логике через Delivery происходит только через Service.
   - Ложно: ```Delivery->Contranct->Repo```

## Пример:
Сервис доставки:
Delivery - HTTP.
Service - BL (Бизнес Логика).
Repo - Postgres.