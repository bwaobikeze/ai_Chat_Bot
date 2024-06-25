import gradio as gr
import ollama as ol

def greet(name, intensity):
    return "Hello," + name + "!" * int(intensity)
    
demo = gr.Interface(fn=greet, inputs=["text", "slider"], outputs="text")
demo.launch()

