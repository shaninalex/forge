//! The pipeline *spec* — the plain, deserialized shape of a pipeline
//! definition. These types describe what the user wrote; they carry no
//! execution behavior.

mod assertion;
mod pipeline;
mod step;

pub use assertion::Assertion;
pub use pipeline::Pipeline;
pub use step::Step;
