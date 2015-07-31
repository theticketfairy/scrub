scrub
=====

[![GoDoc](https://godoc.org/github.com/theticketfairy/scrub?status.png)](https://godoc.org/github.com/theticketfairy/scrub)

Simple form validation library for Go

## Motivation

There are at least a couple of very good validation libraries for Go that rely
on struct tags to declare validation rules. We found these became limiting when
wanting to include custom validation rules. We would need to add extra validation
functions to these types.

This package foregoes the use of struct tags in favour of declaring forms which
are sets of fields which in turn are a set of validation functions. These can
be built and customised easily - in a sense, code over configuration.

