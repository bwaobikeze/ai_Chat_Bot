FROM python:3.9-slim

WORKDIR /ai_chat

COPY . .

RUN pip install gradio ollama

EXPOSE 7860

CMD ["python", "greet.py"]