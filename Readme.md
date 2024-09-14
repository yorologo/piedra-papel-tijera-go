# Rock, Paper, Scissors - Web Game in Go

## Overview

This project is a simple **Rock, Paper, Scissors** web game developed using the Go programming language. The player can choose one of three options (Rock, Paper, or Scissors) to play against an AI that randomly selects its own option. The game determines the winner based on the standard **Rock, Paper, Scissors** rules and displays the result on the web page, along with an improved visual interface.

## Features

- **Web-based interface**: The player interacts with the game through a modern web browser.
- **Random AI selection**: The AI randomly selects one of the three options in each game.
- **Game logic**: The server processes the game, compares the player’s choice with the AI’s, and determines the outcome (win, lose, or draw).
- **Scoring system**: The game keeps a continuous record of player and AI wins.
- **Result display with color change**: The result of each round is displayed on the web page, with the result message changing colors: green if the player wins, red if they lose, and black if it’s a tie.
- **Responsive design**: The game is responsive and adapts well to both desktop and mobile screens.

## Prerequisites

- **Go**: Version 1.16 or higher. You can download it from [the official Go website](https://golang.org/dl/).
- **Modern web browser**: Such as Chrome, Firefox, Safari, etc.

## Project Structure

```
piedra-papel-tijera/
├── main.go          # Main application file
├── index.html       # HTML template for the game interface
├── css/             # CSS styles for the game interface
│   └── styles.css   # Custom styles for the game
├── js/              # JavaScript logic for the game interface
│   └── scripts.js   # Game interaction and result logic
├── go.mod           # Go module definition
└── LICENSE          # Project license
```

## How to Run

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/piedra-papel-tijera-go.git
   cd piedra-papel-tijera-go
   ```

2. **Install dependencies**:

   Run the following command to download the necessary dependencies:

   ```bash
   go mod download
   ```

3. **Run the application**:

   Inside the project directory, start the web server with:

   ```bash
   go run main.go
   ```

4. **Access the game**:

   Open your browser and navigate to `http://localhost:8080/play` to play the game.

## Code Overview

### **main.go**

This is the main file that sets up the HTTP server, defines the game logic, and handles player requests.

### **index.html**

This file contains the HTML structure for the game interface. It includes interactive elements for the player to select their move and displays the result with updated styles based on the game outcome.

**Features**:

- **Bootstrap design**: Uses Bootstrap classes to create a responsive and visually appealing layout.
- **Player hands**: Cards are displayed in a fan-like arrangement for both the player and the AI.
- **Result element**: Displays the result of each round, changing the text color according to the outcome.

### **styles.css**

The CSS file contains the styling for the game interface.

**Key elements**:

- **Card styles**: The cards represent the player’s and AI’s choices, arranged in a fan-like layout with rotation and overlap effects.
- **Hover effects**: Cards enlarge and show a shadow when hovered over, enhancing interactivity.
- **Responsiveness**: Includes media queries to adjust the layout on mobile devices, ensuring an optimal experience across all screen sizes.
- **Result colors**: Defines classes to change the color of the result message according to the game outcome (green for win, red for lose, black for draw).

### **scripts.js**

This file contains the JavaScript logic that manages interaction between the user and the game.

**Main functions**:

- **Player selection**: Records the player’s choice when clicking on a card.
- **AI selection**: Randomly generates the AI’s choice.
- **Determine winner**: Compares the choices and determines the result, updating the score.
- **UI update**: Displays the choices on the board and changes the result message color accordingly.

### **Game Logic**

The game follows the traditional **Rock, Paper, Scissors** rules:

- **Rock** beats **Scissors**.
- **Paper** beats **Rock**.
- **Scissors** beats **Paper**.
- If both the player and the AI choose the same option, the result is a **Draw**.

## Future Improvements

- **Implement rounds**: Add the possibility to play best out of 3 or 5 rounds.
- **Improve user interface**: Add more CSS animations and transitions for a more engaging experience.
- **Multiplayer mode**: Allow two players to play against each other using WebSockets.
- **Sound effects**: Add sounds for interactions and results.
- **Unit tests**: Implement tests to ensure the game logic works correctly.
- **Improve AI**: Make the AI learn from the player’s behavior to offer a greater challenge.

## How to Contribute

Contributions are welcome. If you wish to contribute, please follow these steps:

1. Fork the repository.
2. Create a branch for your feature or fix (`git checkout -b feature/new-feature`).
3. Make your changes and commit them with descriptive messages.
4. Submit a pull request detailing the changes made.

## Screenshots

![example_img1](https://drive.google.com/uc?export=view&id=1ih94hza_GAtCrjjXkkryb7rfCdlvU0kF)
![example_img2](https://drive.google.com/uc?export=view&id=1idxGfZSVGMK5qF6fIvMbripWFRj1hWm9)
![example_gif1](https://drive.google.com/uc?export=view&id=1ilcOGCr_gtOCxkZ-RGwyHHBPvXikbcda)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
