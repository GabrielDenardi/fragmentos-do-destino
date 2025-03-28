package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var (
	clickCount int
	mu         sync.Mutex
)

var normalFragments = []string{
	"Amor Fati: amar até o inevitável.",
	"Você deseja que sua vida se repita eternamente?",
	"Se tudo voltasse exatamente como foi, você mudaria algo?",
	"Você resiste ou dança com o caos?",
	"Aceitar o destino não é desistir, é abraçar a potência do que é.",
	"O universo não tem porquê. E mesmo assim... existe.",
	"Você vive como quem escolheu cada instante?",
	"Tudo o que acontece, acontece como deveria — e nada mais.",
	"Você é autor da sua história ou apenas uma testemunha?",
	"A eternidade se esconde no agora.",
}

type ClickRequest struct {
	ButtonType string `json:"buttonType"`
}

type ClickResponse struct {
	Mode      string   `json:"mode"`
	Fragments []string `json:"fragments"`
}

func generateRhetoricalQuestionWithAI() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY não definida")
	}

	url := "https://api.openai.com/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Você é um gerador de perguntas retóricas filosóficas. Crie uma pergunta retórica profunda e original.",
			},
			{
				"role":    "user",
				"content": "Gere uma pergunta retórica filosófica.",
			},
		},
		"max_tokens":  50,
		"temperature": 0.9,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", errors.New("nenhuma resposta retornada pela API")
	}
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", errors.New("formato da resposta inválido")
	}
	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return "", errors.New("formato da mensagem inválido")
	}
	content, ok := message["content"].(string)
	if !ok {
		return "", errors.New("conteúdo não encontrado na resposta")
	}

	return content, nil
}

func clickHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	var req ClickRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	clickCount++
	currentCount := clickCount
	mu.Unlock()

	var res ClickResponse

	if currentCount >= 10 {
		var generated []string
		for i := 0; i < 3; i++ {
			q, err := generateRhetoricalQuestionWithAI()
			if err != nil {
				log.Println("Erro ao gerar pergunta via IA:", err)
				q = "No âmago do desconhecido, qual é o sentido da existência?"
			}
			generated = append(generated, q)
			time.Sleep(500 * time.Millisecond)
		}
		res = ClickResponse{
			Mode:      "rhetorical",
			Fragments: generated,
		}
	} else {
		res = ClickResponse{
			Mode:      "normal",
			Fragments: normalFragments,
		}
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env:", err)
	}

	rand.Seed(time.Now().UnixNano())

	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	http.HandleFunc("/click", clickHandler)
	log.Println("Servidor rodando na porta 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
