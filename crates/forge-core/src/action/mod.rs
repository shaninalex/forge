//! Actions — the executable unit inside a step.
//!
//! Each action kind lives in its own module and owns both its parsed spec and
//! its execution, so adding a new kind (graphql, webhook_listener, mcp_call…)
//! is a localized change: add a module, add a variant, implement [`Executor`].

mod http;

use std::collections::HashMap;
pub use http::HttpAction;

use serde::Deserialize;

/// A single unit of work within a step, tagged in YAML by its `type` field.
#[derive(Debug, Deserialize)]
#[serde(tag = "type")]
pub enum Action {
    #[serde(rename = "http")]
    Http(HttpAction),
    // Deferred (see architecture plan): graphql, webhook_listener, mcp_call.
}

/// The outcome of executing an action.
#[derive(Debug)]
pub struct ActionResponse {
    pub status: reqwest::StatusCode,
    pub body: String,
    pub headers: HashMap<String, String>,
    pub duration: u128, // milliseconds
}

/// Anything the engine can execute.
pub trait Executor {
    fn execute(&self) -> Result<ActionResponse, Box<dyn std::error::Error>>;
}

impl Executor for Action {
    fn execute(&self) -> Result<ActionResponse, Box<dyn std::error::Error>> {
        match self {
            Action::Http(http) => http.execute(),
        }
    }
}
