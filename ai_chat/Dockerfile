FROM python:3.9-slim

WORKDIR /ai_chat

COPY . .

RUN pip install -r requirements.txt

EXPOSE 7860

CMD ["python", "greet.py"]