<a name="SCM"></a>

# scm
An experiment to investigate greedy optimization of switch-controller mapping

## Usage ##

`go build && ./scm`


## Experiment Set-up 

A set of SDN controllers are initialize to have a total CPU resources CTOTAL and total network bandwidth NTOTAL. A set of switches are initialized and randomly assigned to the SDN controllers. The switches altogether take up CUTIL percentage of CPU resources and NUTIL percentage of network bandwidth. 

## Problem Statement

Find a set of switch migrations SM, so that overloaded controllers can be minimized. 

## Greedy Algorithm

While the problem can be addressed with integer programming, in this experiment we investigate the efficiency of using greedy algorithm. The algorithm is as follows.

```
Sort switches by decreasing load 
for every switch 
  for every controller 
    if migrating switch to controller reduces the cost, migrate 
```

We investigate the effectiveness and efficiencies of the algorithm against the different resource utilization percentages by measuring the runtime of algorithm as well as the final resource utilization of controllers. 

