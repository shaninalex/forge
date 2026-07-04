mod args;
mod render;

use args::Cli;
use clap::Parser;
use forge_core::{Pipeline, engine};
use std::fs;
use std::process;

fn main() {
    let cli = Cli::parse();

    let content = match fs::read_to_string(&cli.file) {
        Ok(content) => content,
        Err(err) => {
            eprintln!("failed to read {}: {}", cli.file, err);
            process::exit(1);
        }
    };

    let pipeline = match Pipeline::from_yaml_str(&content) {
        Ok(pipeline) => pipeline,
        Err(err) => {
            eprintln!("failed to parse pipeline: {}", err);
            process::exit(1);
        }
    };

    let report = engine::run(&pipeline);
    render::text::render(&report);
}
