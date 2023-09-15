# База данних

В якості бази данних використовується Postrges.

Для того аби працювати з базою данних через код golang потрібно підключити такі пакети:

- драйвер бази данних postgres (pgx)
- міграції бази данних (migration)
- генератор запитів query builder (sqlc)

ORM використовуватися не буде (принаймі на початку)

## Драйвер бази данних

Підключаємо драйвер

```bash
go get github.com/jackc/pgx/v5
```

## Migration

Для міграцій будемо використовувати пакет migration

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
```

Копіюємо файл в робочу директорію.

```bash
cp /go/bin/migrate /var/www/app/db/
```

## Query builder (sqlc)

Для підключення sqlc:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0
```

Копіюємо файл в робочу директорію.

```bash
cp /go/bin/sqlc /var/www/app/db/
```
