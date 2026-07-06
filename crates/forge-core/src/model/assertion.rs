use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize)]
pub struct Assertion {
    pub expression: String,

    // query: Query,        // Status | Duration | Body | Header(String) | JsonPath(String) | Regex(String)
    // predicate: Predicate // Exists | Empty | NEmpty | Eq(Value) | Gt(Value) | Contains(Value) | ...
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
pub enum AssertionCondition {
    #[serde(rename = "eq")]
    IsEqualTo,

    #[serde(rename = "neq")]
    IsNotEqualTo,

    #[serde(rename = "gt")]
    IsGreater,

    #[serde(rename = "gteq")]
    IsGreaterOrEqual,

    #[serde(rename = "lt")]
    IsLower,

    #[serde(rename = "lteq")]
    IsLowerOrEqual,

    #[serde(rename = "contains")]
    IsContain,

    #[serde(rename = "ncontains")]
    IsNotContain,

    #[serde(rename = "starts")]
    IsStartWith,

    #[serde(rename = "nstarts")]
    IsNotStartWith,

    #[serde(rename = "ends")]
    IsEndWith,

    #[serde(rename = "nends")]
    IsNotEndWith,

    #[serde(rename = "empty")]
    IsEmpty,

    #[serde(rename = "nempty")]
    IsNotEmpty,

    #[serde(rename = "named")]
    IsNamed,

    #[serde(rename = "time_interval")]
    IsTimeInterval,

    #[serde(rename = "time_part")]
    IsTimePart,
}
