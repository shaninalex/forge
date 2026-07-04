use super::{ActionResponse, Executor};
use http_type::serde_json;
use reqwest::Method;
use reqwest::blocking::Client;
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

impl Executor for HttpAction {
    fn execute(&self) -> Result<ActionResponse, Box<dyn std::error::Error>> {
        let client = Client::new();

        let method: Method = self.method.to_uppercase().parse()?;
        let mut request = client.request(method, &self.url);

        if let Some(query) = &self.query {
            request = request.query(query);
        }

        if let Some(headers) = &self.headers {
            for (key, value) in headers {
                request = request.header(key, value);
            }
        }

        if let Some(body) = &self.body {
            request = request.json(body);
        }

        let response = client.execute(request.build()?)?;
        let status = response.status();
        let body = response.text()?;

        Ok(ActionResponse { status, body })
    }
}
