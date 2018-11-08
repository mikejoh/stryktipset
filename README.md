[![CircleCI](https://circleci.com/gh/mikejoh/stryktipset.svg?style=svg)](https://circleci.com/gh/mikejoh/stryktipset)

# Stryktipset

Stryktipset is one of the brands of the Swedish state-owned company Svenska Spel that operates in the regulated gambling market. It's a betting service where you can play on
various football matches every Saturday.

This repository contains functions and other things related to Stryktipset, like randomizing some bets for you. I created this repository to have a project where i could test and learn Go-lang.

At the moment this repository consists of:

* The `stryktipset` package that exports a number of functions that can be used when playing Stryktipset
* A REST API that serves the `stryktipset` package
* A command line tool called `st` that implements the `stryktipset` package
