#### ai chat model

This project uses Docker to containerize a Python application that leverages the Ollama model for generating responses in a Gradio chatbot. Follow the instructions below to set up and run the project.

## Prerequisites

- [Docker](https://www.docker.com/) installed on your machine
- An internet connection to pull the model from Ollama

## Project Files

- `.dockerignore`: Specifies files to be ignored by Docker.
- `docker-compose.yaml`: Docker Compose configuration file.
- `Dockerfile`: Instructions to build the Docker image.
- `ai_chatbot.py.py`: Python script to run the application.
- `requirements.txt`: Python dependencies.

## Setup Instructions

1. **Clone the Repository**

   First, clone this repository to your local machine:

   ```sh
   git clone <repository_url>
   cd <repository_directory>

2. **Start the application using Docker Compose:**

   Start the application using Docker Compose:

   ```sh
   docker-compose up --build

3. **Access the Application**

   Once the application is running and the model is pulled, you can access it through your web browser at http://localhost:7860.
   
4. **Select model from the "Select model" dropdown**

   After starting the application, you need to pull the model from Ollama.
   Click on the dropdown and select your model.

5. **use AI chat bot**

   Begin using the AI chat bot.

## Project Structure

- `.dockerignore`: Contains patterns to ignore files and directories in Docker builds.
- `docker-compose.yaml`: Defines services, networks, and volumes for the Docker Compose setup.
- `Dockerfile`: Contains instructions to set up the Docker image, including installing dependencies and copying project files.
- `ai_chatbot.py`: Python script that sets up and runs the Gradio interface using the Ollama model.
- `requirements.txt`: Lists the Python packages required for the project (`gradio` and `ollama`).