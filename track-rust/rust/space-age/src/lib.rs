// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

use std::collections::HashMap;

#[derive(Debug)]
pub struct Duration {
    seconds: f64,
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration {
            seconds: s as f64
        }
    }
}

pub trait Planet {
    fn years_during(d: &Duration) -> f64 {
        todo!("convert a duration ({d:?}) to the number of years on this planet for that duration");
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;


const EARTH_YEAR_IN_SECONDS: f64 = 31557600.0;

fn a_year_on_planet_map() -> HashMap<&'static str, f64> {
    return HashMap::from([
        ("Mercury", 0.2408467 * EARTH_YEAR_IN_SECONDS),
        ("Venus", 0.61519726 * EARTH_YEAR_IN_SECONDS),
        ("Earth", 1.0 * EARTH_YEAR_IN_SECONDS),
        ("Mars", 1.8808158 * EARTH_YEAR_IN_SECONDS),
        ("Jupiter", 11.862615 * EARTH_YEAR_IN_SECONDS),
        ("Saturn", 29.447498 * EARTH_YEAR_IN_SECONDS),
        ("Uranus", 84.016846 * EARTH_YEAR_IN_SECONDS),
        ("Neptune", 164.79132 * EARTH_YEAR_IN_SECONDS),
    ]);
}

impl Planet for Mercury {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Mercury").unwrap()
    }
}
impl Planet for Venus {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Venus").unwrap()
    }
}

impl Planet for Earth {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Earth").unwrap()
    }
}

impl Planet for Mars {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Mars").unwrap()
    }
}
impl Planet for Jupiter {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Jupiter").unwrap()
    }
}

impl Planet for Saturn {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Saturn").unwrap()
    }
}

impl Planet for Uranus {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Uranus").unwrap()
    }
}

impl Planet for Neptune {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / a_year_on_planet_map().get("Neptune").unwrap()
    }
}
