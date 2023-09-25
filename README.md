# Бекенд

## Підготовка до локального запуску

Створіть файл в диреторії **dev_local** з назвою **.env** і по прикладу з файлу .**dev_local/env.example** заповніть його інформаціює з вашої локальної машини.

Для того аби побачити id вашого локального користувача в консолі запустіть команду:

```bash
id -u
```

Створіть файл **dev_local/.env** за приклад потрібно взяти файл **dev_local/.env.example** з таким вмістом:

```bash
USER_ID="id локального користувача, скоріше за все це буде 1000"
DATABASE_URL=postgresql://postgres:postgres@mp_postgres:5432/store
```

## Локальний запуск

Для того аби запустити проект локально потрібно спочатку встановити собі docker та docker-compose. Актуальні інструкції по встановленню можна знайти по наведеним посиланням, для кожної операційної системи як то linux, windows, mac.

- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

### Запуск docker-compose

Після встановлення вище наведених залежностей, потрібно викачати собі на локальну машину репозиторій з кодом і перейти в гілку dev.

Cтворити docker мережу за допомогою команди:

```bash
docker network create mp_network
```

Перевірити чи була створена docker мережа за допомогою команди:

```bash
docker network ls
```

Cтворити docker volume за допомогою команди:

```bash
docker volume create --name=mp_postgres_pg_data
```

Перевірити чи буd створений docker volume за допомогою команди:

```bash
docker volume ls
```

Використовуючи консоль зайти в директорію dev_local яка містить файл docker-compose.yml і скориставшись наведеною комадною запустити контейнери для локальної розробки фронтенд застосунку.

```bash
docker-compose up -d
```

Дочекатися поки будуть завантажені усі необхідні образи з інтернету і перевірити стан роботи за допомогою наведеної команди в середині директорії dev_local

```bash
docker-compose ps
```

Якщо контейнери запущені тоді потрібно зайти в середину виконавши таку команду:

```bash
docker exec -it mp_backend bash
```

Перебуваючи в середині контейнера запустіть застосунок.

```bash
go run main.go
```

В випадку успішного запуску ви зможете мати доступ до застосунку через ваш браузре за адресою http://localhost:10000

## Встановлення інструментів для бази данних


### Міграції для бази данних
Перебуваючи в контейрені встановлюємо **migrate**

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
```

Після встановлення копіюємо виконуваний файл в директорію **/bin**

```bash
cp /go/bin/migrate /var/www/app/bin/
```

### Генератор коду для запитів в бд sqlc

Перебуваючи в контейнері встановлюємо **sqlc**

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0
```

Після встановлення копіюємо виконуваний файл в директорію **/bin**

```bash
cp /go/bin/sqlc /var/www/app/bin/
```

В директорії **/bin** потрібно створити два нових файли взявши за приклад файли зі словом **example** в назві. Поля в файлах заповнити актуальною інформацією для вашого локального оточення.

- **./bin/sqlc.yaml** як приклад взяти dev_local/sqlc/sqlc.yaml.example
- **./bin/schema.sql** як приклад взяти dev_local/sqlc/schema.sql.example

## Структура директорій

### dev_local

Тут зберігаються необхідні файли для локального запуску проєкту.

### dev_prod

Тут зберігаються файли для локального тестування контейнеру який буде запускатися на проді.

## Аутентифікація через гугл

Для початку потрібно створити три URL на нашому сайті

- /auth
- /auth-get-googe
- /auth-set-googe

Налаштовуємо свій гугл акаунт для отримання налаштувань від гугла. Отримуємо файл з інформацією для підключення у вигляді JSON файлу.

Формуємо сторінку за адресою /auth.
Сторінка буде містити посилання з написом "увійти через Google" і буде вести на нашу іншу сторінку **/auth-get-googe**

наш URL **/auth-get-googe** прочитає файл JSON який ми отримали від гугла раніше і сформувавши необхідний запит, зробить редірект на вказаний в файлі URL в google

Гугл отримавши від нас інформацію, зробить запит на наш URL який був переданий йому раніше, а саме **/auth-set-googe** і передасть нам токен

Отримавши токен ми робимо ще один запит в гугл з цим токеном і отримаємо інформацію про користувача

Запишемо інформацію про користувача в базу, створимо новий JWT токен і робимо редірект з цим токеном

Користувач отримуж токен

### Пакет для oauth2

https://developers.google.com/identity/protocols/oauth2

Запити в гугл можна робити і без цього пакету але ми будемо його використовувати тому що по ньому більше інформації.

Встановлюємо пакет oauth2

```bash
go go get google.golang.org/api/oauth2/v2
go mod vendor
```

## Gin

Категорії методів:
- прочитати інформацію від клієнта
- провалідувати інформацію від клієнта
- повернути інформацію клієнту



	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		//user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			//db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})