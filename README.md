Performance Difference Between Indices and Values
======================================

# The Problem
Ever wondered which is faster?

### Method A
```go
for index := range list {
  fmt.Sprint(list[index])
}```

### Method B
```go
for _, value := range list {
  fmt.Sprint(value)
}```

## The Experiment
Two benchmark functions which both call `fmt.Sprint` for the targeted object. 
One function uses the indices, the other uses the values. 
Both iterate over a float64 slice with 100,000 (semi-random) elements. ;

In order to neglect noise in the measurements, I decided to let it run for 30 seconds each. 

The command I used is:
`go test -benchtime=30s -bench=.`

## The Results
There is no performance difference at all between method A and method B. Even though the results say there is a 
slight performance difference, this is not the case in the long run. 

### First Run
```
BenchmarkValue      1000          51645997 ns/op
BenchmarkIndex      1000          51507547 ns/op
```

### Second Run
```
BenchmarkValue      1000          51665856 ns/op
BenchmarkIndex      1000          51512520 ns/op
```

# Extending the Problem
What if we need to use the value more than once? Would using the `value` be faster? Let's find out. 

## The Experiment
The set-up is the same as describe above, except that I am now using the value five times in total.

## The Results
There is no performance difference between method A and method B, even when calling the desired object multiple times. 

### First Run
```
BenchmarkValueMultiple       200         264178475 ns/op
BenchmarkIndexMultiple       200         260632222 ns/op
```
### Second Run
```
BenchmarkValueMultiple       200         257902368 ns/op
BenchmarkIndexMultiple       200         263454376 ns/op
```