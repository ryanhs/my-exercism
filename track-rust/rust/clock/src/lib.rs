
#[derive(Debug)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

struct MinutesHourCalcResults {
    minutes: i32,
    hour_surplus: i32,
}

fn calc_minutes(base_minutes: i32, diff_minutes: i32) -> MinutesHourCalcResults {
    let new_minutes = base_minutes + diff_minutes;
    
    // positif
    if new_minutes >= 0 {
        let mod_minutes = new_minutes % 60;
        let hour_surplus = (new_minutes / 60) as i32;
        return  MinutesHourCalcResults { minutes: mod_minutes, hour_surplus };
    }
    
    // negative
    let mut mod_minutes = new_minutes % 60;
    let mut hour_surplus: i32 = (new_minutes / 60) as i32;

    if mod_minutes < 0 {
        hour_surplus -= 1;
        mod_minutes = 60 + mod_minutes;
    }

    return MinutesHourCalcResults{ minutes: mod_minutes, hour_surplus };
}

fn calc_hours(base_hours: i32, diff_hours: i32) -> i32 {
    let new_hours = base_hours + diff_hours;
    
    // positif
    if new_hours >= 0 {
        return new_hours % 24;
    }
    
    // negative
    let mod_hours = new_hours % 24;

    if mod_hours < 0 {
        return 24 + mod_hours;
    }

    return 0
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Clock {
            hours,
            minutes,
        }.add_minutes(0)
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let res = calc_minutes(self.minutes, minutes);
        // println!("add_minutes: {} add {}", self.minutes, minutes);
        // println!("res: {}, {}", res.minutes, res.hour_surplus);

        let fixed_hours: i32 = calc_hours(self.hours, res.hour_surplus);

        Self {
            hours: fixed_hours,
            minutes: res.minutes
        }
    }

    pub fn to_string(&self) -> String {
        format!("{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}

impl PartialEq<Clock> for Clock {
    fn eq(&self, other: &Clock) -> bool {
        self.hours == other.hours && self.minutes == other.minutes
    }
}
