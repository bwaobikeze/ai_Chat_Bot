FROM python:3.9-slim

WORKDIR /ai_chat

COPY . .

RUN pip install -r Requirements.txt

EXPOSE 7860

CMD ["python","-u" ,"ai_chatbot.py"]