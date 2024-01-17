use time::PrimitiveDateTime as DateTime;
use time::Duration;

// A gigasecond is one thousand million seconds
// That is a one with nine zeros after it.
const GIGASECOND: i64 = 1_000_000_000;

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    start + Duration::seconds(GIGASECOND)
}
