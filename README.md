# crystal-vs-golang

Compare with golang to check crystal's concurrency ability.


### Why is golang?

Crystal has the sample CONCURRENCY MODEL (using `channels`) like Go, we all konwn that Golang has a good concurrency ability and  It's very hot last frew yeas.

Compare with it, we can more easily to get the crystal concurrency ability, is Powerful or Not.

### Testing System

- My MacBook Air (11-inch, Early 2015), 1.6 GHz Intel Core i5 (2 core), 4 GB 1600 MHz DDR3
- Crystal v0.21.1
- Golang v1.7.3


### Examples

#### countfilelines

This is Crystal official example to count the total number of some files. For this example, we want to get Crystal [Source Code ](https://github.com/crystal-lang/crystal/tree/master/src) lines.

Run `make countfilelines` and get output:

```
real	0m17.833s
user	0m9.668s
sys	0m8.123s
finish countfilelines-crystal test

real	0m8.549s
user	0m3.209s
sys	0m6.861s
finish countfilelines-go test
```

You can see the Go is about 2x faster than Crystal.
