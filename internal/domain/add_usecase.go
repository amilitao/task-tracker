package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func Add(arg string) {

	uuid, errId := uuid.NewRandom()

	if errId != nil {
		fmt.Println("Erro ao gerar id:", errId)
		return
	}

	task := Task{
		Id:          uuid.String(),
		Description: arg,
		Status:      "done",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	dadosJson, err := json.MarshalIndent(task, "", " ")

	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	caminho := "internal/dados/pedido.json"

	err = os.WriteFile(caminho, dadosJson, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
		return
	}

	fmt.Println("Arquivo salvo com sucesso em:", caminho)

	fmt.Printf("Output: Task added successfully (ID: %s)", uuid)
}
