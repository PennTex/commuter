# COMMUTER


## Synopsis

Ever wonder how long your commute home in traffic is going to take?  Commuter is a CLI tool that will tell you what your commute looks like now and in the future.

## Installation

To install, all you need is a Google Directions API key:
* [Click here](https://github.com/marioharper/commuter/releases) and download the latest commuter binary
* Open and install the binary file
* [Click here](https://developers.google.com/maps/documentation/directions/get-api-key) to go to Google's Developer website
* Click "Get a Key" and answer a few questions
* Copy your key
* Open your terminal and type 
```sh 
export MAPS_API_KEY=<your-api-key> 
```
You're now ready to use Commuter!

## Code Example

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

#### Options:
* number: specify number of future commutes to show (default: 5)
```sh
commuter -n 20
```
* interval: specify time between commute results in minutes (default: 15)
```
commuter -n 20 -i 10
```
* from/to: specify which locations you which to use (default: work/home) 
```
commuter -f work -t store
```
```
commuter -f home -t gym
```

## Contributors

#### Mario Harper
#### John Stamatakos
##### Copyright 2016
