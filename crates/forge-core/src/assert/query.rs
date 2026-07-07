// The query — *where* to look and *what* to point at.

use crate::action::ActionResponse;
use serde_json::Value;

#[derive(Debug, Clone, PartialEq)]
pub enum Query {
    // HTTP status code -> number.
    Status,
    // Round-trip time in milliseconds -> number.
    Duration,
    // Raw response body -> string.
    Body,
    // A single response header by name -> string.
    Header(String),
    // A JSONPath expression over a JSON body -> the matched value(s).
    JsonPath(String),
    // A regex capture over a text body -> the captured string.
    Regex(String),
}

impl Query {
    pub fn resolve(&self, _response: &ActionResponse) -> Option<Value> {
        todo!("resolve the query against the response")
    }
}
