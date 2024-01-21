
#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

#[derive(PartialEq)]
pub enum WhichThrow {
    FirstThrow,
    SecondThrow,
}


// https://www.bowlinggenius.com/

const TOTAL_PINS: u16 = 10;
const STRIKE: u16 = 10;
const NORMAL_FRAMES: usize = 10;

pub struct BowlingGame {
    frame_i : usize,
    which_throw: WhichThrow,
    pins_remaining: u16,
    pins_down: Vec<u16>,
}

impl BowlingGame {
    pub fn new() -> Self {
        BowlingGame {
            frame_i: 0,
            which_throw: WhichThrow::FirstThrow,
            pins_remaining: TOTAL_PINS,
            pins_down: Vec::with_capacity((NORMAL_FRAMES * 2) + 4), // +2 frame for double strike in the end
        }
    }

    pub fn is_done(&self) -> bool {
        if self.frame_i <= 9 {
            return false;
        }

        let last_throw_i: usize = (9 * 2) as usize;
        let last_frame_first: u16 = self.pins_down[last_throw_i];
        let last_frame_second: u16 = self.pins_down[last_throw_i + 1];
        let last_frame_final = last_frame_first + last_frame_second;

        // next 1 frame due to great score
        if self.frame_i == NORMAL_FRAMES {
            
            let is_first_throw = self.pins_down.len() == NORMAL_FRAMES * 2;

            // last frame is strike / spare
            if last_frame_final == TOTAL_PINS && is_first_throw {
                return false;
            }

            // last frame is strike exclusively, oke for 2 throw
            if last_frame_first == STRIKE {
                return false;
            }
        }

        // next 2 frame due to double strike score
        if self.frame_i == 11 && last_frame_first == TOTAL_PINS {
            
            let is_first_throw = self.pins_down.len() == 11 * 2;

            let last_bonus_frame: u16 = self.pins_down[last_throw_i + 2];

            if last_bonus_frame == STRIKE && is_first_throw {
                return false;
            }
        }

        true
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if pins > self.pins_remaining {
            return Err(Error::NotEnoughPinsLeft);
        }
        
        if self.is_done() {
            return Err(Error::GameComplete);
        }
        
        self.pins_down.push(pins);

        let mut make_new_frame = self.which_throw == WhichThrow::SecondThrow;

        // finisher throw, new frame
        if self.pins_remaining == pins {

            // strike / spare, 
            if self.which_throw == WhichThrow::FirstThrow {
                self.pins_down.push(0);
                make_new_frame = true;
            }
        }

        println!("# frame: {}, pins: {}", self.frame_i + 1, pins);
        self.print_score_board();

        // new frame
        if make_new_frame {
            self.pins_remaining = TOTAL_PINS;
            self.frame_i += 1;
            self.which_throw = WhichThrow::FirstThrow;
        }

        // next throw in current frame
        else if self.which_throw == WhichThrow::FirstThrow {
            self.which_throw = WhichThrow::SecondThrow;
            self.pins_remaining -= pins;
        }

        println!("");

        return Ok(());
    }

    pub fn print_score_board(&self) {
        for i in 0..(self.frame_i + 1) {
            // println!("print_score_board {}", i);
            let first_throw_i: usize = ((i) * 2) as usize;
            let first_throw = self.pins_down[first_throw_i];

            let second_throw_i = first_throw_i + 1;
            let frame_score = self.score_at_frame(i);

            if (self.pins_down.len() - 1) < second_throw_i {
                print!("[{},  ]({}) ", first_throw, frame_score);
                continue;
            } else {
                let second_throw = self.pins_down[second_throw_i];

                if second_throw == 0 {
                    print!("[{}, -]({}) ", first_throw, frame_score);
                } else {
                    print!("[{}, {}]({}) ", first_throw, second_throw, frame_score);
                }
            }
        }
    }

    

    pub fn score_at_frame(&self, frame_i: usize) -> u16 {
        let first_throw_i = (frame_i * 2) as usize;

        // out of index
        if self.pins_down.len() < first_throw_i + 1 {
            return 0;
        }

        let first_throw = self.pins_down[first_throw_i];

        // if strike
        if first_throw == STRIKE {
            let mut final_score: u16 = TOTAL_PINS;
            let next_throw_i = first_throw_i + 2;
            let next_throw_pins = if self.pins_down.len() > next_throw_i { self.pins_down[next_throw_i] } else { 0 };

            final_score += next_throw_pins;

            // double strike
            if next_throw_pins == STRIKE {
                let next_next_throw_i = first_throw_i + 4;
                let next_next_throw_pins = if self.pins_down.len() > next_next_throw_i { self.pins_down[next_next_throw_i] } else { 0 };
                final_score += next_next_throw_pins;
            } else {
                let next_next_throw_i = first_throw_i + 3;
                let next_next_throw_pins = if self.pins_down.len() > next_next_throw_i { self.pins_down[next_next_throw_i] } else { 0 };
                final_score += next_next_throw_pins;
            }

            return final_score;
        }

        let second_throw_i = first_throw_i + 1;

        // frame still in progress, second throw not yet thrown
        if (self.pins_down.len() - 1) < second_throw_i {
            return 0;
        }

        let second_throw = self.pins_down[second_throw_i];
        let mut final_score = first_throw + second_throw;

        // if spare
        if final_score == TOTAL_PINS {
            let next_throw_i = second_throw_i + 1;
            let next_throw_pins = if self.pins_down.len() > next_throw_i { self.pins_down[next_throw_i] } else { 0 };
            final_score += next_throw_pins
        }

        final_score
    }

    pub fn score(&self) -> Option<u16> {
        if !self.is_done() {
            return None;
        }

        let mut scores: Vec<u16> = Vec::with_capacity(20);

        for i in 0..(self.frame_i + 1) {
            let tmp_frame_score = self.score_at_frame(i);

            if i <= 9 {
                scores.push(tmp_frame_score);
            }
        }

        // println!("> scores: {:?}", scores);

        Some(scores.iter().sum())
    }
}
