# A Racy Guide to Concurrent Programming
3 small applications to demonstrate concurrency and data races in Go.

## Simple Race
A version of our "raffle winner" application from the June meetup. But now with added concurrency!

## Mutexes
A toy bank app; withdraw money concurrently from accounts handled by a `Bank` struct.

## Channels
A URL scraper

> All of these applications are subtly broken! The goal is to use the race detector to figure out how to fix them.
