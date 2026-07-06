use super::Assertion;
use crate::action::Action;
use serde::Deserialize;
use std::collections::HashMap;

#[derive(Debug, Deserialize)]
pub struct Step {
    pub id: String,

    #[serde(flatten)]
    pub action: Action,

    pub asserts: Option<Vec<Assertion>>,
    // pub expect_match: Option<HashMap<String, String>>,
}
