use clap::Parser;

/// Run a pipeline defined in a YAML file.
#[derive(Parser, Debug)]
#[command(name = "forge", version, about)]
pub struct Cli {
    /// Path to the pipeline YAML file.
    #[arg(short, long)]
    pub file: String,
}
