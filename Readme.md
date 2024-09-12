# Rock, Paper, Scissors Web Game in Go
## Overview
This project is a simple **Rock, Paper, Scissors** web game built using the Go programming language. The player can choose one of the three options (Rock, Paper, or Scissors) to play against an AI that randomly selects its own option. The game then determines the winner based on the standard rules of **Rock, Paper, Scissors** and displays the result on the web page.
## Features
- **Web-based interface**: The player interacts with the game through a web browser.
- **Random AI choice**: The AI selects one of the three choices randomly for each game.
- **Game logic**: The server processes the game, compares the player's choice with the AI's, and determines the outcome (win, lose, or draw).
- **Result display**: The result of each round is displayed on the web page.
## Prerequisites
- Go (version 1.16+ recommended)
- Web browser
## Project Structure
```
piedra-papel-tijera/
├── main.go          # Main application file
└── index.html       # HTML template for the game interface
```
## How to Run
1. **Clone the repository** (if you have uploaded it to GitHub):   
```bash
git clone https://github.com/yourusername/piedra-papel-tijera-go.git
cd piedra-papel-tijera-go
```
2. **Install Go**: Follow the installation guide for your operating system from [the official Go website](https://golang.org/dl/).
3. **Run the application**:   Inside the project directory, run the following command to start the web server:  
```bash
go run main.go
```
5. **Access the game**:   Open your web browser and navigate to `http://localhost:8080/play` to play the game.
## Code Overview
### `main.go`
This is the main file that sets up the HTTP server, defines the game logic, and handles the player's requests. Key components include:
- **GameResult struct**: Holds the player's choice, AI's choice, and the game result.
- **getAIChoice()**: Function that randomly selects the AI's move (Rock, Paper, or Scissors).
- **determineWinner()**: Function that compares the player's choice with the AI's and returns the result ("Win", "Lose", or "Draw").
- **playHandler()**: The HTTP handler for the `/play` route that processes the game and renders the result in the HTML template.
### `index.html`
This file contains the HTML structure and template for the game interface. It includes a form where the player selects their choice (Rock, Paper, or Scissors), and upon form submission, the game result is displayed.
### Example HTML Structure:
```html
<form action="/play" method="POST">
    <select name="choice">
        <option value="Rock">Rock</option>
        <option value="Paper">Paper</option>
        <option value="Scissors">Scissors</option>
    </select>
    <button type="submit">Play</button>
</form>

{{if .PlayerChoice}}
<p>You chose: {{.PlayerChoice}}</p>
<p>AI chose: {{.AIChoice}}</p>
<p>Result: {{.Result}}</p>
{{end}}
```
## Game Logic
The game follows the traditional **Rock, Paper, Scissors** rules:
- **Rock** beats **Scissors**.
- **Paper** beats **Rock**.
- **Scissors** beats **Paper**.
- If both the player and the AI choose the same option, the result is a **Draw**.
## Future Improvements
- Add rounds and a scoring system.
- Improve the user interface using CSS.
- Create a multiplayer version using WebSockets.
- Add animations or sounds for user interaction.
## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.
