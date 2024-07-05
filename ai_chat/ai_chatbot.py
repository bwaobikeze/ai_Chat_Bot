import gradio as gr
import ollama as ol

# Defines a function that takes a user input and returns a chatbot response
def LLMChatFunction(user_input, model_name, history):
    message=[{'role': 'user', 'content': user_input}]
    ollama_server_Url = 'http://ollama:11434'
    ollama_client = ol.Client(ollama_server_Url)
    try:
        response = ollama_client.chat(model= model_name, messages=message)
        bot_response = str(response['message']['content'])
        history.append((user_input, bot_response))
        return history
        
    except Exception as e:
        return e
    
#Defines a function that returns the list of available models 
def ListOfmodels():
    available_models = ['llama3']
    return available_models


#Defines a function that pulls the model from ollama server   
def PullOllamaModel(modelname):
    ollama_server_Url = 'http://ollama:11434'
    ollama_client = ol.Client(ollama_server_Url)
    try:
        response = ollama_client.pull(model= modelname)
        return response
    except Exception as e:
        return e

# Create a Gradio interface
with gr.Blocks() as demo:
    chatbot = gr.Chatbot()
    with gr.Row():
       with gr.Column(scale=4):
            user_input = gr.Textbox(show_label=False, placeholder="Type your message here...")
       with gr.Column(scale=1):
           send_Button = gr.Button("Send")
    

    available_models = ListOfmodels()
    model_dropdown = gr.Dropdown(available_models, label="Select Model")


    PullOllamaModel(model_dropdown)

    chat_history = gr.State([])
    

    send_Button.click(LLMChatFunction, inputs=[user_input,model_dropdown,chat_history], outputs=[chatbot])

demo.launch(server_name="0.0.0.0", server_port=7860)