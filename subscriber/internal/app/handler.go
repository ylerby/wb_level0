package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (a *App) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		formId := r.FormValue("id")
		id, err := strconv.Atoi(formId)
		if err != nil {
			log.Printf("ошибка при конвертации id записи %s", err)
			_, err = w.Write([]byte("Некорректное значение id"))
			if err != nil {
				log.Printf("ошибка при записи %s", err)
			}
			return
		}

		log.Println("Попытка получения значения из кеша")
		val, ok := a.Cache.GetById(id)
		if ok {
			response, err := json.Marshal(*val)
			if err != nil {
				log.Printf("ошибка при сериализации объекта %s", err)
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}

			log.Println("Значение получено из кеша")
			_, err = w.Write(response)
			if err != nil {
				log.Printf("ошибка при ответе %s", err)
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
			return
		} else {
			log.Println("Попытка получения значения из sql")
			sqlVal, ok := a.Sql.GetById(id)
			if !ok {
				_, err = w.Write([]byte("Заказ с таким id не найден"))
				if err != nil {
					log.Printf("ошибка при ответе %s", err)
					http.Redirect(w, r, "/", http.StatusBadRequest)
				}
			} else {
				response, err := json.Marshal(*sqlVal)
				if err != nil {
					log.Printf("ошибка при сериализации объекта %s", err)
					http.Redirect(w, r, "/", http.StatusBadRequest)
				}

				_, err = w.Write(response)
				if err != nil {
					log.Printf("ошибка при ответе %s", err)
				}
			}
		}
	} else if r.Method == "GET" {
		html := `<!DOCTYPE html>
			<html>
			<head>
			   <title>Главная</title>
			</head>
			<body>
			   <h1>Получение записи по id </h1>
			   <form action="/" method="post">
				  <label for="id">Введите ID:</label><br>
				  <input type="text" id="id" name="id"><br>
				  <input type="submit" value="Отправить">
			   </form>
			</body>
			</html>`

		_, err := fmt.Fprintf(w, html)
		if err != nil {
			log.Printf("ошибка при отображении html-файла %s", err)
		}
	}
}
