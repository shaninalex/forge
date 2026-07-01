use std::collections::HashMap;
use serde::Deserialize;
use http_type::{serde_json, Method};

#[derive(Debug, Deserialize)]
#[serde(tag = "type")]
pub enum Action {
    #[serde(rename = "http")]
    Http {
        method: Method, // GET, POST
        url: String,
        headers: Option<HashMap<String, String>>,
        body: Option<serde_json::Value>,
    },

    #[serde(rename = "graphql")]
    GraphQL {
        url: String,
        query: String,
        variables: Option<serde_json::Value>,
        headers: Option<HashMap<String, String>>,
    },

    #[serde(rename = "webhook_listener")]
    WebhookListener {
        port: u16,
        timeout_seconds: u64,
    },

    #[serde(rename = "mcp_call")]
    McpCall {
        server: String,
        prompt: String,
    }
}

#[derive(Debug, Deserialize)]
pub struct Assertion {
    pub expression: String,
}

#[derive(Debug, Deserialize)]
pub struct Expectation {
    pub selector: String,
    pub value: String,
}


#[derive(Debug, Deserialize)]
pub struct Step {
    pub id: String,

    #[serde(flatten)]
    pub action: Action,

    pub asserts: Option<Vec<Assertion>>,
    pub expect_match: Option<Vec<Expectation>>,
}