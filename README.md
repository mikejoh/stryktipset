# Stryktipset

Stryktipset is one of the brands of the Swedish state-owned company Svenska Spel that operates in the regulated gambling market.

This repository contains tools and other things related to Stryktipset, i created this repository to have a project where i could test and proof-of-concept Go-lang specific design patterns and basically just write some code.

At the moment this repository consists of:

* The `stryktipset` package that exports a number of functions that can be used when playing Stryktipset
* A REST API that serves the `stryktipset` package
* A command line tool called `st` that implements the `stryktipset` package