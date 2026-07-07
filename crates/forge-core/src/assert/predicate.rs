use serde_json::Value;

#[derive(Debug, Clone, PartialEq)]
pub enum Predicate {
    // equality
    Eq(Value),
    Neq(Value),

    // ordering — numbers, durations
    Gt(Value),
    Gteq(Value),
    Lt(Value),
    Lteq(Value),

    // substring / element membership
    Contains(Value),
    NContains(Value),

    // affixes (strings only)
    Starts(String),
    NStarts(String),
    Ends(String),
    NEnds(String),

    // regex
    Matches(String),
    NMatches(String),

    // presence
    Exists,
    Empty,
    NEmpty,
}

impl Predicate {

    // Test the resolved (and possibly transformed) value against this
    // predicate.
    // NOTE: type-aware comparison.
    //  - Eq/Neq: compare Values directly (number vs string never match).
    //  - Gt/Gteq/Lt/Lteq: only for numbers
    //  - Contains: for strings
    //  - Starts/Ends/Matches: for strings.
    //  - Exists: actual.is_some(); Empty/NEmpty: zero-length vs present.
    pub fn test(&self, _actual: Option<&Value>) -> bool {
        todo!("apply operator to the resolved value")
    }
}
