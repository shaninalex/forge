use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Assertion {
    pub expression: String,
}
