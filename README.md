# Wildberries task level 0

### Для запуска необходимо:
1) Находясь в директории проекта, перейти в /subscriber
   ```
   cd subscriber
   ```
2) Выполнить команду:
   ```
   docker compose up 
   ```
3) Находясь в директории проекта, перейти в /publisher/cmd
   ```
   cd publisher/cmd
   ```
4) Выполнить команду
   ```
   go run main.go
   ```
5) Находясь в директории проекта, перейти в /subscriber/cmd
   ```
   cd subscriber/cmd
   ```
6) Выполнить команду
   ```
   go run main.go
   ```

#### Проверка производится локально на localhost:8080

#### База данных заполняется "моковыми" данными, с использованием пакета gofakeit
