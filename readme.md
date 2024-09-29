
# Telegram Bot in Go

This repository contains a simple Telegram bot created using Go. The bot utilizes environment variables for configuration and responds to various commands.

## Features

- Responds to a set of predefined commands
- Utilizes inline and reply buttons for user interaction
- Loads configuration from a `.env` file
- Easy to extend with additional commands

## Technologies Used

- Go
- [telebot](https://github.com/tucnak/telebot) - A simple and powerful Telegram bot API library for Go
- [godotenv](https://github.com/joho/godotenv) - Loads environment variables from a `.env` file

## Getting Started

### Prerequisites

- Go installed on your machine
- A Telegram bot token (you can create a new bot using [BotFather](https://core.telegram.org/bots#botfather))

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/your-repo-name.git
   cd your-repo-name
   ```

2. Create a `.env` file in the root directory of the project with the following content:

   ```env
   BOT_TOKEN=your_telegram_bot_token_here
   ```

3. Install the required dependencies:

   ```bash
   go mod tidy
   ```

4. Run the bot:

   ```bash
   go run main.go
   ```

## Commands

The bot responds to the following commands:

- `/start`: Initializes the bot and greets the user.
- `/test1`: Provides a description for test command 1.
- `/test2`: Provides a description for test command 2.
- `/test3`: Provides a description for test command 3.
- `/test4`: Provides a description for test command 4.
- `/test5`: Provides a description for test command 5.
- `/test6`: Provides a description for test command 6.

Additionally, it can handle simple text messages and provide relevant responses.

## Code Overview

The main function initializes the bot and sets up command handlers. Key components include:

- **Environment Variable Loading**: The bot uses `godotenv` to load the Telegram bot token from the `.env` file.
  
  ```go
  godotenv.Load()
  ```

- **Command Registration**: Each command is defined as a `tb.Command`, and they are registered with the bot.

  ```go
  var (
      CommandTest1 = tb.Command{Text: "test1", Description: "test1"}
      ...
      commands     = []tb.Command{CommandTest1, CommandTest2, ...}
  )
  ```

- **Button Setup**: Reply and inline buttons are defined to enhance user interaction.

  ```go
  btnHelp     = menu.Text("ℹ Help")
  btnSettings = menu.Text("⚙ Settings")
  ```

- **Handling Commands**: Each command is handled with a function that sends a corresponding message back to the user.

  ```go
  b.Handle("/start", func(c tb.Context) error {
      return c.Send("Hello world")
  })
  ```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.


