## Go read input

This repository is the attachment for [this article](https://nabil.top/blogs/go-read-input) where i write about all the different ways to read an input in Golang with its pros and cons

The data.csv file is not big enough to prove all strategies strength and weaknesses, so you can run this command to create 3 big file with its different sizes.

```
$ for i in {1..5000}; do cat data.csv >> data_sm.csv; done
$ for i in {1..90000}; do cat data.csv >> data_md.csv; done
$ for i in {1..300000}; do cat data.csv >> data_big.csv; done
```

```
$ go test -bench=Small
$ go test -bench=Med
$ go test -bench=Big
```