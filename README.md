# COMMUTER


## Synopsis

Ever wonder how long your commute home in traffic is going to take?  Commuter is a CLI tool that will tell you what your commute looks like now and in the future.

## Installation

* [Click here](https://github.com/marioharper/commuter/releases) and download the latest commuter binary.
* Ensure the binary is added to your $PATH.

You're now ready to use Commuter!

## Usage

#### To use commuter you must first initialize the tool:
```sh
commuter init
```

#### To use commuter with the default values(work -> home), just type:
```sh
commuter
```

#### To add or delete addresses:
```sh
commuter add
```
```sh
commuter delete
```

#### To view your commute in google maps:
```sh
# for default values(work -> home)
commuter view 

# for specific
commuter view -f home -t work
```

#### Options:
**number**: specify number of future commutes to show (default: 5)
```sh
commuter -n 20
```

**interval**: specify time between commute results in minutes (default: 15)
```sh
commuter -n 20 -i 10
```
**from/to**: specify which locations you which to use (default: work/home) 
```sh
commuter -f work -t store
```
```sh
commuter -f home -t gym
```
**start**: specify when you will start your commute. Future times only. (default: current time)

Two syntaxs are accepted:
* MMDD:HHMM 24 hour clock example: 0712:1632 = July 12th 4:32PM
* MMDD:HHMM(AM|PM) 12 hour clock example: 0712:0432PM = July 12th 4:43PM
```sh
# start December 14th 4:00PM
commuter -s 1214:1600

# start December 14th 4:00PM
commuter -s 1215:400PM
```

## Contributors

#### [Mario Harper](https://www.marioharper.me)
#### [John Stamatakos](https://github.com/johnstamatakos)
##### Copyright 2016
