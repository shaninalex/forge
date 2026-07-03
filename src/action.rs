use http_type::serde_json;
use reqwest::blocking::Client;
use reqwest::Method;
use serde::Deserialize;
use std::collections::HashMap;

#[derive(Debug, Deserialize)]
pub struct HttpAction {
    pub method: String, // GET, POST...
    pub url: String,
    pub query: Option<HashMap<String, String>>,
    pub headers: Option<HashMap<String, String>>,
    pub body: Option<serde_json::Value>,
}

#[derive(Debug, Deserialize)]
#[serde(tag = "type")]
pub enum Action {
    #[serde(rename = "http")]
    Http(HttpAction),

    // NOTE: we do not deal with that in MVP stage
    // #[serde(rename = "graphql")]
    // GraphQL {
    //     url: String,
    //     query: String,
    //     variables: Option<serde_json::Value>,
    //     headers: Option<HashMap<String, String>>,
    // },

    // NOTE: we do not deal with that in MVP stage
    // #[serde(rename = "webhook_listener")]
    // WebhookListener {
    //     port: u16,
    //     timeout_seconds: u64,
    // },

    // NOTE: we do not deal with that in MVP stage
    // #[serde(rename = "mcp_call")]
    // McpCall {
    //     server: String,
    //     prompt: String,
    // },
}

#[derive(Debug)]
pub struct ActionResponse {
    pub status: reqwest::StatusCode,
    pub body: String,
}

impl Action {
    pub fn process(&self) -> Result<ActionResponse, Box<dyn std::error::Error>> {
        match self {
            Action::Http(http) => process_http(http),
        }
    }
}


fn process_http(action: &HttpAction) -> Result<ActionResponse, Box<dyn std::error::Error>> {
    let client = Client::new();

    let method: Method = action.method.to_uppercase().parse()?;
    let mut request = client.request(method, &action.url);

    if let Some(query) = &action.query {
        request = request.query(query);
    }

    if let Some(headers) = &action.headers {
        for (key, value) in headers {
            request = request.header(key, value);
        }
    }

    if let Some(body) = &action.body {
        request = request.json(body);
    }

    let request = request.build()?;
    println!("URL: {}", request.url());
    let response = client.execute(request)?;

    let status = response.status();
    let body = response.text()?;

    Ok(ActionResponse { status, body })
}