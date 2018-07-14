# Simple Race
This application is a simple raffle winner chooser. It selects a winner from a list of contestants by getting the most frequent contestant from 10 random selections.

It is also horribly broken. It'll often select the wrong winner, i.e. one who wasn't selected the most, and sometimes it'll crash.

Can you fix it? Hint: Rand is not thread safe and mutexes can be used to protect data used in multiple go routines.

See the solution folder for the answers.