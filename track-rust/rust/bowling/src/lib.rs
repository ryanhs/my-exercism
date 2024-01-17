use std::fmt::{self};


#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

pub struct BowlingFrame {
    scores: Vec<u16>,
    current_pins_left: u16,
}


impl BowlingFrame {
    pub fn new() -> Self {
        Self {
            current_pins_left: 10,
            scores: Vec::new(),
        }
    }
}

impl fmt::Debug for BowlingFrame {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "scores {:?} => {}", self.scores, self.current_pins_left)
    }
}

pub struct BowlingGame {
    current_frame: usize,
    frames: Vec<BowlingFrame>,
    
    is_finished: bool,
}

impl BowlingGame {
    pub fn new() -> Self {
        Self {
            frames: vec![BowlingFrame::new()],
            current_frame: 0,
            is_finished: false,
        }
    }

    fn current_frame_pins_remaining(&mut self) -> u16{
        if self.current_frame < self.frames.len() {
            return 0;
        }

        self.frames[self.current_frame].current_pins_left
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if pins > 10 || pins > self.current_frame_pins_remaining() {
            return Err(Error::NotEnoughPinsLeft);
        }

        let current_frame = &mut self.frames[self.current_frame];
        
        current_frame.scores.push(pins);
        current_frame.current_pins_left -= pins;

        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        todo!("Return the score if the game is complete, or None if not.");
    }
}
