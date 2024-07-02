import gradio as gr
import requests

def LLMChatFunction(user_input, history):
    # history is a list of previous messages in the chat
    message=[{'role': 'user', 'content': user_input}]
    try:
        # model='llama3' is the model name
        url= "http://ollama:11434/v1/chat"
        headers = {'Content-Type': 'application/json'}
        data = {
            "model": "llama3",
            "messages": message
        }
        response = requests.post(url, headers=headers, json=data)
        return str(response.json()['messages']['content'])
    except Exception as e:
        return e
    
demo = gr.ChatInterface(LLMChatFunction).launch()


    
