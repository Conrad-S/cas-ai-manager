from flask import Flask, request, jsonify
import os
import asyncio
import semantic_kernel as sk
from semantic_kernel.connectors.ai.open_ai import AzureChatCompletion
from semantic_kernel.utils.logging import setup_logging
from semantic_kernel.contents.chat_history import ChatHistory

app = Flask(__name__)

# Initialize Semantic Kernel with Azure OpenAI
api_key = os.getenv("AZURE_OPENAI_API_KEY")
endpoint = os.getenv("AZURE_OPENAI_ENDPOINT")
deployment_name = os.getenv("AZURE_OPENAI_DEPLOYMENT_ID")

# Create the Kernel and AzureChatCompletion instance
kernel = sk.Kernel()
azure_chat_completion = AzureChatCompletion(
    deployment_name=deployment_name,
    api_key=api_key,
    base_url=endpoint,
    
)

# Set up logging for the Semantic Kernel
setup_logging()

@app.route('/process', methods=['POST'])
def process():
    try:
        data = request.json
        user_input = data.get("message", "What can you help me with today?")

        # Create a new history for the conversation
        history = ChatHistory()

        # Add the user's message to the history
        history.add_user_message(user_input)
        
        result = azure_chat_completion.AsyncAzureOpenAI = asyncio.run(azure_chat_completion.complete_chat(            
            chat_history=history,
            settings={
                "max_tokens": 150
            }
        ))

        # Await the result of complete_chat asynchronously
        result = asyncio.run(azure_chat_completion.complete_chat(            
            chat_history=history,
            settings={
                "max_tokens": 150
            }
        ))

        # Return the processed result
        return jsonify({"response": result})

    except Exception as e:
        # Handle errors
        app.logger.error(f"Error processing request: {e}")
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8082)