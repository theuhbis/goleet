# Two Sum Solution

[View Problem on Leetcode](https://leetcode.com/problems/two-sum/)

## Algorithm

It is possible to solve the two sum problem with one pass through the array
using a hash table.

While iterating and inserting elements into the hash table, also check the
keys to see if the current element's complement already exists in the hash
table. If it exists, a solution has been found and indices can be returned 
immediately.