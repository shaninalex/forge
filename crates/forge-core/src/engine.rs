use crate::action::Executor;
use crate::model::Pipeline;
use crate::report::{RunReport, StepOutcome};

/// Execute every step of a pipeline in order and collect the results.
///
/// The engine performs no presentation: it returns a [`RunReport`] that an
/// interface (CLI/TUI/GUI) is free to render however it likes.
pub fn run(pipeline: &Pipeline) -> RunReport {
    let mut steps = Vec::with_capacity(pipeline.steps.len());

    for step in &pipeline.steps {
        let result = step.action.execute().map_err(|err| err.to_string());
        if step.asserts.is_some() {
            // check condition
            // if validation failed => break
        }

        steps.push(StepOutcome {
            id: step.id.clone(),
            result,
        });
    }

    RunReport {
        pipeline: pipeline.pipeline.clone(),
        steps,
    }
}
