# API scheme:

* `/api`
  * `/polls`
    * **GET** `/` - *get list of all polls*
    * **GET** `/:poll_id` - *get poll with id==poll_id*
    * **POST** `/new` - *add new poll*
    * **GET** `/results/:poll_id` - *get results for poll with id==poll_id*
    * **POST** `/:poll_id` - *sends results of poll with id==poll_id*
  * `/user`
    * **GET** `/` - *gets all users*
    * **GET** `/auth` - *check authentication*
    * **POST** `/login` - *log in*
    * **POST** `/registartion` - *sign in*
    
# Repository scheme:

- `/cmd/main`
  - `/main.go` - *entry point* 
- `/internal`
  - `/config`
  - `/user`
    - `/handler.go` - обработчик запросов
    - `/model.go` - модель данных пользователя
    - `/service.go` - "сваязывает handler и storage". Определяет интерфейс взаимодействия обработчика с хранилищем
    - `/storage.go` - 
    - `/db/postgre.go` - 
  - `/poll`
  - `/question`
  - `/...`