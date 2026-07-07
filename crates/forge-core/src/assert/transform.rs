use serde_json::Value;

#[derive(Debug, Clone, PartialEq)]
pub enum Transform {
    Count,
}

impl Transform {
    pub fn apply(&self, _value: Option<Value>) -> Option<Value> {
        todo!("reduce the resolved value")
    }
}
