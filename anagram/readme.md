# Valid Anagram Solution

[View Problem on Leetcode](https://leetcode.com/problems/valid-anagram/)

## Algorithm

It is possible to solve the valid anagram problem with one pass through each string using a hash table.

Iterate through the first string and insert/count characters from the string in the hash table. Next iterate through the second string and check the characters against the hash table if it isn't in the hash table return false. If it is, decrement until 0 and then remove. Return true if second string is fully iterated through and hash table is empty.