# This is how Load Average Metric is calculated in Linux

<hr>

## Versions
- 1.0
    - Hypothetically calculates Load Average for 1 minute
    - Amount of **Active tasks** (Running + Runnable + Uninterraptible Sleep) starts from **zero** and increases on every iteration by 60% chance and decreases on every iteration by 40% chance. 
    - Exponent is taken 0.08
    - [Exponentially weighted moving average (EWMA)](https://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average) is used for calculations
- 2.0
    - Added bitwise operations
    - Full copy from real Linux Sources where this [algorithm](https://github.com/torvalds/linux/blob/2ba9f676d0a2e408aef14d679984c26373bf37b7/include/linux/sched/loadavg.h#L4) is written in C
