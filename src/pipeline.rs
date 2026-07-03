use crate::step::Step;
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Pipeline {
    pub pipeline: String,
    pub steps: Vec<Step>,
}
