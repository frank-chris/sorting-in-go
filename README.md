# Sorting in Go

A simple sorting program written in Go to get familiar with the Go syntax.

## Sort specification

This project reads, writes, and sorts files consisting of zero or
more records.  A record is a 100 byte binary key-value pair, consisting
of a 10-byte key and a 90-byte value.  Each key is interpreted
as an unsigned 10-byte (80-bit) integer.  The sort is ascending,
meaning that the output has the record with the smallest key first,
then the second-smallest, and so on.

## Utility scripts

The utility scripts in `utils\` work on x86_64-based Linux, Intel MacOS and Apple M1 MacOS.

### Gensort

Gensort generates random input.  If the 'randseed' parameter is provided,
the given seed is used to ensure deterministic output.

'size' can be provided as a non-negative integer to generate that many
bytes of output. However human-readable strings can be used as well,
such as "10 mb" for 10 megabytes, "1 gb" for one gigabyte", "256 kb"
for 256 kilobytes, etc.

If the specified size is not a multiple of 100 bytes, the requested
size will be rounded up to the next multiple of 100.

Usage: 

```
utils/{os_architecture}/bin/gensort outputfile size
  -randseed int
    	Random seed
```

### Showsort

Showsort shows the records in the provided file in a human-readable
format, with the key followed by a space followed by an
abbreviated version of the value.

Usage: 

```
utils/{os_architecture}/bin/showsort inputfile
```

### Valsort

Valsort scans the provided input file to check if it is sorted.

Usage: 

```
utils/{os_architecture}/bin/valsort inputfile
```

## Sort program

The sort program (source in `src\`, binary in `bin\`) reads in an input file and
produces a sorted version of the output file.

Usage: 

```
bin/sort inputfile outputfile
```

## Building

To build the sort program:

```
go build -o bin/sort src/sort.go
```

## Verifying sort implementation

A simple way to verify the correctness of the implementation of sort
is to run the standard unix sort command on the output of 'showsort'.
For example, to generate, sort, and verify a 1 megabyte file:

```bash
$ utils/linux-amd64/bin/gensort example1.dat "1 mb"
No random seed provided, using current timestamp
Initializing the random seed to 1641144051385376000
Requested 1 mb (= 1048576) bytes
Increasing size to 1048600 to be divisible by 100

$ bin/sort example1.dat example1-sorted.dat
2022/03/02 09:14:36 sort.go:59: Read in 10486 records

$ utils/linux-amd64/bin/valsort example1-sorted.dat
File is sorted

$ utils/linux-amd64/bin/showsort example1.dat | sort > example1-chk.txt
$ utils/linux-amd64/bin/showsort example1-sorted.dat| sort > example1-sorted-chk.txt
$ diff example1-chk.txt example1-sorted-chk.txt
```

This last 'diff' should simply return to the command prompt. If it
indicates any differences that means that there is an error
in the sort routine.