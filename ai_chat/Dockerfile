FROM python:3.9-slim

WORKDIR /ai_chat

COPY . .

RUN pip install gradio ollama

CMD ["python", "greet.py"]