# strutil

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/odknt/strutil/master/LICENSE) [![Coverage](http://gocover.io/_badge/github.com/odknt/strutil)](http://gocover.io/github.com/odknt/strutil) [![Build Status](https://api.travis-ci.com/odknt/strutil.svg?branch=master)](https://travis-ci.com/odknt/strutil)

Provides utility functions for string and byte slice that no use heap.

## Benchmarks

```
BenchmarkIsAlnum_Strutil-6              30000000                54.0 ns/op             0 B/op          0 allocs/op
BenchmarkIsAlnum_Regexp-6               10000000               144 ns/op               0 B/op          0 allocs/op
BenchmarkValidSubDomain_Strutil-6       100000000               15.0 ns/op             0 B/op          0 allocs/op
BenchmarkValidSubDomain_Regexp-6        10000000               201 ns/op               0 B/op          0 allocs/op
```

```
# The benchmark code uses heap.
# Actual memory allocation count is 1 less.
BenchmarkToBytes_Strutil-6      20000000                84.4 ns/op            32 B/op          1 allocs/op
BenchmarkToBytes_Native-6       10000000               140 ns/op              48 B/op          2 allocs/op
BenchmarkFromBytes_Strutil-6    20000000                68.1 ns/op            16 B/op          1 allocs/op
BenchmarkFromBytes_Native-6     10000000               101 ns/op              32 B/op          2 allocs/op
```
