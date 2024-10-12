# Тестовое задание в Открытый контакт

## Деплой
## Шаг 1: клонирование репозитория
Сначала клонируйте репозиторий на свой локальный компьютер или сервер:

``` bash
https://github.com/0Abracadaber0/exchange-service.git
cd exchange-service
```

## Шаг 2: создание ```.env``` файла
В корневом каталоге вашего проекта создайте файл ```.env```, чтобы сохранить переменные окружения. Мой ```.env``` для тестирования выглядил так (файл без значений ```.env.example``` лежит в корне репозитория):
```
MYSQL_ROOT_PASSWORD=root_pass
DB_USER=root
DB_PASS=root_pass
DB_PORT=3306
DB_HOST=mysql
DB_NAME=exchange
APP_HOST=0.0.0.0
APP_PORT=8080
```
Эти значение устанвливаются по умолчанию, в случае их отсутствия в вашем окружении (кроме пароля).

## Шаг 3: настройка ```docker-compose.yaml```
Ваш файл должен выглядить следующим образом (заполненный пример находится в корне проекта):
```yaml
version: '3'
services:
  mysql:
    image: mysql:latest
    env_file:
      - .env
    environment:
      MYSQL_DATABASE: exchange
    ports:
      - "[выбранный вами порт]:[выбранный вами порт]"
    volumes:
      - mysql_data:/var/lib/mysql
    healthcheck:
      test: ["CMD-SHELL", "mysql -u root -p${MYSQL_ROOT_PASSWORD} -e 'SELECT 1'"]

      interval: 30s
      timeout: 10s
      retries: 5
      start_period: 30s



  exchange:
    build:
      context: .
      dockerfile: Dockerfile
    env_file:
      - .env
    ports:
      - "[выбранный вами порт]:[выбранный вами порт]"
    depends_on:
      mysql:
        condition: service_healthy

volumes:
  mysql_data:
```
healthcheck необходим для проверки готовности контейнера mysql к подключению приложния.

## Шаг 4: сборка и запуск приложения
Чтобы собрать и запустить приложение, выполните следующую команду в терминале:
```
docker-compose up -d --build
```
Уберите флаг ```-d```, если хотите видеть отоброжение логов.
