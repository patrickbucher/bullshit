# Bullshit

Ultra fast Bullshit Bingo card generator written in Go.

## Usage

First, compile the `bullshit` program:

```bash
$ go build cmd/bullshit.go
```

Generate 100 Bullshit Bingo cards with four columns and three rows based on the
sample input `computing.txt` into the `output/` directory:

```bash
$ ./bullshit < input/computing.txt -n 100 -c 4 -r 3
```

Generate 25 Bullshit Bingo cards with five columns and four rows based on the
sample input `business.txt` into the `output/` directory:

```bash
$ ./bullshit < input/business.txt -n 25 -c 5 -r 4
```

## Performance

This Bullshit Bingo card generator is ultra fast:

```bash
$ time ./bullshit < input/computing.txt -n 1000 -c 4 -r 4
```

    real    0m0.155s
    user    0m0.058s
    sys     0m0.099s

It generates one thousand 4x4 Bullshit Bingo cards in less than 0.2 seconds on
a weak Intel Core i5.
