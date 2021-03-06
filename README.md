_Note: This repository is complete **bullshit**._

# Bullshit

Ultra fast, concurrent Bullshit Bingo card generator written in Go.

## Usage

First, compile the `bullshit` program:

```bash
$ go build cmd/bullshit.go
```

Generate 100 Bullshit Bingo cards with four columns and three rows based on the
sample input `computing` into the `output/` directory:

```bash
$ ./bullshit < input/computing -n 100 -c 4 -r 3
```

Generate 25 Bullshit Bingo cards with five columns and four rows based on the
sample input `business` into the `output/` directory:

```bash
$ ./bullshit < input/business -n 25 -c 5 -r 4
```

## Performance

This Bullshit Bingo card generator is ultra fast, thanks to its concurrent
implementation:

```bash
$ time ./bullshit < input/computing -n 1000 -c 4 -r 4
```

    real    0m0.083s
    user    0m0.070s
    sys     0m0.141s
