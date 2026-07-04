use forge_core::RunReport;

/// Render a run report as plain, human-readable text (the default output).
pub fn render(report: &RunReport) {
    println!("Pipeline: {:?}", report.pipeline);
    for step in &report.steps {
        println!("Process step: {:?}", step.id);
        match &step.result {
            Ok(response) => {
                println!("Status: {}", response.status);
                println!("Response:\n{}", response.body);
            }
            Err(err) => eprintln!("Action failed: {}", err),
        }
    }
}
