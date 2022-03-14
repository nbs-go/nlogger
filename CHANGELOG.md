# CHANGELOG

## v2.2.2

- feat(stdlogger): Change namespace writing format
- feat(stdlogger): Change level argument to optional

## v2.2.1

- feat(stdlogger): Change env key for namespace

## v2.2.0

- feat(stdlogger): Change StdPrinterFunc to Printer interface
- feat(level): Add helper to get Level string
- feat(context): Add logContext package

## v2.1.0

- fix: Change module name, add prefix /v2
- feat(stdlogger): Remove error tracing responsibility
- feat(option): Add Context as built-in options

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

