use crate::pupeline::Pipeline;
use std::fs;

pub mod step;
pub mod pupeline;

fn main() {
    let content = fs::read_to_string("docs/example.yaml").unwrap();
    let p: Pipeline = serde_yaml::from_str(&content).expect("error");

    println!("Parsed file data: {:?}", p);
}
