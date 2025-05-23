# About
Blackjack is one of the first applications I have developed to learn the go programming language.

## How to use

``` shell
make build
./blackjack

# To view options for the game.
./blackjack -h

# Set number of decks to play with and decide whether dealer gets to hit on soft 17 or not.
./blackjack -decks 6 -h17=false
```


## How to debug

``` shell
export DEBUG=1
make run
```

In a second shell:

``` shell
tail -f debug.log
```

# Acknowledgements
* [bubbletea](https://github.com/charmbracelet/bubbletea)
