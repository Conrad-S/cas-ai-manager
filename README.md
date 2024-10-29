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

This is an Azure Container Apps application that manages all aspects of connecting to and calling Azure OpenAI endpoints. The result is that applications and gateways no longer have to handle endpoints or keys.

For example, one approach is to enbed XML policies into Azure API management:

```xml
<policies>
    <inbound>
        <base />
    </inbound>
    <backend>
        <base />
    </backend>
    <outbound>
        <base />
    </outbound>
    <on-error>
        <retry condition="@(context.Response.StatusCode != 200 || context.Response.StatusCode == 408)" count="1" interval="5">
            <send-request mode="new" response-variable-name="fallbackResponse" timeout="10" ignore-error="false">
                <set-url>{{Fallback_OpenAI_Endpoint}}</set-url>
                <set-method>POST</set-method>
                <set-header name="Authorization" exists-action="override">
                    <value>{{Fallback_Authorization_Token}}</value>
                </set-header>
                <set-body>@(context.Request.Body.As<string>(preserveContent: true))</set-body>
            </send-request>
        </retry>
        <set-variable name="response" value="@((string)context.Variables["fallbackResponse"] ?? "Fallback service response unavailable")" />
        <return-response>
            <set-status code="200" reason="OK" />
            <set-body>@((string)context.Variables["response"])</set-body>
        </return-response>
    </on-error>
    <backend>
        <base />
        <send-request mode="new" response-variable-name="primaryResponse" timeout="10" ignore-error="true">
            <set-url>{{Primary_OpenAI_Endpoint}}</set-url>
            <set-method>POST</set-method>
            <set-header name="Authorization" exists-action="override">
                <value>{{Primary_Authorization_Token}}</value>
            </set-header>
            <set-body>@(context.Request.Body.As<string>(preserveContent: true))</set-body>
        </send-request>
    </backend>
    <outbound>
        <choose>
            <when condition="@(context.Variables["primaryResponse"]?.StatusCode == 200)">
                <return-response>
                    <set-status code="200" reason="OK" />
                    <set-body>@(context.Variables["primaryResponse"].Body.As<string>())</set-body>
                </return-response>
            </when>
            <otherwise>
                <return-response>
                    <set-status code="503" reason="Service Unavailable" />
                    <set-body>@("Both primary and fallback OpenAI endpoints failed.")</set-body>
                </return-response>
            </otherwise>
        </choose>
    </outbound>
</policies>

```

```csharp
var payload = new
    {
        modelName = "exampleModel",
        modelVersion = "v1",
        modelRegion = "eastus",
        body = new {
            systemMessage = "Hello, this is the system message",
            userMessage = "Hi, I have a question about the model"
        }
    };
    var content = new StringContent(JsonConvert.SerializeObject(payload), Encoding.UTF8, "application/json");
    var response = await client.PostAsync(url, content);
    return await response.Content.ReadAsStringAsync();
```


**Features**
