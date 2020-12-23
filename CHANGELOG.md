# CHANGELOG

## v0.2.0 (2020-12-23)

### Others

- build: run go mod tidy to remove go-cmp

- test(evaluate): replace go-cmp with assert

- test(parser): replace go-cmp with assert

- test(scanner): replace go-cmp with assert for buffer_scanner

- test(scanner): replace go-cmp with assert

- refactor(evaluate): gofmt binary

- chore: bump golangci-lint v1.33 in github action

- chore(changelog): generate v0.1.0

## v0.1.0 (2020-12-23)

### Added

- feat: add visit in, not in

- feat: add visit <, <=, >, >=

- feat: represent TokenText in string

- feat: add string method for expression

- feat: add string represent for token

- feat: evaluator with simple example

### Fixed

- fix: assert right expr not left

- fix: correct visit equal using switch type

- fix: make scanner not lowercase all ident

- fix: token stores lowercase scanner.Ident as Text

### Others

- chore: add MIT LICENSE

- chore: add github action for test, lint

- test: unit test for not in

- test: simple unittest for visit in

- refactor: make new expression return pointer not interface

- test: unittest for visit unary not

- test: unittest for visit <, <=, >, >=

- test: unittest for and, or evaluate

- test: move evaluate array to visit literal testcase

- refactor: make evaluate visitor receive pointer

- refactor: new expression return expression interface

- refactor: use new bool literal instead of struct

- refactor: use new fn to init expression when visit var

- refactor: rename visitor -> evaluate visitor

- refactor: remove evaluate

- refactor: remove examples

- test: complex parser to text reduce parenthesis

- refactor: return token string represent for error

- test: testing scanner ident lower, upper, mixed case

- refactor: better error when visit unary, binary

- refactor: remove parenthesis

- refactor: use new expression fn instead of raw init

- refactor: replace visit array with visit literal

- refactor: use single visit literal

- refactor: better string represent for token and expression

- refactor: make visit return expression not interface

- test: remove benchmark

- refactor: rename parseExpression -> parseWithPrecedence

- refactor: remove naked return in parse expression

- test: unittest for parser binary expression

- refactor: use pointer for expression

- refactor: improve return error for parser

- refactor: lower text right after scanner

- refactor: get token precendence from map

- refactor: nullDenotation, leftDenotation -> nud, led for consistent

- refactor: split parser to nud, led file

- refactor: split visitor implement from evaluate to visitor pkg

- refactor: rename expression_xxx -> xxx

- docs: nud, led explain right before fn

- test: unittest for buffer scanner peek

- refactor: use go-cmp to replace testify

- test: use go-cmp to test buffer scanner scan

- refactor: testCase -> scannerTestCase

- docs: explain scan and peek method for buffer scanner

- refactor: remove undoScan method of buffer scanner

- docs: add comment to explain scan method of buffer scanner

- test: unittest for eof, illegal scanner

- docs: add comment for scanner when remove ""

- test: unittest for scanner lower, mixed, uppercase in, not in

- style: rename bufferscanner -> buffer_scanner

- chore: add Makefile to support test, lint