import gradio as gr
import ollama as ol

def LLMChatFunction(user_input, history):
    # history is a list of previous messages in the chat
    message=[{'role': 'user', 'content': user_input}]
    try:
        # model='llama3' is the model name
        response = ol.chat(model='llama3', messages=message)
        return str(response['message']['content'])
    except Exception as e:
        return e
    
demo = gr.ChatInterface(LLMChatFunction).launch(server_name="0.0.0.0", server_port=7860)

