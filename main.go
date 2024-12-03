/*
	Билет №18. Форматирование строки "ASDFG" -> "asdfg"

Необходимо написать веб-севрер на GO, декодирующий поданную строку. Сервер должен запускаться по адресу `127.0.0.1:8081`.

У севрера должна быть ручка (handler) `GET /encode`. Эта ручка ожидает, что через query-параметр `?src_string=<передаваемая_строка>` будет передана строка вида `ASDFG`.

При обработке http-запроса должно происходить преобразование вида `"ASDFG" —> "asdfg"

В качестве ответа сервер должен возвращать JSON с единственным полем `result`.

Примерм запроса (curl):
```
curl --request GET http://127.0.0.1:8081/encode?src_string=ASDFG
```

Пример ответа:

	```

{"result":"asdfg"}

	```

Автор: Пелевина Татьяна Владимировна
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Input struct {
	FirstString *string `json:"first_string"`
}

type Output struct {
	Result string `json:"result"`
}

// Обработчик HTTP-запроса
func EncodeHandler(w http.ResponseWriter, r *http.Request) {

	srcString := r.URL.Query().Get("src_string")
	if srcString == "" {
		http.Error(w, "Отсутвуют нужные параметры", http.StatusBadRequest)
		return
	}

	var input Input

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	var output Output

	output.Result = strings.ToLower(srcString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	respBytes, _ := json.Marshal(output)
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/encode", EncodeHandler)

	// Запускаем веб-сервер на порту 8081
	fmt.Println("starting server...")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
