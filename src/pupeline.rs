use crate::step::Step;
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Pipeline {
    pipeline: String,
    steps: Vec<Step>,
}
