<a href="https://">
    <img src="https://aka.ms/deploytoazurebutton" alt="Deploy to Azure">
</a>

(Button not yet enabled)

### Development-Status
```json
{
  "version": "0.1",
  "working": "Yes",
  "documentation": "No",
  "developer-required": "Yes",
  "distributed-publicly": "No"
}
```

**Project description**

This is an application that handles connectivity and calls to Azure OpenAI endpoints, eliminating the need to add complex code to applications and gateways.

The app:
 - Handles connectivity to Azure OpenAI endpoints. The caller only needs to pass in the model name, model version, and desired region.
 - Handles calls to Azure OpenAI endpoints.
 - Enables the use of Semantic Kernel
 - Enables the use of Assistant API

For example, once implemented, the only code needed to call an Azure OpenAI endpoint is as follows (example in C#):

```csharp
var payload = new {modelName = "exampleModel", modelVersion = "v1", modelRegion = "eastus",
    body = new {systemMessage = "You are a helpful assistant", userMessage = "How is the weather in Boston?"}
};

    var content = new StringContent(JsonConvert.SerializeObject(payload), Encoding.UTF8, "application/json");
    var response = await client.PostAsync(url, content);
    return await response.Content.ReadAsStringAsync();
```

**Architecture**
The application is implemented as a series of microservices deployed as an Azure Container App.

--- TODO...
