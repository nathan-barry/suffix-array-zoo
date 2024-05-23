# Suffix Array Zoo

This repo contains multiple implementation of suffix array construction algorithms, including internal and external memory methods.

## TODO
- [X] Add naive construction algorithm
- [X] Load text files
- [X] Add binary search exact match
- [ ] Implement MM 1990
- [ ] Implement Skew 2003
- [ ] Implement SA-IS 2009
- [ ] Implement eSAIS
- [ ]

## Papers
- [X] Suffix Arrays: A New Method for On-Line String Searches, Manber & Meyers (1990)
- [X] Simple Linear Work Suffix Array Construction (2003) (DC3/skew)
- [ ] Fast Lightweight Suffix Array Construction and Checking (2003)
- [ ] Replacing Suffix Trees with Enhanced Suffix Arrays, Abouelhoda, Kurtz & Ohlebusch (2004)
- [X] A Taxonomy of Suffix Array Construction Algorithms (2007)
- [ ] Better External Memory Suffix Array Construction (2008)
- [ ] Scalable Parallel Suffix Array Construction (2008)
- [ ] Linear Suffix Array Construction by Almost Pure Induced-Sorting, Nong, Zhang & Chan (2009) (SA-IS)
- [ ] Optimal In-Place Suffix Sorting (2016)
- [X] Deduplicating Training Data Makes Language Models Better, Google Research (2022)
    - Based on SA-IS, modified to use external memory but original text must fit in memory.

## Investigate
- [ ] QSufSort (based on 1999 Larsson-Sadakane algorithm)
- [ ] Yuta Mori's DivSufSort "fastest known suffix algo in main memory" as of 2017
- [ ] Ilya Grebnov's even faster implementation 
- [ ] Look at google research's SA-IS extern memory implementation
