// Assertions — the `<query> [transform] <predicate>` model.
//
// Examples:
//
// ```text
// status eq 200
// jsonpath "$[0].status" eq "PAID"
// jsonpath "$.items" count gteq 1
// header "Content-Type" contains "application/json"
// ```

mod predicate;
mod query;
mod transform;

pub use predicate::Predicate;
pub use query::Query;
pub use transform::Transform;

use crate::action::ActionResponse;

#[derive(Debug, Clone, PartialEq)]
pub struct Assertion {
    pub query: Query,
    pub transform: Option<Transform>,
    pub predicate: Predicate,
}

#[derive(Debug, Clone, PartialEq)]
pub enum ParseError {
    Todo,
}

#[derive(Debug, Clone)]
pub struct AssertionOutcome {
    pub passed: bool,
}

impl Assertion {
    pub fn parse(_line: &str) -> Result<Self, ParseError> {
        todo!("tokenize `<query kind> [arg] [transform] <operator> [value]`")
    }

    // Evaluate this assertion against a response.
    pub fn evaluate(&self, response: &ActionResponse) -> AssertionOutcome {
        let resolved = self.query.resolve(response);
        let value = match &self.transform {
            Some(t) => t.apply(resolved),
            None => resolved,
        };
        AssertionOutcome {
            passed: self.predicate.test(value.as_ref()),
        }
    }
}
