let playerScore = 0;
let aiScore = 0;

// Función para determinar el ganador
function determineWinner(playerChoice, aiChoice) {
    let resultMessage = "";

    if (playerChoice === aiChoice) {
        resultMessage = "It's a Tie!";
    } else if (
        (playerChoice === "Rock" && aiChoice === "Scissors") ||
        (playerChoice === "Paper" && aiChoice === "Rock") ||
        (playerChoice === "Scissors" && aiChoice === "Paper")
    ) {
        resultMessage = "You Win!";
        playerScore++; // El jugador gana, aumentamos su puntaje
    } else {
        resultMessage = "You Lose!";
        aiScore++; // La IA gana, aumentamos su puntaje
    }

    // Actualiza el mensaje de resultado
    const resultElement = document.getElementById("resultMessage");
    resultElement.innerHTML = resultMessage;

    // Remueve clases anteriores y añade la clase correspondiente
    resultElement.classList.remove("win", "lose", "tie");
    if (resultMessage === "You Win!") {
        resultElement.classList.add("win");
    } else if (resultMessage === "You Lose!") {
        resultElement.classList.add("lose");
    } else {
        resultElement.classList.add("tie");
    }

    // Actualiza el puntaje
    updateScore();
}

// Función para actualizar el puntaje en la interfaz
function updateScore() {
    document.querySelector('.score-board').innerHTML = `Player ${playerScore} - AI ${aiScore}`;
}

// Función para elegir una carta al hacer clic
function chooseCard(choice) {
    const playerChoice = choice;
    const aiChoice = getAIChoice();
    determineWinner(playerChoice, aiChoice);

    document.getElementById('playerCard').innerHTML = getEmojiForChoice(playerChoice);
    document.getElementById('aiCard').innerHTML = getEmojiForChoice(aiChoice);
}

// Función para obtener la elección de la IA
function getAIChoice() {
    const choices = ['Rock', 'Paper', 'Scissors'];
    const randomIndex = Math.floor(Math.random() * choices.length);
    return choices[randomIndex];
}

// Función para obtener el emoji de la elección
function getEmojiForChoice(choice) {
    switch (choice) {
        case 'Rock':
            return '🪨';
        case 'Paper':
            return '📄';
        case 'Scissors':
            return '✂️';
        default:
            return '❓';
    }
}
