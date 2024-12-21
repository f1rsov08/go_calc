package application

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/f1rsov08/go_calc/pkg/calculation"
)

type Config struct {
	Addr string // Порт, на котором будет запущен сервер
}

// Функция для создания конфигурации из переменных окружения
func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

// Структура приложения, содержащая конфигурацию
type Application struct {
	config *Config
}

// Функция для создания нового экземпляра приложения
func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

// Метод для запуска HTTP-сервера
func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}

// Структура запроса для обработки входящих данных
type Request struct {
	Expression string `json:"expression"`
}

// Структура ответа для отправки клиенту
type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// Обработчик HTTP-запросов для выполнения вычислений
func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if err.Error() == "Expression is not valid" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(Response{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: result})
}
