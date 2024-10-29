from flask import Flask, request, jsonify
import os
from openai import AzureOpenAI

app = Flask(__name__)

# Set up the Azure OpenAI API credentials
client = AzureOpenAI(
    api_key=os.getenv("AZURE_OPENAI_API_KEY"),
    azure_endpoint=os.getenv("AZURE_OPENAI_ENDPOINT"),
    api_version="2023-07-01-preview"  # Update with the correct version
)

deployment_id = os.getenv("AZURE_OPENAI_DEPLOYMENT_ID")  # Your model deployment ID

@app.route('/process', methods=['POST'])
def process():
    try:
        data = request.json
        assistant_message = data.get("message", "What can you help me with today?")
        
        # Use the new API structure
        response = client.chat.completions.create(
            model=deployment_id,
            messages=[{"role": "user", "content": assistant_message}],
            max_tokens=150,
            temperature=0.7,
        )
        
        # Extract the response text
        response_text = response['choices'][0]['message']['content'].strip()
        
        # Return the response as JSON
        return jsonify({"response": response_text})
    except Exception as e:
        # Log the exception
        app.logger.error(f"Error processing request: {e}")
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8083)