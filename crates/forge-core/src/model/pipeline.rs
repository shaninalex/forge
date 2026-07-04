use super::Step;
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Pipeline {
    pub pipeline: String,
    pub steps: Vec<Step>,
}

impl Pipeline {
    /// Parse a pipeline from a YAML document.
    ///
    /// Reading the document from disk (or anywhere else) is left to the
    /// caller, so the core stays free of any filesystem concern.
    pub fn from_yaml_str(yaml: &str) -> Result<Self, serde_yaml::Error> {
        serde_yaml::from_str(yaml)
    }
}
