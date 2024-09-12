package main

import (
    "fmt"
    "html/template"
    "math/rand"
    "net/http"
    "time"
)

// Opciones del juego
var choices = []string{"Piedra", "Papel", "Tijera"}

// Estructura para almacenar los resultados
type GameResult struct {
    PlayerChoice string
    AIChoice     string
    Result       string
}

// Función que genera la elección de la IA
func getAIChoice() string {
    rand.Seed(time.Now().UnixNano())
    return choices[rand.Intn(len(choices))]
}

// Función para determinar el resultado del juego
func determineWinner(playerChoice, aiChoice string) string {
    if playerChoice == aiChoice {
        return "Empate"
    }
    if (playerChoice == "Piedra" && aiChoice == "Tijera") ||
        (playerChoice == "Papel" && aiChoice == "Piedra") ||
        (playerChoice == "Tijera" && aiChoice == "Papel") {
        return "¡Ganaste!"
    }
    return "Perdiste"
}

// Controlador para manejar el juego
func playHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Obtener la elección del jugador
        playerChoice := r.FormValue("choice")
        aiChoice := getAIChoice()
        result := determineWinner(playerChoice, aiChoice)

        // Crear un resultado del juego
        gameResult := GameResult{
            PlayerChoice: playerChoice,
            AIChoice:     aiChoice,
            Result:       result,
        }

        // Renderizar el resultado en la plantilla
        tmpl := template.Must(template.ParseFiles("index.html"))
        tmpl.Execute(w, gameResult)
    } else {
        http.ServeFile(w, r, "index.html")
    }
}

func main() {
    // Ruta del juego
    http.HandleFunc("/play", playHandler)

    // Servidor en puerto 8080
    fmt.Println("Servidor escuchando en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
