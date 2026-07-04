use crate::action::ActionResponse;

/// The result of running a whole pipeline.
///
/// This is the seam between the engine and any interface. Its shape is kept
/// deliberately small for now and will grow as assertions and richer output
/// land (see the architecture plan).
#[derive(Debug)]
pub struct RunReport {
    pub pipeline: String,
    pub steps: Vec<StepOutcome>,
}

/// What happened for a single step.
#[derive(Debug)]
pub struct StepOutcome {
    pub id: String,
    /// The action's response, or a stringified error if it failed.
    pub result: Result<ActionResponse, String>,
}
