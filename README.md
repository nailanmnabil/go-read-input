## Go read input

This repository serves as an attachment for the article where I write about various methods of reading input in Golang, along with their pros and cons.

The data.csv file is not large enough to demonstrate the strengths and weaknesses of all strategies. Therefore, you can run the following command to create three large files of different sizes:


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