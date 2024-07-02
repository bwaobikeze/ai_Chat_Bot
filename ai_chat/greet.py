import gradio as gr
import ollama as ol

def LLMChatFunction(user_input, history):
    # history is a list of previous messages in the chat
    message=[{'role': 'user', 'content': user_input}]
    ollama_server_Url = 'http://ollama:11434'
    
    ollama_client = ol.Client(ollama_server_Url)
    try:
        response = ollama_client.chat(model='llama3', messages=message)
        return str(response['message']['content'])
    except Exception as e:
        return e
    
demo = gr.ChatInterface(LLMChatFunction).launch(server_name="0.0.0.0", server_port=7860)


    
