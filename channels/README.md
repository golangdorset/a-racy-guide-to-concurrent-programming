# Channels
This application contacts a set of URLs and gets statistics on them (response code and latency).
We want these calls to happen concurrently for sped purposes.

But yet again this code has data races and you'll notice causes a hanging process.

Can you fix it? Hint: Think about how you pass data to the goroutines and how you can determine all urls have been contacted.

See the solution folder for the answers.