# Data Enricher & Dispatcher

Простий сервіс на Go, який отримує користувачів з API, фільтрує тих, чий email закінчується на `.biz`, та надсилає їх у інший API з підтримкою повторних спроб (retry).

## Функціонал 📊
- Отримання користувачів з API A (`GET`)
- Фільтрація email-адрес на `.biz`
- Надсилання користувачів у API B (`POST`)
- Retry-механізм з 3 спробами і затримкою
- Логування
- Підтримка context
- Покриття тестами основної логіки


## Як запустити 🚀

1. Клонувати репозиторій:
```bash
git clone https://github.com/Git5737/-Data-Enricher-Dispatcher.git
cd -Data-Enricher-Dispatcher/
```

2. Створити файл .env:
```bash
API_A_URL=https://jsonplaceholder.typicode.com/users
API_B_URL=https://webhook.site/your-custom-id
```

3. Завантаження пакетів:
```bash
go mod tidy
```

4. Запустити програму:
```bash
go run ./cmd/data_enricher_dispatcher
```


