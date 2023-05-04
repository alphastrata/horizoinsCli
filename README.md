# Horizons API CLI

A command-line interface (CLI) for downloading data from the NASA Horizons system.

## Usage

To run the program using go run, clone the repository and run the following command from the root directory:

```bash
go run main.go horizonsAPI.go --format=text --command='499' --obj-data=YES --make-ephem=YES --ephem-type
=OBSERVER --center='500@399' --start-time='2006-01-01' --stop-time='2006-01-20' --step-size='1%20d' --quantities='1'
```
Alternatively, you can build the program using go build and then run the executable file:

```bash
go build -o horizons main.go horizonsAPI.go
./horizons --format=text --command='499' --obj-data=YES
```
