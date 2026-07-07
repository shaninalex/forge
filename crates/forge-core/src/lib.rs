//! ForgeCore — the pipeline engine.
//!
//! This crate is interface-agnostic: it parses pipeline definitions and
//! executes their steps. It performs no argument parsing and owns no
//! user-facing presentation — that belongs to the interfaces built on top of
//! it (the reference CLI in this repo, plus separate TUI/GUI projects).
//!
//! The typical flow for any interface is:
//!
//! 1. Load a [`Pipeline`] (e.g. [`Pipeline::from_yaml_str`]).
//! 2. Hand it to [`engine::run`], which returns a [`RunReport`].
//! 3. Render the report however the interface likes.

pub mod action;
pub mod assert;
pub mod engine;
pub mod model;
pub mod report;

// Re-export the core types so consumers can `use forge_core::{Pipeline, ...}`
// without reaching into module paths.
pub use action::{Action, ActionResponse, Executor, HttpAction};
pub use model::{Assertion, Pipeline, Step};
pub use report::{RunReport, StepOutcome};
