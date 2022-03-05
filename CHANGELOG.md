# CHANGELOG

## v2.0.0

- feat(stdlogger): Change implementation on StdLogger to follow Logger interface
- BREAKING CHANGE: Change args type to distinct between logging option and formatted logging functions

## v1.4.3

- [FIXED] Fix FormatArgs option evaluation

## v1.4.2

- [FIXED] Fix formatted args logging

## v1.4.1

- [ADDED] Add context aware logger child and custom print formatter
- [FIXED] Fix formatted args option evaluation

## v1.3.0

- [CHANGED] Change Options type to `map[string]interface{}` for re-usability

## v1.2.0

- [ADDED] Add NewChild function

## v1.1.0

- [ADDED] Add NewChild implementation in StdLogger
- [ADDED] Add NewChild function to Logger interface

## v1.0.0

- [ADDED] Add NewChild implementation in StdLogger
- [ADDED] Add NewChild function to Logger interface
- [ADDED] Add StdLogger as fallback Logger implementation
- [ADDED] Add logger interface

