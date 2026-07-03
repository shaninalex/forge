use crate::pipeline::Pipeline;
use clap::Parser;
use std::fs;

pub mod step;
pub mod pipeline;
pub mod action;

/// Run a pipeline defined in a YAML file.
#[derive(Parser, Debug)]
#[command(version, about)]
struct Cli {
    /// Path to the pipeline YAML file.
    #[arg(short, long)]
    file: String,
}

fn main() {
    let cli = Cli::parse();

    let content = fs::read_to_string(&cli.file).unwrap();

    let p: Pipeline = serde_yaml::from_str(&content).expect("error");

    println!("Pipeline: {:?}", p.pipeline);
    for step in p.steps {
        println!("Process step: {:?}", step.id);
        match step.action.process() {
            Ok(response) => {
                println!("Status: {}", response.status);
                println!("Response:\n{}", response.body);
            }
            Err(e) => eprintln!("Action failed: {}", e),
        }
    }
}
